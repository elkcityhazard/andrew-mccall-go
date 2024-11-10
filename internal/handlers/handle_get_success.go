package handlers

import (
	"net/http"

	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
	"github.com/elkcityhazard/andrew-mccall-go/internal/render"
)

func (hr *HandlerRepo) HandleGetSuccess(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["PageTitle"] = "Success"

	render.RenderTemplate(w, r, "default-success.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})

}
