package sqldbconn

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	"html"
	"time"

	"github.com/elkcityhazard/andrew-mccall-go/internal/config"
	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
)

type SQLDbConn struct {
	app  *config.AppConfig
	conn *sql.DB
}

func NewSQLDbConn(app *config.AppConfig, conn *sql.DB) *SQLDbConn {
	return &SQLDbConn{
		app:  app,
		conn: conn,
	}
}
func (sdc *SQLDbConn) ActivateUser(u *models.User) (int64, error) {

	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*10)

	defer cancel()

	idChan := make(chan int64)
	errChan := make(chan error)

	sdc.app.WG.Add(1)

	go func() {
		defer sdc.app.WG.Done()
		defer close(idChan)
		defer close(errChan)

		tx, err := sdc.conn.BeginTx(ctx, &sql.TxOptions{})

		if err != nil {
			tx.Rollback()
			errChan <- err
			return
		}

		stmt := `
		UPDATE users 
		SET users.is_active = ?,
		users.updated_at = ?,
		users.version = users.version + 1
		WHERE users.id = ?
		AND users.version = ?
		`

		args := []any{1, time.Now(), u.ID, u.Version}

		result, err := tx.ExecContext(ctx, stmt, args...)

		if err != nil {
			tx.Rollback()
			errChan <- err
			return
		}

		id, err := result.LastInsertId()

		if err != nil {
			tx.Rollback()
			errChan <- err
			return
		}

		// delete token
		delStmt := `
	DELETE FROM activation_tokens
	WHERE activation_tokens.user_id = ?
	`

		delResult, err := tx.ExecContext(ctx, delStmt, u.ID)

		if err != nil {
			tx.Rollback()
			errChan <- err
			return
		}

		affected, err := delResult.RowsAffected()

		if err != nil {
			tx.Rollback()
			errChan <- err
			return
		}

		if affected == 0 {
			tx.Rollback()
			errChan <- errors.New("an error has occurred")
			return
		}

		err = tx.Commit()

		if err != nil {
			tx.Rollback()
			errChan <- err
			return
		}

		idChan <- id

	}()

	select {
	case id := <-idChan:
		return id, nil
	case err := <-errChan:
		return 0, err
	}

}
func (sdc *SQLDbConn) GetActivationToken(tokenString string) (*models.User, *models.ActivationToken, error) {

	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*10)

	defer cancel()

	tokenHash := sha256.Sum256([]byte(tokenString))
	userChan := make(chan *models.User)
	tokenChan := make(chan *models.ActivationToken)
	errChan := make(chan error)
	doneChan := make(chan bool)

	sdc.app.WG.Add(1)

	go func() {
		defer sdc.app.WG.Done()

		stmt := `
		select
		users.id,
		users.email,
		users.username,
		users.password,
		users.created_at,
		users.updated_at,
		users.is_active,
		users.role,
		users.version,
		passwords.id,
		passwords.hash,
		passwords.created_at,
		passwords.updated_at,
		passwords.is_active,
		passwords.version,
		activation_tokens.id,
		activation_tokens.hash,
		activation_tokens.user_id,
		activation_tokens.expiry,
		activation_tokens.scope
		FROM users
		INNER JOIN
		activation_tokens ON
		users.id = activation_tokens.user_id
		INNER JOIN passwords ON
		users.password = passwords.id
		WHERE activation_tokens.hash = ?
		AND activation_tokens.expiry > ?
		AND passwords.is_active = 0
		`

		pw := models.Password{}
		u := models.User{}
		u.Password = &pw
		at := models.ActivationToken{}

		args := []any{tokenHash[:], time.Now()}

		err := sdc.conn.QueryRowContext(ctx, stmt, args...).Scan(
			&u.ID,
			&u.Email,
			&u.Username,
			&u.Password.ID,
			&u.CreatedAt,
			&u.UpdatedAt,
			&u.IsActive,
			&u.Role,
			&u.Version,
			&pw.ID,
			&pw.Hash,
			&pw.CreatedAt,
			&pw.UpdatedAt,
			&pw.IsActive,
			&pw.Version,
			&at.ID,
			&at.Hash,
			&at.UserID,
			&at.Expiry,
			&at.Scope,
		)

		if err != nil {
			errChan <- err
			doneChan <- true
			return
		}

		at.Plaintext = tokenString

		u.Password = &pw

		userChan <- &u

		tokenChan <- &at

		doneChan <- true

	}()

	var u *models.User
	var t *models.ActivationToken
	var e error

	for {
		select {
		case usr := <-userChan:
			u = usr
		case acTok := <-tokenChan:
			t = acTok
		case err := <-errChan:
			e = err
			return u, t, e
		case <-doneChan:
			close(tokenChan)
			close(userChan)
			close(errChan)
			close(doneChan)
			return u, t, e
		case <-ctx.Done():
			e = ctx.Err()
			return u, t, e
		}
	}

}
func (sdc *SQLDbConn) InsertActivationToken(at *models.ActivationToken) (int64, error) {
	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*10)

	defer cancel()

	tokenChan := make(chan *models.ActivationToken)
	errorChan := make(chan error)

	sdc.app.WG.Add(1)

	go func() {
		defer close(tokenChan)
		defer close(errorChan)

		defer sdc.app.WG.Done()

		stmt := `

		INSERT INTO activation_tokens 
		(hash, user_id, expiry,scope)
		VALUES(?,?,?,?)
		`

		args := []any{at.Hash, at.UserID, at.Expiry, at.Scope}

		result, err := sdc.conn.ExecContext(ctx, stmt, args...)

		if err != nil {
			errorChan <- err
			return
		}

		id, err := result.LastInsertId()

		if err != nil {
			errorChan <- err
			return
		}

		at.ID = id

		tokenChan <- at

	}()

	select {
	case token := <-tokenChan:
		return token.ID, nil
	case err := <-errorChan:
		return 0, err
	}

}
func (sdc *SQLDbConn) GetNextPrevPost(post *models.Content, increment bool) (*models.Content, error) {
	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*10)

	defer cancel()

	postChan := make(chan *models.Content)
	errorChan := make(chan error)
	sdc.app.WG.Add(1)

	go func() {
		defer sdc.app.WG.Done()
		defer close(postChan)
		defer close(errorChan)

		var stmt string

		if !increment {
			stmt = `
		SELECT 
			posts.id, 
			posts.title,
			posts.slug 
			FROM posts 
			WHERE posts.created_at < ? AND posts.status = "published"
			ORDER BY posts.created_at
			DESC
			LIMIT 1;
		`
		} else {

			stmt = `
		SELECT 
			posts.id, 
			posts.title, 
			posts.slug 
			FROM posts 
			WHERE posts.created_at > ?  AND posts.status = "published"
			ORDER BY posts.created_at
			ASC
			LIMIT 1;
		`
		}

		var p models.Content

		err := sdc.conn.QueryRowContext(ctx, stmt, post.CreatedAt).Scan(&p.ID, &p.Title, &p.Slug)

		if err != nil {
			errorChan <- err
			return
		}

		postChan <- &p

	}()

	select {
	case post := <-postChan:
		return post, nil
	case err := <-errorChan:
		return nil, err
	}

}
func (sdc *SQLDbConn) GetUserByID(id int64) (*models.User, error) {
	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*15)

	defer cancel()

	userChan := make(chan *models.User)
	errChan := make(chan error)
	sdc.app.WG.Add(1)

	go func() {
		defer close(userChan)
		defer close(errChan)
		defer sdc.app.WG.Done()

		stmt := `
		SELECT id,email,username FROM users WHERE id = ?
		`
		user := models.User{}

		err := sdc.conn.QueryRowContext(ctx, stmt, id).Scan(&user.ID, &user.Email, &user.Username)

		if err != nil {
			errChan <- err
			return
		}

		userChan <- &user

	}()

	select {
	case err := <-errChan:
		return nil, err
	case user := <-userChan:
		return user, nil
	}
}
func (sdc *SQLDbConn) ListPosts(limit, offset int) ([]*models.Content, error) {
	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*15)

	defer cancel()

	postChan := make(chan *models.Content)
	errChan := make(chan error)
	doneChan := make(chan bool)

	sdc.app.WG.Add(1)

	go func() {
		defer sdc.app.WG.Done()

		stmt := `
		SELECT 
			posts.id,
			posts.user_id,
			posts.title,
			posts.slug,
			posts.description,
			posts.content,
			posts.delta,
			posts.created_at,
			posts.updated_at,
			posts.status,
			posts.version,
			users.email
		FROM posts
		INNER JOIN
		users
		On posts.user_id = users.id
		WHERE
			posts.created_at <= CURRENT_TIMESTAMP
		AND posts.status = ?
		ORDER BY posts.created_at DESC
		LIMIT ?
		OFFSET ?
		`

		args := []any{"published", limit, offset}

		rows, err := sdc.conn.QueryContext(ctx, stmt, args...)

		if err != nil {
			errChan <- err
			return
		}

		defer rows.Close()

		for rows.Next() {
			var c models.Content

			var email string

			err := rows.Scan(
				&c.ID,
				&c.UserId,
				&c.Title,
				&c.Slug,
				&c.Description,
				&c.Content,
				&c.Delta,
				&c.CreatedAt,
				&c.UpdatedAt,
				&c.Status,

				&c.Version,
				&email)
			if err != nil {
				errChan <- err
				return
			}

			postChan <- &c
		}

		doneChan <- true

	}()

	var posts []*models.Content

	for {
		select {
		case post := <-postChan:
			posts = append(posts, post)
		case err := <-errChan:
			if err != nil {
				close(postChan)
				close(errChan)
				close(doneChan)
				return nil, err
			}
		case <-doneChan:
			close(postChan)
			close(errChan)
			close(doneChan)
			return posts, nil

		}
	}

}

func (sdc *SQLDbConn) GetPaginatedPosts(userID int64, offset, size int) ([]*models.Content, error) {
	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*15)

	defer cancel()

	postChan := make(chan *models.Content)
	errChan := make(chan error)
	doneChan := make(chan bool)

	sdc.app.WG.Add(1)

	go func() {
		defer sdc.app.WG.Done()

		stmt := `
		SELECT id,title,slug,description,created_at,updated_at,status,version
		FROM posts
		WHERE user_id = ? 
		LIMIT ? 
		OFFSET ?
		`

		args := []any{userID, size, offset}

		rows, err := sdc.conn.QueryContext(ctx, stmt, args...)

		if err != nil {
			errChan <- err
			doneChan <- true
			return
		}

		defer rows.Close()

		for rows.Next() {

			var c = models.Content{}

			err := rows.Scan(&c.ID, &c.Title, &c.Slug, &c.Description, &c.CreatedAt, &c.UpdatedAt, &c.Status, &c.Version)

			if err != nil {
				fmt.Println(err)
				errChan <- err
				doneChan <- true
			}
			postChan <- &c
		}
		if err = rows.Err(); err != nil {
			errChan <- errors.New("no rows")
		}
		doneChan <- true

	}()

	var posts = []*models.Content{}

	var err error

	for {
		select {
		case post := <-postChan:
			posts = append(posts, post)

		case errOp := <-errChan:
			err = errOp

		case <-doneChan:
			close(postChan)
			close(errChan)
			close(doneChan)

			if err != nil {
				return nil, err
			}

			return posts, nil
		}
	}

}
func (sdc *SQLDbConn) UpdatePost(c *models.Content) (int64, error) { //	returns version
	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*15)
	defer cancel()

	affectedChan := make(chan int64)
	errorChan := make(chan error)

	sdc.app.WG.Add(1)

	go func() {
		defer close(affectedChan)
		defer close(errorChan)
		defer sdc.app.WG.Done()

		stmt := `UPDATE posts
		SET
		title = ?,
		slug = ?,
		description = ?,
		content = ?,
		delta = ?,
		updated_at = ?,
		version = version + 1,
		status = ?
		WHERE id = ? AND version = ?`

		args := []any{c.Title, c.Slug, c.Description, c.Content, c.Delta, c.UpdatedAt, c.Status, c.ID, c.Version}

		tx, err := sdc.conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault})

		if err != nil {
			errorChan <- err
			return
		}

		result, err := tx.ExecContext(ctx, stmt, args...)

		if err != nil {
			_ = tx.Rollback()
			errorChan <- err
			return
		}

		if err := tx.Commit(); err != nil {
			errorChan <- err
			return
		}

		affected, err := result.RowsAffected()

		if err != nil {
			errorChan <- err
			return
		}

		if affected > 0 {

			affectedChan <- affected

		}
	}()
	for {
		select {
		case affected := <-affectedChan:
			return affected, nil
		case err := <-errorChan:
			return 0, err
		}
	}

}
func (sdc *SQLDbConn) GetBlogPostByID(id int64) (*models.Content, error) {

	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*15)

	defer cancel()

	postChan := make(chan *models.Content)
	errorChan := make(chan error)

	sdc.app.WG.Add(1)

	go func() {
		defer sdc.app.WG.Done()
		defer close(postChan)
		defer close(errorChan)

		stmt := `SELECT id,user_id,title,slug,description,content,delta,created_at,updated_at,status,version FROM posts WHERE id = ?`

		post := models.Content{}

		err := sdc.conn.QueryRowContext(ctx, stmt, id).Scan(&post.ID, &post.UserId, &post.Title, &post.Slug, &post.Description, &post.Content, &post.Delta, &post.CreatedAt, &post.UpdatedAt, &post.Status, &post.Version)

		if err != nil {
			errorChan <- err
			return
		}

		postChan <- &post

	}()

	select {
	case err := <-errorChan:
		return nil, err
	case post := <-postChan:
		return post, nil
	}
}
func (sdc *SQLDbConn) GetBlogPost(routekey string) (*models.Content, error) {

	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*15)

	defer cancel()

	postChan := make(chan *models.Content)
	errorChan := make(chan error)

	sdc.app.WG.Add(1)

	go func() {
		defer sdc.app.WG.Done()
		defer close(postChan)
		defer close(errorChan)

		stmt := `SELECT id,user_id,title,slug,description,content,created_at,updated_at,status,version FROM posts WHERE slug = ?`

		post := models.Content{}

		err := sdc.conn.QueryRowContext(ctx, stmt, routekey).Scan(&post.ID, &post.UserId, &post.Title, &post.Slug, &post.Description, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.Status, &post.Version)

		if err != nil {
			errorChan <- err
			return
		}

		postChan <- &post

	}()

	select {
	case err := <-errorChan:
		return nil, err
	case post := <-postChan:
		return post, nil
	}
}

func (sdc *SQLDbConn) InsertEditorContent(c *models.Content) (int64, error) {

	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*15)

	defer cancel()

	idChan := make(chan int64)
	errorChan := make(chan error)

	sdc.app.WG.Add(1)

	go func() {
		defer sdc.app.WG.Done()
		defer close(idChan)
		defer close(errorChan)

		stmt := `INSERT INTO 
		posts 
		(
			user_id,
			title,
			slug,
			description,
			content,
			delta,
			created_at,
			updated_at,
			status,
			version
		) 
		VALUES(?,?,?,?,?,?,CURRENT_TIMESTAMP,CURRENT_TIMESTAMP,?,1)`

		args := []any{c.UserId, c.Title, c.Slug, c.Description, c.Content, c.Delta, c.Status}

		tx, err := sdc.conn.BeginTx(ctx, nil)

		if err != nil {
			tx.Rollback()
			return
		}

		result, err := tx.ExecContext(ctx, stmt, args...)

		if err != nil {
			errorChan <- err
			tx.Rollback()
			return
		}

		insertedID, err := result.LastInsertId()

		if err != nil {
			errorChan <- err
			tx.Rollback()
			return
		}

		err = tx.Commit()

		if err != nil {
			errorChan <- err
			tx.Rollback()
			return
		}

		c.ID = insertedID

		idChan <- insertedID
	}()

	select {
	case insertedID := <-idChan:
		return insertedID, nil
	case err := <-errorChan:
		return 0, err
	}

}

func (sdc *SQLDbConn) GetUserByEmail(email string) (*models.User, error) {

	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*10)

	defer cancel()

	userChan := make(chan *models.User, 1)
	errorChan := make(chan error, 1)

	sdc.app.WG.Add(1)

	go func() {
		defer sdc.app.WG.Done()
		defer close(userChan)
		defer close(errorChan)

		stmt := `
		select 
		users.id,
		users.email,
		users.username,
		users.created_at,
		users.updated_at,
		users.is_active,
		users.role,
		users.version,
		passwords.id,
		passwords.hash
		FROM users
		inner join 
		passwords on
		users.password = passwords.id
		WHERE users.email = ?
		`

		u := &models.User{}
		p := &models.Password{}
		err := sdc.conn.QueryRowContext(ctx, stmt, email).Scan(&u.ID, &u.Email, &u.Username, &u.CreatedAt, &u.UpdatedAt, &u.IsActive, &u.Role, &u.Version, &p.ID, &p.Hash)

		if err != nil {
			errorChan <- err
			return
		}

		u.Password = p

		userChan <- u
	}()

	for {
		select {
		case user := <-userChan:
			return user, nil
		case err := <-errorChan:
			return nil, err
		}
	}

}

func (sdc *SQLDbConn) InsertUser(u *models.User) (int64, error) {
	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*10)

	defer cancel()

	idChan := make(chan int64, 1)
	errorChan := make(chan error, 1)

	sdc.app.WG.Add(1)

	go func() {
		defer sdc.app.WG.Done()
		defer close(idChan)
		defer close(errorChan)
		tx, err := sdc.conn.BeginTx(ctx, nil)
		if err != nil {
			errorChan <- err
			tx.Rollback()
			return
		}
		pwStmt := `INSERT INTO passwords (hash) VALUES(?)`
		pwResult, err := tx.ExecContext(ctx, pwStmt, u.Password.Hash)
		if err != nil {
			errorChan <- err
			tx.Rollback()
			return
		}

		pwID, err := pwResult.LastInsertId()
		if err != nil {
			errorChan <- err
			tx.Rollback()
			return
		}

		u.Password.ID = pwID

		usrStmt := `INSERT INTO users (email,username,password,role, is_active) VALUES(?,?,?,?,?)`

		usrArgs := []interface{}{u.Email, u.Username, u.Password.ID, u.Role, u.IsActive}

		usrResult, err := tx.ExecContext(ctx, usrStmt, usrArgs...)

		if err != nil {
			errorChan <- err
			tx.Rollback()
			return
		}

		usrID, err := usrResult.LastInsertId()
		if err != nil {
			errorChan <- err
			tx.Rollback()
			return
		}

		u.ID = usrID

		idChan <- u.ID

		err = tx.Commit()

		if err != nil {
			errorChan <- err
			tx.Rollback()
		}

	}()

	return <-idChan, <-errorChan

}

func (sdc *SQLDbConn) GetTotalCount(table string) (int, error) {
	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*15)

	defer cancel()

	countChan := make(chan int, 1)
	errChan := make(chan error, 1)
	sdc.app.WG.Add(1)

	go func() {
		defer close(countChan)
		defer close(errChan)
		defer sdc.app.WG.Done()

		var count int

		stmt := fmt.Sprintf("SELECT COUNT(id) FROM %s", html.EscapeString(table))

		err := sdc.conn.QueryRowContext(ctx, stmt).Scan(&count)

		if err != nil {
			errChan <- err
			return
		}

		countChan <- count

	}()

	select {
	case count := <-countChan:
		return count, nil
	case err := <-errChan:
		return 0, err
	}

}
