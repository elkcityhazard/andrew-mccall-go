package handlers

import (
	"encoding/json"
	"net/http"
)

func (hr *HandlerRepo) HandleGetResume(w http.ResponseWriter, r *http.Request) {

	userID := hr.app.SessionManager.GetInt64(r.Context(), "id")

	if userID == 0 {
		userID = 1
	}

	resume, err := hr.conn.GetResumeById(userID)

	if err != nil {
		returnErr(w, err)
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
		returnErr(w, err)
		return
	}

	resume.ContactDetail = contactDetails

	socialList, err := hr.conn.GetResumeSocialMedia(resume.ID)

	if err != nil {
		returnErr(w, err)
		return
	}
	resume.SocialMediaList = socialList

	awardList, err := hr.conn.GetAwardItems(resume.ID)

	if err != nil {
		returnErr(w, err)
		return
	}
	resume.AwardsList = awardList

	skillList, err := hr.conn.GetSkillItems(resume.ID)

	if err != nil {
		returnErr(w, err)
		return
	}

	resume.SkillList = skillList

	employmentList, err := hr.conn.GetEmploymentList(resume.ID)

	if err != nil {
		returnErr(w, err)
		return
	}

	resume.EmploymentList = employmentList

	educationList, err := hr.conn.GetEducationList(resume.ID)

	if err != nil {
		returnErr(w, err)
		return
	}

	resume.EducationList = educationList

	refList, err := hr.conn.GetReferenceList(resume.ID)

	if err != nil {
		returnErr(w, err)
		return
	}

	resume.ReferenceList = refList

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(resume)

	if err != nil {
		returnErr(w, err)
		return
	}

}
