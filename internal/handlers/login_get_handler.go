package handlers

import (
	"net/http"

	"github.com/elkcityhazard/andrew-mccall-go/internal/forms"
	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
	"github.com/elkcityhazard/andrew-mccall-go/internal/render"
)

func (hr *HandlerRepo) LoginGetHandler(w http.ResponseWriter, r *http.Request) {

	var stringMap = make(map[string]string)
	stringMap["Title"] = "Login"
	switch r.Method {
	case "GET":

		render.RenderTemplate(w, r, "login.gohtml", &models.TemplateData{
			StringMap: stringMap,
			Form:      forms.New(nil),
		})
	case "POST":

	default:
		http.NotFound(w, r)
	}

}
