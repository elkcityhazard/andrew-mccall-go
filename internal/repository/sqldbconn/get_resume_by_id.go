package sqldbconn

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
)

func (sdc *SQLDbConn) GetResumeById(userID int64) (*models.Resume, error) {
	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*15)

	defer cancel()

	resumeChan := make(chan *models.Resume)
	errorChan := make(chan error)

	sdc.app.WG.Add(1)
	go func() {
		defer sdc.app.WG.Done()
		defer close(resumeChan)
		defer close(errorChan)

		stmt := `
		select 
		resumes.id as resume_id,
		resumes.user_id as resume_user_id,
		resumes.job_title as resume_job_title,
		resumes.created_at as resume_created_at,
		resumes.updated_at as resume_updated_at,
		resumes.version as resume_version
		FROM resumes WHERE id = ?
		`
		args := []any{userID}

		row := sdc.conn.QueryRowContext(ctx, stmt, args...)

		res := models.Resume{}

		err := row.Scan(&res.ID, &res.UserID, &res.JobTitle, &res.CreatedAt, &res.UpdatedAt, &res.Version)

		if err != nil {
			if err == sql.ErrNoRows {
				err = errors.New("could not find a resume")
			}
			errorChan <- err
			return
		}
		resumeChan <- &res

	}()

	select {
	case err := <-errorChan:
		return nil, err
	case resume := <-resumeChan:
		return resume, nil
	case <-ctx.Done():
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		return nil, errors.New("query timedout")
	}

}
