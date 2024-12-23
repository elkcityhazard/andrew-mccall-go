package sqldbconn

import (
	"context"
	"time"

	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
)

func (sdc *SQLDbConn) GetResumeObjective(resumeID int64) (*models.Objective, error) {
	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*15)

	defer cancel()

	objectiveChan := make(chan *models.Objective)
	errorChan := make(chan error)
	sdc.app.WG.Add(1)
	go func() {
		defer close(objectiveChan)
		defer close(errorChan)
		defer sdc.app.WG.Done()

		stmt := `
		select
			objectives.id as res_obj_id,
			objectives.resume_id as res_res_id,
			objectives.content as res_content,
			objectives.created_at as res_created_at,
			objectives.updated_at as res_updated_at,
			objectives.version as res_obj_version
			FROM objectives
			WHERE resume_id = ?
			ORDER BY created_AT ASC
		`

		args := []any{resumeID}

		row := sdc.conn.QueryRowContext(ctx, stmt, args...)

		obj := models.Objective{}

		err := row.Scan(&obj.ID, &obj.ResumeID, &obj.Content, &obj.CreatedAt, &obj.UpdatedAt, &obj.Version)

		if err != nil {
			errorChan <- err
			return
		}

		objectiveChan <- &obj

	}()

	select {
	case err := <-errorChan:
		return nil, err
	case objective := <-objectiveChan:
		return objective, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
