package sqldbconn

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
)

func (sdc *SQLDbConn) InsertResume(rme *models.Resume) (int64, error) {
	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*30)

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

		stmt := `
		INSERT INTO
		resumes
		(
			user_id,
			job_title,
			created_at,
			updated_at,
			version
		) VALUES(
		?,?,?,?,?	
		)
		`

		args := []any{rme.UserID, rme.JobTitle, rme.CreatedAt, rme.UpdatedAt, rme.Version}

		result, err := tx.ExecContext(ctx, stmt, args...)

		if err != nil {
			fmt.Println("resumes")
			errorChan <- err
			err = tx.Rollback()
			if err != nil {
				errorChan <- err
			}
			return
		}

		resumeID, err := result.LastInsertId()

		if err != nil {
			fmt.Println("resumes id")
			errorChan <- err
			err = tx.Rollback()
			if err != nil {
				errorChan <- err
			}
			return
		}

		// objective

		stmt = `
		INSERT INTO
			objectives
			(
				resume_id,
			content
		)
			VALUES(?,?);
		`

		args = []any{resumeID, rme.Objective.Content}

		_, err = tx.ExecContext(ctx, stmt, args...)

		if err != nil {
			errorChan <- err
			err = tx.Rollback()
			if err != nil {
				errorChan <- err
			}
			return
		}

		// contact details

		stmt = `
		insert into
			resume_contact_details
			(
				resume_id,
				first_name,
				last_name,
				address_1,
				address_2,
				city,
				state,
				zipcode,
				email, 
				phone_number,
				web_address
			)
			VALUES(?,?,?,?,?,?,?,?,?,?,?)
		`

		cd := rme.ContactDetail

		contactDetailsArgs := []any{
			resumeID,
			cd.Firstname,
			cd.Lastname,
			cd.AddressLine1,
			cd.AddressLine2,
			cd.City,
			cd.State,
			cd.Zipcode,
			cd.Email,
			cd.PhoneNumber,
			cd.WebAddress,
		}

		result, err = tx.ExecContext(ctx, stmt, contactDetailsArgs...)

		if err != nil {
			errorChan <- err
			err = tx.Rollback()
			if err != nil {
				errorChan <- err
			}
			return
		}

		contactDetailID, err := result.LastInsertId()
		if err != nil {
			errorChan <- err
			err = tx.Rollback()
			if err != nil {
				errorChan <- err
			}
			return
		}

		rme.ContactDetail.ID = contactDetailID

		//	skill list group

		if len(rme.SkillList.Items) > 0 {

			stmt = `
		insert into
		skill_lists
		(resume_id)
		values(?)
		`

			args = []any{resumeID}

			result, err = tx.ExecContext(ctx, stmt, args...)
			if err != nil {
				errorChan <- err
				err = tx.Rollback()
				if err != nil {
					errorChan <- err
				}
				return
			}

			skillListGroupID, err := result.LastInsertId()
			if err != nil {
				errorChan <- err
				err = tx.Rollback()
				if err != nil {
					errorChan <- err
				}
				return
			}

			rme.SkillList.ID = skillListGroupID

			// Skill List Items

			for i := range rme.SkillList.Items {
				rme.SkillList.Items[i].SKillListID = skillListGroupID
			}

			stmt = `
		insert into
		skill_list_items
		(
			skill_lists_id,
			title,
			content,
			duration
		)
		values
		`

			// build multi row insert

			skillListItemRows := buildMultiRowInsert(len(rme.SkillList.Items), 4)

			stmt += skillListItemRows

			// build args

			args = []any{}

			for i := range rme.SkillList.Items {
				var item = rme.SkillList.Items[i]
				args = append(args, skillListGroupID, item.Title, item.Content, item.Duration)
			}

			result, err = tx.ExecContext(ctx, stmt, args...)
			if err != nil {
				errorChan <- err
				err = tx.Rollback()
				if err != nil {
					errorChan <- err
				}
				return
			}

			lastSkillListItemID, err := result.LastInsertId()
			if err != nil {
				errorChan <- err
				err = tx.Rollback()
				if err != nil {
					errorChan <- err
				}
				return
			}

			for i := range rme.SkillList.Items {
				rme.SkillList.Items[i].ID = lastSkillListItemID
			}
		}

		// Employment History Group

		if len(rme.EmploymentList.Employers) > 0 {

			stmt = `
		insert into
		employment_lists
		(
			resume_id
		)
		values (?)
		`

			args = []any{resumeID}

			result, err = tx.ExecContext(ctx, stmt, args...)

			if err != nil {
				errorChan <- err
				err = tx.Rollback()
				if err != nil {
					errorChan <- err
				}
				return
			}

			employmentListID, err := result.LastInsertId()

			if err != nil {
				errorChan <- err
				err = tx.Rollback()
				if err != nil {
					errorChan <- err
				}
				return
			}

			rme.EmploymentList.ID = employmentListID

			// Employment History Items

			stmt = `
		insert into
			employment_list_items
			(
				employment_lists_id,
				title,
				date_from,
				date_to,
				job_title,
				summary
			)
			values
		`

			stmt += buildMultiRowInsert(len(rme.EmploymentList.Employers), 6)

			args = []any{}

			for i := range rme.EmploymentList.Employers {
				rme.EmploymentList.Employers[i].EmploymentListID = employmentListID
				var item = rme.EmploymentList.Employers[i]
				args = append(args, item.EmploymentListID, item.Title, item.From, item.To, item.JobTitle, item.Summary)
			}

			result, err = tx.ExecContext(ctx, stmt, args...)

			if err != nil {
				errorChan <- err
				err = tx.Rollback()
				if err != nil {
					errorChan <- err
				}
				return
			}
		}

		// Social Media Group

		if len(rme.SocialMediaList.SocialMediaListItems) > 0 {

			stmt = `
		insert into
			social_media_lists
			(resume_id)
			values (?)
		`

			args = []any{resumeID}

			result, err = tx.ExecContext(ctx, stmt, args...)

			if err != nil {
				errorChan <- err
				err = tx.Rollback()
				if err != nil {
					errorChan <- err
				}
				return
			}

			socialMediaListID, err := result.LastInsertId()
			if err != nil {
				errorChan <- err
				err = tx.Rollback()
				if err != nil {
					errorChan <- err
				}
				return
			}

			rme.SocialMediaList.ID = socialMediaListID

			for i := range rme.SocialMediaList.SocialMediaListItems {
				rme.SocialMediaList.SocialMediaListItems[i].SocialMediaListID = socialMediaListID
			}

			// Social Media Items

			stmt = `
		insert into
			social_media_list_items
			(
				social_media_lists_id,
				company_name,
				username,
				web_address
			)
			values
		`

			smItems := rme.SocialMediaList.SocialMediaListItems

			stmt += buildMultiRowInsert(len(smItems), 4)

			args = []any{}

			for i := 0; i < len(rme.SocialMediaList.SocialMediaListItems); i++ {
				var item = rme.SocialMediaList.SocialMediaListItems[i]
				args = append(args, item.SocialMediaListID, item.CompanyName, item.UserName, item.WebAddress)
			}

			result, err = tx.ExecContext(ctx, stmt, args...)
			if err != nil {
				errorChan <- err
				err = tx.Rollback()
				if err != nil {
					errorChan <- err
				}
				return
			}

			lastSMItemID, err := result.LastInsertId()
			if err != nil {
				errorChan <- err
				err = tx.Rollback()
				if err != nil {
					errorChan <- err
				}
				return
			}

			for i := range rme.SocialMediaList.SocialMediaListItems {
				rme.SocialMediaList.SocialMediaListItems[i].ID = lastSMItemID
			}
		}

		// Education List Group

		if len(rme.EducationList.Education) > 0 {

			stmt = `
		insert into
			education_lists
			(
				resume_id
			)
			values (?)
		`

			args = []any{}

			args = append(args, resumeID)

			result, err = tx.ExecContext(ctx, stmt, args...)
			if err != nil {
				errorChan <- err
				err = tx.Rollback()
				if err != nil {
					errorChan <- err
				}
				return
			}

			educationListID, err := result.LastInsertId()
			if err != nil {
				errorChan <- err
				err = tx.Rollback()
				if err != nil {
					errorChan <- err
				}
				return
			}

			rme.EducationList.ID = educationListID

			// Education List Items

			stmt = `
		insert into
			education_items
			(
				education_lists_id,
				name,
				degree_year,
				degree,
				address_1,
				address_2,
				city,
				state,
				zipcode
			)
			values 
		`

			stmt += buildMultiRowInsert(len(rme.EducationList.Education), 9)

			args = []any{}

			for i := range rme.EducationList.Education {
				rme.EducationList.Education[i].EducationListID = educationListID
				item := rme.EducationList.Education[i]
				args = append(args, item.EducationListID, item.Name, item.DegreeYear, item.Degree, item.Address1, item.Address2, item.City, item.State, item.Zipcode)
			}

			result, err = tx.ExecContext(ctx, stmt, args...)
			if err != nil {
				errorChan <- err
				err = tx.Rollback()
				if err != nil {
					errorChan <- err
				}
				return
			}
		}

		// Award List Group

		if len(rme.AwardsList.Awards) > 0 {

			stmt = `
		insert into
			award_lists
			(resume_id)
			values (?)
		`

			args = []any{}
			args = append(args, resumeID)

			result, err = tx.ExecContext(ctx, stmt, args...)
			if err != nil {
				errorChan <- err
				err = tx.Rollback()
				if err != nil {
					errorChan <- err
				}
				return
			}

			awardListID, err := result.LastInsertId()
			if err != nil {
				errorChan <- err
				err = tx.Rollback()
				if err != nil {
					errorChan <- err
				}
				return
			}

			// Award List Items

			stmt = `
		insert into
			award_items
			(
				award_list_id,
				title,
				org_name,
				received_year,
				content
			)
			values 
		`

			stmt += buildMultiRowInsert(len(rme.AwardsList.Awards), 5)

			args = []any{}

			for i := range rme.AwardsList.Awards {
				rme.AwardsList.Awards[i].AwardListID = awardListID
				item := rme.AwardsList.Awards[i]

				args = append(args, item.AwardListID, item.Title, item.OrganizationName, item.Year, item.Content)
			}

			result, err = tx.ExecContext(ctx, stmt, args...)
			if err != nil {
				errorChan <- err
				err = tx.Rollback()
				if err != nil {
					errorChan <- err
				}
				return
			}
		}

		// Reference List Group
		if len(rme.ReferenceList.ReferenceList) > 0 {
			stmt = `
		insert into
			reference_lists
			(resume_id)
			values (?)
		`
			args = []any{resumeID}

			result, err = tx.ExecContext(ctx, stmt, args...)
			if err != nil {
				errorChan <- err
				err = tx.Rollback()
				if err != nil {
					errorChan <- err
				}
				return
			}

			referenceListID, err := result.LastInsertId()
			if err != nil {
				errorChan <- err
				err = tx.Rollback()
				if err != nil {
					errorChan <- err
				}
				return
			}

			rme.ReferenceList.ID = referenceListID

			for i := range rme.ReferenceList.ReferenceList {
				rme.ReferenceList.ReferenceList[i].ReferenceListID = referenceListID
			}

			// Reference List Items

			stmt = `
		insert into
			reference_items
			(
				ref_list_id,
				first_name,
				last_name,
				email,
				phone_number,
				job_title,
				organization,
				type,
				address_1,
				address_2,
				city,
				state,
				zipcode,
				content
			)
			values  
		`
			stmt += buildMultiRowInsert(len(rme.ReferenceList.ReferenceList), 14)

			args = []any{}

			for i := range rme.ReferenceList.ReferenceList {
				item := rme.ReferenceList.ReferenceList[i]
				rme.ReferenceList.ReferenceList[i].ReferenceListID = referenceListID
				args = append(args,
					item.ReferenceListID,
					item.FirstName,
					item.LastName,
					item.Email,
					item.PhoneNumber,
					item.JobTitle,
					item.Organization,
					item.Type,
					item.Address1,
					item.Address2,
					item.City,
					item.State,
					item.Zipcode,
					item.Content)
			}

			result, err = tx.ExecContext(ctx, stmt, args...)
			if err != nil {
				errorChan <- err
				err = tx.Rollback()
				if err != nil {
					errorChan <- err
				}
				return
			}
		}

		// commit
		err = tx.Commit()
		if err != nil {
			err = tx.Rollback()
			errorChan <- err
			return
		}

		idChan <- resumeID

	}()

	select {
	case id := <-idChan:
		rme.ID = id
		updateResumeID(rme, id)
		return id, nil
	case err := <-errorChan:
		return 0, err
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

func updateResumeID(resume *models.Resume, resumeID int64) {
	resume.Objective.ResumeID = resumeID
	resume.ContactDetail.ResumeID = resumeID
	resume.SkillList.ResumeID = resumeID
	resume.EmploymentList.ResumeID = resumeID
	resume.SocialMediaList.ResumeID = resumeID
	resume.EducationList.ResumeID = resumeID
	resume.AwardsList.ResumeID = resumeID
	resume.ReferenceList.ResumeID = resumeID
}

// buildMultiRowInsert creates an empty slice with the capacity of
// numValsPerRow and returns the joined string for db inserts
// append after VALUES in insert statement
// this is used to build prepared statements
func buildMultiRowInsert(numRows, numValsPerRow int) string {
	if numRows < 1 || numValsPerRow < 1 {
		return "()"
	}

	var out []string

	for i := 0; i < numRows; i++ {
		s := make([]string, numValsPerRow) // null val slice with length zero but capacity of numValsPerRow
		for j := 0; j < numValsPerRow; j++ {
			s[j] = "?"
		}
		var v = strings.Join(s, ", ")
		out = append(out, fmt.Sprintf("(%s)", v))
	}

	return strings.Join(out, ",")

}
