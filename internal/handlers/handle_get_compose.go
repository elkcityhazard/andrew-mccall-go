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

	data := make(map[string]any)

	cats, err := hr.conn.ListCategories()

	if err != nil {
		hr.app.SessionManager.Put(r.Context(), "error", "there was an error fetching the categories")
	}

	data["Categories"] = cats

	render.RenderTemplate(w, r, "compose.gohtml", &models.TemplateData{
		StringMap: stringMap,
		Data:      data,
		Form:      forms.New(nil),
	})
}
