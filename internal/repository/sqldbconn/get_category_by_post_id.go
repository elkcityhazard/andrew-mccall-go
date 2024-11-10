package sqldbconn

import (
	"context"
	"time"

	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
)

func (sdc *SQLDbConn) GetCategoryByPostID(postID int64) (*models.Category, error) {
	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*10)

	defer cancel()

	catChan := make(chan *models.Category)
	errChan := make(chan error)

	sdc.app.WG.Add(1)

	go func() {
		defer sdc.app.WG.Done()
		defer close(catChan)
		defer close(errChan)

		stmt := `
		SELECT categories.id,categories.name,categories.created_at, categories.updated_at,categories.version FROM categories
		INNER JOIN category_joins
		ON category_joins.post_id = ?
		`

		c := &models.Category{}

		err := sdc.conn.QueryRowContext(ctx, stmt, postID).Scan(&c.ID, &c.Name, &c.CreatedAt, &c.UpdatedAt, &c.Version)

		if err != nil {
			errChan <- err
			return
		}

		catChan <- c

	}()

	select {
	case cat := <-catChan:
		return cat, nil
	case err := <-errChan:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}

}
