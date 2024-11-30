package sqldbconn

import (
	"context"
	"time"

	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
)

func (sdc *SQLDbConn) InsertResume(rme *models.Resume) (int64, error) {
	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*15)

	defer cancel()

	idChan := make(chan int64)
	errorChan := make(chan error)

	sdc.app.WG.Add(1)

	go func() {
		defer sdc.app.WG.Done()
		defer close(idChan)
		defer close(errorChan)

		tx, err := sdc.conn.BeginTx(ctx, nil)

		if err != nil {
			err = tx.Rollback()
			errorChan <- err
			return
		}

		// commit
		err = tx.Commit()
		if err != nil {
			err = tx.Rollback()
			errorChan <- err
			return
		}
	}()

	select {
	case id := <-idChan:
		return id, nil
	case err := <-errorChan:
		return 0, err
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}
