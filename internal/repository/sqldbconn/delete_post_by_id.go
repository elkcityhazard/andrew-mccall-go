package sqldbconn

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
)

func (sdc *SQLDbConn) DeletePostById(id, userID int64) (int64, error) {
	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*10)
	defer cancel()

	affectedChan := make(chan int64)
	errChan := make(chan error)
	sdc.app.WG.Add(1)

	go func() {
		defer sdc.app.WG.Done()
		defer close(affectedChan)
		defer close(errChan)

		tx, err := sdc.conn.BeginTx(ctx, &sql.TxOptions{})

		if err != nil {
			tx.Rollback()
			errChan <- err
			return
		}

		// check user owns document
		gpStmt := `SELECT posts.id,posts.user_id from posts WHERE posts.id = ?`

		var p models.Content

		err = sdc.conn.QueryRowContext(ctx, gpStmt, id).Scan(&p.ID, &p.UserId)

		if err != nil {
			tx.Rollback()
			errChan <- err
			return
		}

		if p.UserId != userID {
			tx.Rollback()
			err := errors.New("you are not authorized")
			errChan <- err
			return
		}

		stmt := `DELETE FROM posts WHERE posts.id = ?`

		args := []any{id}

		result, err := sdc.conn.ExecContext(ctx, stmt, args...)

		if err != nil {
			errChan <- err
			return
		}

		affected, err := result.RowsAffected()

		if err != nil {
			errChan <- err
			return
		}
		affectedChan <- affected

		err = tx.Commit()

		if err != nil {
			tx.Rollback()
			errChan <- err
			return
		}

	}()

	select {
	case err := <-errChan:
		return 0, err
	case affected := <-affectedChan:
		return affected, nil
	}

}
