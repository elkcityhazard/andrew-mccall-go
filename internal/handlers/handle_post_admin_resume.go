package handlers

import (
	"encoding/json"
	"html"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elkcityhazard/andrew-mccall-go/internal/forms"
	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
	"github.com/elkcityhazard/andrew-mccall-go/internal/render"
)

func (hr *HandlerRepo) HandlePostAdminResume(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		returnErr(w, err)
		return
	}
	resumeForm := forms.New(r.Form)

	// required fields

	resumeForm.Required(
		"resume_job_title",
		"resume_contact-details_firstname",
		"resume_contact-details_lastname",
		"resume_contact-details_address1",
		"resume_contact-details_city",
		"resume_contact-details_state",
		"resume_contact-details_zipcode",
		"resume_contact-details_email",
		"resume_contact-details_phone",
	)

	if !resumeForm.IsEmail(resumeForm.Get("resume_contact-details_email")) {
		resumeForm.Errors.Add("resume_contact-details_email", "invalid email address")
	}

	if !resumeForm.Valid() {
		data, stringMap, intMap, _ := hr.CreateEmptyTemplatePayload()

		stringMap["PageTitle"] = "Resume"
		data["UserID"] = hr.app.SessionManager.GetInt64(r.Context(), "id")
		render.RenderTemplate(w, r, "admin-resume.gohtml", &models.TemplateData{
			StringMap: stringMap,
			Data:      data,
			IntMap:    intMap,
			Form:      resumeForm,
		})
		return
	}

	// New resume
	var res = models.NewResume()
	res.UserID = hr.app.SessionManager.GetInt64(r.Context(), "id")
	res.JobTitle = html.EscapeString(resumeForm.Get("resume_job_title"))

	// contact details
	res.ContactDetail = parseFormIntoContactDetails(&resumeForm.Values)

	// Objective

	var resObjective = models.NewObjective()
	resObjective.Content = html.EscapeString(resumeForm.Get("resume_objective"))
	res.Objective = resObjective

	res.SocialMediaList = parseSocialList(resumeForm)

	res.SkillList = parseFormIntoSkillList(resumeForm)

	res.AwardsList = parseAwardsList(resumeForm)

	res.EmploymentList = parseEmploymentList(resumeForm)

	res.EducationList = parseEducationList(resumeForm)

	res.ReferenceList = parseReferenceList(resumeForm)

	resumeID, err := hr.conn.InsertResume(res)

	if err != nil {
		returnErr(w, err)
		return
	}

	res.ID = resumeID

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(res)

}

func escapeHTML(s string) string {

	if len(s) < 1 {
		return ""
	}
	s = strings.TrimSpace(s)
	return html.EscapeString(s)
}

func parseStringToInt(s string) int {
	val, err := strconv.Atoi(s)

	if err != nil {
		return 0
	}
	return val
}

// getFormKey is a helper function to make sure the key exists based on the index,
// if it does not, we return an empty string
func getFormKey(data map[string][]string, key string, index int) string {
	if values, ok := data[key]; ok {
		if index >= 0 && index < len(values) {
			return values[index]
		}
	}
	return ""
}

// parseReferenceList takes in *forms.Form and parses the values into an Reference List and returns the list,
// sanitizing any strings along the way
func parseReferenceList(f *forms.Form) *models.ReferenceList {
	var list = models.NewReferenceList()
	if !hasLength(f.Values["resume_reference-list_firstname"]) {
		return list
	}

	for i := range f.Values["resume_reference-list_firstname"] {
		var item = models.NewReferenceItem()

		item.FirstName = escapeHTML(getFormKey(f.Values, "resume_reference-list_firstname", i))
		item.LastName = escapeHTML(getFormKey(f.Values, "resume_reference-list_lastname", i))
		item.Email = escapeHTML(getFormKey(f.Values, "resume_reference-list_email", i))
		item.PhoneNumber = escapeHTML(getFormKey(f.Values, "resume_reference-list_phonenumber", i))
		item.JobTitle = escapeHTML(getFormKey(f.Values, "resume_reference-list_jobtitle", i))
		item.Organization = escapeHTML(getFormKey(f.Values, "resume_reference-list_organization", i))
		item.Type = escapeHTML(getFormKey(f.Values, "resume_reference-list_type", i))
		item.Address1 = escapeHTML(getFormKey(f.Values, "resume_reference-list_address1", i))
		item.Address2 = escapeHTML(getFormKey(f.Values, "resume_reference-list_address2", i))
		item.City = escapeHTML(getFormKey(f.Values, "resume_reference-list_city", i))
		item.State = escapeHTML(getFormKey(f.Values, "resume_reference-list_state", i))
		item.Zipcode = escapeHTML(getFormKey(f.Values, "resume_reference-list_zipcode", i))
		item.Content = escapeHTML(getFormKey(f.Values, "resume_reference-list_content", i))

		list.ReferenceList = append(list.ReferenceList, item)
	}

	return list
}

// parseAwardsList takes in *forms.Form and parsed the values into an AwardList struct and returns it, sanitizing any strings
func parseAwardsList(f *forms.Form) *models.AwardsList {
	var list = models.NewAwardsList()
	if !hasLength(f.Values["resume_award-list_title"]) {
		return list
	}

	for i := range f.Values["resume_award-list_title"] {
		var item = models.NewAwardItem()
		item.Title = escapeHTML(getFormKey(f.Values, "resume_award-list_title", i))
		item.OrganizationName = escapeHTML(getFormKey(f.Values, "resume_award-list_name", i))
		item.Year = parseStringToInt(escapeHTML(getFormKey(f.Values, "resume_award-list_year", i)))
		item.Content = escapeHTML(getFormKey(f.Values, "resume_award-list_content", i))

		list.Awards = append(list.Awards, item)
	}

	return list
}

// parseEducationList takes in url values and creates a new education list, santizing along the way

func parseEducationList(f *forms.Form) *models.EducationList {
	var list = models.NewEducationList()
	if !hasLength(f.Values["resume_education-list_name"]) {
		return list
	}

	for i := range f.Values["resume_education-list_name"] {
		var item = models.NewEducationItem()

		item.Name = escapeHTML(getFormKey(f.Values, "resume_education-list_name", i))
		item.DegreeYear = parseStringToInt(escapeHTML(getFormKey(f.Values, "resume_education-list_degreeyear", i)))
		item.Degree = escapeHTML(getFormKey(f.Values, "resume_education-list_degree", i))
		item.Address1 = escapeHTML(getFormKey(f.Values, "resume_education-list_address1", i))
		item.Address2 = escapeHTML(getFormKey(f.Values, "resume_education-list_address2", i))
		item.City = escapeHTML(getFormKey(f.Values, "resume_education-list_city", i))
		item.State = escapeHTML(getFormKey(f.Values, "resume_education-list_state", i))
		item.Zipcode = escapeHTML(getFormKey(f.Values, "resume_education-list_zipcode", i))

		list.Education = append(list.Education, item)
	}

	return list
}

//	parseEmploymentList takes in url values and creates an new past employment list, sanitizing values along the way

func parseEmploymentList(f *forms.Form) *models.EmploymentList {
	var el = models.NewEmploymentList()
	if !hasLength(f.Values["resume_employment-list_title"]) {
		return el
	}

	for i := range f.Values["resume_employment-list_title"] {
		var item = models.NewEmploymentListItem()

		item.Title = escapeHTML(getFormKey(f.Values, "resume_employment-list_title", i))
		item.From = parseStringToInt(escapeHTML(getFormKey(f.Values, "resume_employment-list_from_date", i)))
		item.To = parseStringToInt(escapeHTML(getFormKey(f.Values, "resume_employment-list_to_date", i)))
		item.JobTitle = escapeHTML(getFormKey(f.Values, "resume_employment-list_job_title", i))
		item.Summary = escapeHTML(getFormKey(f.Values, "resume_employment-list_summary", i))

		el.Employers = append(el.Employers, item)
	}
	return el
}

// parseFormIntoSkillList takes url values and creates a new skill list, sanitizing values on the way
func parseFormIntoSkillList(fv *forms.Form) *models.SkillList {
	skillList := models.NewSkillList()
	if !hasLength(fv.Values["resume_skill-list_title"]) {
		return skillList
	}

	for i := range fv.Values["resume_skill-list_title"] {
		item := models.NewSkillListItem()

		item.Title = escapeHTML(getFormKey(fv.Values, "resume_skill-list_title", i))
		dur, err := strconv.Atoi(escapeHTML(getFormKey(fv.Values, "resume_skill-list_duration", i)))
		if err != nil {
			dur = 1
		}
		item.Duration = dur
		item.Content = escapeHTML(getFormKey(fv.Values, "resume_skill-list_content", i))

		skillList.Items = append(skillList.Items, item)
	}
	return skillList
}

// parseFormIntoContactDetails takes in the form values, creates a new contact details struct, and returns it, sanitizing the input values on the way
func parseFormIntoContactDetails(f *url.Values) *models.ContactDetail {

	resContactDetails := models.NewContactDetail()

	resContactDetails.Firstname = html.EscapeString(f.Get("resume_contact-details_firstname"))
	resContactDetails.Lastname = html.EscapeString(f.Get("resume_contact-details_lastname"))
	resContactDetails.AddressLine1 = html.EscapeString(f.Get("resume_contact-details_address1"))
	resContactDetails.AddressLine2 = html.EscapeString(f.Get("resume_contact-details_address2"))
	resContactDetails.City = html.EscapeString(f.Get("resume_contact-details_city"))
	resContactDetails.State = html.EscapeString(f.Get("resume_contact-details_state"))
	resContactDetails.Zipcode = html.EscapeString(f.Get("resume_contact-details_zipcode"))
	resContactDetails.Email = html.EscapeString(f.Get("resume_contact-details_email"))
	resContactDetails.PhoneNumber = html.EscapeString(f.Get("resume_contact-details_phone"))
	resContactDetails.WebAddress = html.EscapeString(f.Get("resume_contact-details_webaddress"))
	return resContactDetails
}

// parseSocialList takes in *forms.Form and parses value into *models.SocialMediaList
// sanitizing any string on the way
func parseSocialList(f *forms.Form) *models.SocialMediaList {
	list := models.NewSocialMediaList()

	if !hasLength(f.Values["resume_social-media_company"]) {
		return list
	}

	for i := range f.Values["resume_social-media_company"] {
		var listItem = models.NewSocialMediaListItems()

		listItem.CompanyName = escapeHTML(getFormKey(f.Values, "resume_social-media_company", i))
		listItem.UserName = escapeHTML(getFormKey(f.Values, "resume_social-media_username", i))
		listItem.WebAddress = escapeHTML(getFormKey(f.Values, "resume_social-media_address", i))

		list.SocialMediaListItems = append(list.SocialMediaListItems, listItem)
	}
	return list
}

// hasLength takes in a slice of strings and returns whether it has values or not

func hasLength(ss []string) bool {
	if len(ss) > 0 {
		return true
	}
	return false
}
