package handlers

import (
	"net/http"

	"github.com/elkcityhazard/andrew-mccall-go/internal/forms"
	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
	"github.com/elkcityhazard/andrew-mccall-go/internal/render"
)

func (hr *HandlerRepo) HandleGetCompose(w http.ResponseWriter, r *http.Request) {
	var stringMap = map[string]string{}
	stringMap["PageTitle"] = "Compose"

	render.RenderTemplate(w, r, "compose.gohtml", &models.TemplateData{
		StringMap: stringMap,
		Form:      forms.New(nil),
	})
}
