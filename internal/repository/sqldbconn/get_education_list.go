package sqldbconn

import (
	"context"
	"time"

	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
)

func (sdc *SQLDbConn) GetEducationList(resumeID int64) (*models.EducationList, error) {
	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*15)

	defer cancel()
	sdc.app.WG.Add(1)
	eduListChan := make(chan *models.EducationList)
	errChan := make(chan error)

	go func() {
		defer sdc.app.WG.Done()
		defer close(eduListChan)
		defer close(errChan)

		stmt := `
		select
	education_lists.id,
	education_lists.resume_id,
	education_lists.created_at,
	education_lists.updated_at,
	education_lists.version,
	education_items.id,
	education_items.education_lists_id,
	education_items.name,
	education_items.degree_year,
	education_items.degree,
	education_items.address_1,
	education_items.address_2,
	education_items.city,
	education_items.state,
	education_items.zipcode,
	education_items.created_at,
	education_items.updated_at,
	education_items.version
FROM education_lists
INNER JOIN education_items
ON education_items.education_lists_id = education_lists.id
WHERE education_lists.resume_id = ?
ORDER BY education_items.created_at ASC;
		`

		args := []any{resumeID}

		rows, err := sdc.conn.QueryContext(ctx, stmt, args...)

		if err != nil {
			errChan <- err
			return
		}

		defer rows.Close()

		eduList := models.EducationList{}

		for rows.Next() {
			eduItem := models.EducationItem{}

			err := rows.Scan(
				&eduList.ID,
				&eduList.ResumeID,
				&eduList.CreatedAt,
				&eduList.UpdatedAt,
				&eduList.Version,
				&eduItem.ID,
				&eduItem.EducationListID,
				&eduItem.Name,
				&eduItem.DegreeYear,
				&eduItem.Degree,
				&eduItem.Address1,
				&eduItem.Address2,
				&eduItem.City,
				&eduItem.State,
				&eduItem.Zipcode,
				&eduItem.CreatedAt,
				&eduItem.UpdatedAt,
				&eduItem.Version,
			)

			if err != nil {
				errChan <- err
				return
			}

			eduList.Education = append(eduList.Education, &eduItem)

		}

		eduListChan <- &eduList
	}()

	select {
	case err := <-errChan:
		return nil, err
	case eduList := <-eduListChan:
		return eduList, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}

}
