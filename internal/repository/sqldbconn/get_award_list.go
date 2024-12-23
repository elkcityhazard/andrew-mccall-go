package sqldbconn

import (
	"context"
	"time"

	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
)

func (sdc *SQLDbConn) GetAwardItems(resumeID int64) (*models.AwardsList, error) {
	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*15)

	defer cancel()

	awardListChan := make(chan *models.AwardsList)
	errChan := make(chan error)

	sdc.app.WG.Add(1)

	go func() {
		defer sdc.app.WG.Done()
		defer close(awardListChan)
		defer close(errChan)

		stmt := `
		select
			award_lists.id,
			award_lists.resume_id,
			award_lists.created_at,
			award_lists.updated_at,
			award_lists.version,
			award_items.id,
			award_items.award_list_id,
			award_items.title,
			award_items.org_name,
			award_items.received_year,
			award_items.content,
			award_items.created_at,
			award_items.updated_at,
			award_items.version
		FROM award_lists
		INNER JOIN
			award_items ON award_items.award_list_id =award_list_id 
		WHERE award_lists.resume_id = ?
		ORDER BY award_items.created_at ASC;
		`

		args := []any{resumeID}

		rows, err := sdc.conn.QueryContext(ctx, stmt, args...)
		if err != nil {
			errChan <- err
			return
		}

		defer rows.Close()

		awList := models.AwardsList{}

		for rows.Next() {

			awItem := models.AwardItem{}

			err := rows.Scan(
				&awList.ID,
				&awList.ResumeID,
				&awList.CreatedAt,
				&awList.UpdatedAt,
				&awList.Version,
				&awItem.ID,
				&awItem.AwardListID,
				&awItem.Title,
				&awItem.OrganizationName,
				&awItem.Year,
				&awItem.Content,
				&awItem.CreatedAt,
				&awItem.UpdatedAt,
				&awItem.Version,
			)

			if err != nil {
				errChan <- err
				return
			}

			awList.Awards = append(awList.Awards, &awItem)

		}

		awardListChan <- &awList
	}()

	select {
	case err := <-errChan:
		return nil, err
	case awList := <-awardListChan:
		return awList, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
