package handlers

import (
	"net/http"

	"github.com/elkcityhazard/andrew-mccall-go/internal/forms"
	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
	"github.com/elkcityhazard/andrew-mccall-go/internal/render"
)

func (hr *HandlerRepo) HandleGetContact(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)

	stringMap["PageTitle"] = "Contact"
	render.RenderTemplate(w, r, "contact.gohtml", &models.TemplateData{
		StringMap: stringMap,
		Form:      forms.New(nil),
	})
}
