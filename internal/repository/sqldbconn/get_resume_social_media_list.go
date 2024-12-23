package sqldbconn

import (
	"context"
	"time"

	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
)

func (sdc *SQLDbConn) GetResumeSocialMedia(resumeID int64) (*models.SocialMediaList, error) {
	ctx, cancel := context.WithTimeout(sdc.app.Context, time.Second*15)

	defer cancel()

	sdc.app.WG.Add(1)
	errorChan := make(chan error)
	smListChan := make(chan *models.SocialMediaList)

	go func() {
		defer sdc.app.WG.Done()
		defer close(errorChan)
		defer close(smListChan)

		stmt := `
	select
		social_media_lists.id as sml_id,
		social_media_lists.resume_id as sml_resume_id,
		social_media_lists.created_at as sml_created_at,
		social_media_lists.updated_at as sml_updated_at,
		social_media_lists.version as sml_version,
		social_media_list_items.id as sml_item_id,
		social_media_list_items.social_media_lists_id as sml_item_list_id,
		social_media_list_items.company_name as sml_item_company_name,
		social_media_list_items.username as sml_item_username,
		social_media_list_items.web_address as sml_item_web_address,
		social_media_list_items.created_at as sml_item_created_at,
		social_media_list_items.updated_at as sml_item_updated_at,
		social_media_list_items.version as sml_item_version
	from social_media_lists
	left join social_media_list_items ON
		social_media_list_items.social_media_lists_id = social_media_lists.id
	WHERE
		social_media_lists.resume_id = ?;
		`

		args := []any{resumeID}

		rows, err := sdc.conn.QueryContext(ctx, stmt, args...)

		if err != nil {
			errorChan <- err
			return
		}

		defer rows.Close()

		smList := models.SocialMediaList{}

		for rows.Next() {

			listItem := models.SocialMediaListItems{}

			err := rows.Scan(
				&smList.ID,
				&smList.ResumeID,
				&smList.CreatedAt,
				&smList.UpdatedAt,
				&smList.Version,
				&listItem.ID,
				&listItem.SocialMediaListID,
				&listItem.CompanyName,
				&listItem.UserName,
				&listItem.WebAddress,
				&listItem.CreatedAt,
				&listItem.UpdatedAt,
				&listItem.Version,
			)

			if err != nil {
				errorChan <- err
				return
			}

			smList.SocialMediaListItems = append(smList.SocialMediaListItems, &listItem)

		}

		smListChan <- &smList

	}()

	select {
	case err := <-errorChan:
		return nil, err
	case smList := <-smListChan:
		return smList, nil
	}
}
