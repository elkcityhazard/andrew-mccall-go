package handlers

import (
	"net/http"

	"github.com/elkcityhazard/andrew-mccall-go/internal/forms"
	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
	"github.com/elkcityhazard/andrew-mccall-go/internal/render"
)

func (hr *HandlerRepo) SignupGetHandler(w http.ResponseWriter, r *http.Request) {

	var stringMap = map[string]string{}

	stringMap["Title"] = "Signup"

	render.RenderTemplate(w, r, "signup.gohtml", &models.TemplateData{
		StringMap: stringMap,
		Form:      forms.New(nil),
	})
}
