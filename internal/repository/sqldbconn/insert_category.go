package sqldbconn

import (
	"context"
	"time"

	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
)

func (sdc *SQLDbConn) InsertCategory(cat *models.Category) (int64, error) {
	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*15)

	defer cancel()

	idChan := make(chan int64)
	errChan := make(chan error)

	sdc.app.WG.Add(1)

	go func() {
		defer sdc.app.WG.Done()
		defer close(idChan)
		defer close(errChan)

		stmt := `INSERT INTO categories (name,created_at,updated_at,version) VALUES(?,?,?,?)`

		args := []any{cat.Name, cat.CreatedAt, cat.UpdatedAt, cat.Version}

		result, err := sdc.conn.ExecContext(ctx, stmt, args...)

		if err != nil {
			errChan <- err
			return
		}

		id, err := result.LastInsertId()

		if err != nil {
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
