package sqldbconn

import (
	"context"
	"time"

	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
)

func (sdc *SQLDbConn) ListCategories() ([]*models.Category, error) {
	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*15)

	defer cancel()

	catChan := make(chan *models.Category)
	errChan := make(chan error)
	doneChan := make(chan bool)

	sdc.app.WG.Add(1)

	go func() {
		defer sdc.app.WG.Done()
		defer close(catChan)
		defer close(errChan)
		defer close(doneChan)

		stmt := `SELECT id, name, created_at, updated_at, version FROM categories`

		rows, err := sdc.conn.QueryContext(ctx, stmt)

		if err != nil {
			errChan <- err
			doneChan <- true
			return
		}

		defer rows.Close()

		for rows.Next() {
			cat := models.Category{}
			err := rows.Scan(&cat.ID, &cat.Name, &cat.CreatedAt, &cat.UpdatedAt, &cat.Version)

			if err != nil {
				errChan <- err
				doneChan <- true
				return
			}

			catChan <- &cat
		}

		if rows.Err() != nil {
			errChan <- rows.Err()
			doneChan <- true
			return
		}

		doneChan <- true

	}()

	cats := []*models.Category{}
	var err error

	for {
		select {
		case cat := <-catChan:
			cats = append(cats, cat)
		case e := <-errChan:
			err = e
		case <-doneChan:
			return cats, err
		}
	}

}
