package sqldbconn

import (
	"context"
	"fmt"
	"reflect"
	"strings"
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
			errorChan <- err
			err = tx.Rollback()
			if err != nil {
				errorChan <- err
			}
			return
		}

		resumeID, err := result.LastInsertId()

		if err != nil {
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
			(resume_id,content)
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

		_, err = tx.ExecContext(ctx, stmt, contactDetailsArgs...)

		if err != nil {
			errorChan <- err
			err = tx.Rollback()
			if err != nil {
				errorChan <- err
			}
			return
		}

		//	skill list group

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

		stmt = `
		insert into
		skill_list_items
		()
		VALUES 
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

		// Employment History Group

		// Employment History Items

		// Social Media Group

		// Social Media Items

		// Education List Group

		// Education List Items

		// Award List Group

		// Award List Items

		// Reference List Group

		// Reference List Items

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
func buildMultiRowInsert(numRows, numValsPerRow int) string {

	var out []string

	for i := 0; i < numRows; i++ {
		s := make([]string, 0, numValsPerRow) // null val slice with length zero but capacity of numValsPerRow
		for j := 0; j < numValsPerRow; j++ {
			s[j] = "?"
		}
		var v = strings.Join(s, ", ")
		out = append(out, fmt.Sprintf("(%s)", v))
	}

	return strings.Join(out, ",")

}


func buildArgsMultiRowInsert(items []any, fieldNames...string) []any {
	vals := []any{}

	for i,v := range items {
		v := reflect.ValueOf(v)

		fields := reflect.VisibleFields(v)
	} 

	}

}
