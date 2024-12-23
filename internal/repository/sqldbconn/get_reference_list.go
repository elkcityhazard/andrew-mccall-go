package sqldbconn

import (
	"context"
	"time"

	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
)

func (sdc *SQLDbConn) GetReferenceList(resumeID int64) (*models.ReferenceList, error) {
	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*15)

	defer cancel()

	sdc.app.WG.Add(1)
	refListChan := make(chan *models.ReferenceList)
	errChan := make(chan error)

	go func() {
		defer sdc.app.WG.Done()
		defer close(refListChan)
		defer close(errChan)

		stmt := `
		SELECT 
 reference_lists.id,
 reference_lists.resume_id,
 reference_lists.created_at,
 reference_lists.updated_at,
 reference_lists.version,
 reference_items.id,
 reference_items.ref_list_id,
 reference_items.first_name,
 reference_items.last_name,
 reference_items.email,
 reference_items.phone_number,
 reference_items.job_title,
 reference_items.organization,
 reference_items.type,
 reference_items.address_1,
 reference_items.address_2,
 reference_items.city,
 reference_items.state,
 reference_items.zipcode,
 reference_items.content,
 reference_items.created_at,
 reference_items.updated_at,
 reference_items.version
 FROM reference_lists
 INNER JOIN
 reference_items
 ON reference_items.ref_list_id = reference_lists.id
 WHERE reference_lists.resume_id = ?
 ORDER BY reference_items.created_at ASC;
		`

		args := []any{resumeID}

		rows, err := sdc.conn.QueryContext(ctx, stmt, args...)

		if err != nil {
			errChan <- err
			return
		}

		defer rows.Close()

		refList := models.ReferenceList{}

		for rows.Next() {

			refItem := models.ReferenceItem{}

			err := rows.Scan(
				&refList.ID,
				&refList.ResumeID,
				&refList.CreatedAt,
				&refList.UpdatedAt,
				&refList.Version,
				&refItem.ID,
				&refItem.ReferenceListID,
				&refItem.FirstName,
				&refItem.LastName,
				&refItem.Email,
				&refItem.PhoneNumber,
				&refItem.JobTitle,
				&refItem.Organization,
				&refItem.Type,
				&refItem.Address1,
				&refItem.Address2,
				&refItem.City,
				&refItem.State,
				&refItem.Zipcode,
				&refItem.Content,
				&refItem.CreatedAt,
				&refItem.UpdatedAt,
				&refItem.Version,
			)

			if err != nil {
				errChan <- err
				return
			}

			refList.ReferenceList = append(refList.ReferenceList, &refItem)

		}

		refListChan <- &refList

	}()

	select {
	case err := <-errChan:
		return nil, err
	case refList := <-refListChan:
		return refList, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}

}
