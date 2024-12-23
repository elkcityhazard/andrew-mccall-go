package sqldbconn

import (
	"context"
	"time"

	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
)

// GetSkillItems accepts a resumeID of type int64 and returns *models.SkillList or an error
// it performs a join query to get all unique skills that belong to a group
// using the resumeID to return the skill list items that belong to that resume
func (sdc *SQLDbConn) GetSkillItems(resumeID int64) (*models.SkillList, error) {
	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*15)

	defer cancel()

	skillListChan := make(chan *models.SkillList)
	errChan := make(chan error)

	sdc.app.WG.Add(1)

	go func() {
		defer sdc.app.WG.Done()
		defer close(skillListChan)
		defer close(errChan)

		stmt := `
		select
	skill_lists.id,
	skill_lists.resume_id,
	skill_lists.title,
	skill_lists.created_at,
	skill_lists.updated_at,
	skill_lists.version,
	skill_list_items.id,
	skill_list_items.skill_lists_id,
	skill_list_items.title,
	skill_list_items.content,
	skill_list_items.duration,
	skill_list_items.created_at,
	skill_list_items.updated_at,
	skill_list_items.version
	FROM skill_lists
	INNER JOIN skill_list_items ON skill_list_items.skill_lists_id = skill_lists.id
	WHERE skill_lists.resume_id = ?
	ORDER BY skill_list_items.created_at asc;
		`

		args := []any{resumeID}

		rows, err := sdc.conn.QueryContext(ctx, stmt, args...)

		if err != nil {
			errChan <- err
			return
		}

		skList := models.SkillList{}

		defer rows.Close()

		for rows.Next() {
			skItem := models.SkillListItem{}

			err := rows.Scan(
				&skList.ID,
				&skList.ResumeID,
				&skList.Title,
				&skList.CreatedAt,
				&skList.UpdatedAt,
				&skList.Version,
				&skItem.ID,
				&skItem.SKillListID,
				&skItem.Title,
				&skItem.Content,
				&skItem.Duration,
				&skItem.CreatedAt,
				&skItem.UpdatedAt,
				&skItem.Version,
			)

			if err != nil {
				errChan <- err
				return
			}

			skList.Items = append(skList.Items, &skItem)

		}

		skillListChan <- &skList

	}()

	select {
	case err := <-errChan:
		return nil, err
	case skillList := <-skillListChan:
		return skillList, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}

}
