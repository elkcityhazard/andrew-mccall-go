package sqldbconn

import (
	"context"
	"time"

	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
)

func (sdc *SQLDbConn) GetEmploymentList(resumeID int64) (*models.EmploymentList, error) {
	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*15)

	defer cancel()

	empListChan := make(chan *models.EmploymentList)
	errChan := make(chan error)
	sdc.app.WG.Add(1)

	go func() {
		defer sdc.app.WG.Done()
		defer close(empListChan)
		defer close(errChan)

		stmt := `
		select
			employment_lists.id,
			employment_lists.resume_id,
			employment_lists.created_at,
			employment_lists.updated_at,
			employment_lists.version,
			employment_list_items.id,
			employment_list_items.employment_lists_id,
			employment_list_items.title as employment_list_item_org_title,
			employment_list_items.date_from,
			employment_list_items.date_to,
			employment_list_items.job_title as employment_list_item_job_title,
			employment_list_items.summary,
			employment_list_items.created_at,
			employment_list_items.updated_at,
			employment_list_items.version
		FROM employment_lists
		INNER JOIN employment_list_items
		ON employment_list_items.employment_lists_id = employment_lists.id
		WHERE employment_lists.resume_id = ?
		ORDER BY employment_list_items.created_at ASC
			`

		args := []any{resumeID}

		rows, err := sdc.conn.QueryContext(ctx, stmt, args...)

		if err != nil {
			errChan <- err
			return
		}

		defer rows.Close()

		empList := models.EmploymentList{}

		for rows.Next() {

			empListItem := models.EmploymentListItem{}

			err := rows.Scan(
				&empList.ID,
				&empList.ResumeID,
				&empList.CreatedAt,
				&empList.UpdatedAt,
				&empList.Version,
				&empListItem.ID,
				&empListItem.EmploymentListID,
				&empListItem.Title,
				&empListItem.From,
				&empListItem.To,
				&empListItem.JobTitle,
				&empListItem.Summary,
				&empListItem.CreatedAt,
				&empListItem.UpdatedAt,
				&empListItem.Version,
			)

			if err != nil {
				errChan <- err
				return
			}

			empList.Employers = append(empList.Employers, &empListItem)
		}

		empListChan <- &empList

	}()

	select {
	case err := <-errChan:
		return nil, err
	case empList := <-empListChan:
		return empList, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}

}
