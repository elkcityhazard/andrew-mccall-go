package handlers

import (
	"net/http"

	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
	"github.com/elkcityhazard/andrew-mccall-go/internal/render"
)

func (hr *HandlerRepo) HandleGetResume(w http.ResponseWriter, r *http.Request) {

	userID := hr.app.SessionManager.GetInt64(r.Context(), "id")

	if userID == 0 {
		userID = 1
	}

	resume, err := hr.conn.GetResumeById(userID)

	if err != nil {
		render.RenderTemplate(w, r, "404.gohtml", &models.TemplateData{})
		return
	}

	objective, err := hr.conn.GetResumeObjective(resume.ID)

	if err != nil {
		returnErr(w, err)
		return
	}

	resume.Objective = objective

	contactDetails, err := hr.conn.GetResumeContactDetails(resume.ID)

	if err != nil {
		contactDetails = models.NewContactDetail()
	}

	resume.ContactDetail = contactDetails

	socialList, err := hr.conn.GetResumeSocialMedia(resume.ID)

	if err != nil {
		socialList = models.NewSocialMediaList()
		return
	}
	resume.SocialMediaList = socialList

	awardList, err := hr.conn.GetAwardItems(resume.ID)

	if err != nil {
		awardList = models.NewAwardsList()
	}
	resume.AwardsList = awardList

	skillList, err := hr.conn.GetSkillItems(resume.ID)

	if err != nil {
		skillList = models.NewSkillList()
	}

	resume.SkillList = skillList

	employmentList, err := hr.conn.GetEmploymentList(resume.ID)

	if err != nil {
		employmentList = models.NewEmploymentList()
	}

	resume.EmploymentList = employmentList

	educationList, err := hr.conn.GetEducationList(resume.ID)

	if err != nil {
		educationList = models.NewEducationList()
	}

	resume.EducationList = educationList

	refList, err := hr.conn.GetReferenceList(resume.ID)

	if err != nil {
		refList = models.NewReferenceList()
	}

	resume.ReferenceList = refList

	data := make(map[string]any)
	stringMap := make(map[string]string)

	data["Resume"] = resume

	stringMap["PageTitle"] = "Resume"

	render.RenderTemplate(w, r, "resume.gohtml", &models.TemplateData{
		StringMap: stringMap,
		Data:      data,
		Form:      nil,
	})

}
