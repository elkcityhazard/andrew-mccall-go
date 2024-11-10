package sqldbconn

import (
	"context"
	"time"

	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
)

func (sdc *SQLDbConn) InsertMessage(msg *models.ContactMsg) (int64, error) {
	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*10)

	defer cancel()

	idChan := make(chan int64)
	errChan := make(chan error)

	sdc.app.WG.Add(1)

	go func() {
		defer sdc.app.WG.Done()
		defer close(idChan)
		defer close(errChan)

		stmt := `
		INSERT INTO messages (email, message,created_at,version) VALUES(?,?,?,?);
		`

		args := []any{msg.Email, msg.Message, msg.CreatedAt, msg.Version}

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
	case err := <-errChan:
		return 0, err
	case id := <-idChan:
		return id, nil
	}

}
