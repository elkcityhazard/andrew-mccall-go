package handlers

import (
	"net/http"

	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
	"github.com/elkcityhazard/andrew-mccall-go/internal/render"
)

func (hr *HandlerRepo) HandleGetAdminResume(w http.ResponseWriter, r *http.Request) {

	data, stringMap, intMap, form := hr.CreateEmptyTemplatePayload()

	stringMap["PageTitle"] = "Resume"
	data["UserID"] = hr.app.SessionManager.GetInt64(r.Context(), "id")
	render.RenderTemplate(w, r, "admin-resume.gohtml", &models.TemplateData{
		StringMap: stringMap,
		Data:      data,
		IntMap:    intMap,
		Form:      form,
	})
}
