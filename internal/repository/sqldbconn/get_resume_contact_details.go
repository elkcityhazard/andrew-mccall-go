package sqldbconn

import (
	"context"
	"time"

	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
)

func (sdc *SQLDbConn) GetResumeContactDetails(resumeID int64) (*models.ContactDetail, error) {
	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*15)

	defer cancel()

	contactDetailChan := make(chan *models.ContactDetail)
	errorChan := make(chan error)
	sdc.app.WG.Add(1)

	go func() {
		defer sdc.app.WG.Done()
		defer close(contactDetailChan)
		defer close(errorChan)

		stmt := `
		select 
			resume_contact_details.id as res_cd_id, 
			resume_contact_details.resume_id as res_cd_res_id,
			resume_contact_details.first_name as res_cd_firstname,
			resume_contact_details.last_name as res_cd_lastname,
			resume_contact_details.address_1 as res_cd_address1,
			resume_contact_details.address_2 as res_cd_address2,
			resume_contact_details.city as res_cd_city,
			resume_contact_details.state as res_cd_state,
			resume_contact_details.zipcode as res_cd_zipcode,
			resume_contact_details.email as res_cd_email,
			resume_contact_details.phone_number as res_cd_phone,
			resume_contact_details.web_address as res_cd_web_address,
			resume_contact_details.created_at as res_cd_created_at,
			resume_contact_details.updated_at as res_cd_updated_at,
			resume_contact_details.version as res_cd_version
			from resume_contact_details
		WHERE resume_contact_details.resume_id = ?
		ORDER BY 
			resume_contact_details.created_at ASC;
		`

		args := []any{resumeID}

		cd := models.ContactDetail{}

		err := sdc.conn.QueryRowContext(ctx, stmt, args...).Scan(
			&cd.ID,
			&cd.ResumeID,
			&cd.Firstname,
			&cd.Lastname,
			&cd.AddressLine1,
			&cd.AddressLine2,
			&cd.City,
			&cd.State,
			&cd.Zipcode,
			&cd.Email,
			&cd.PhoneNumber,
			&cd.WebAddress,
			&cd.CreatedAt,
			&cd.UpdatedAt,
			&cd.Version)

		if err != nil {
			errorChan <- err
			return
		}

		contactDetailChan <- &cd

	}()

	return <-contactDetailChan, <-errorChan

}
