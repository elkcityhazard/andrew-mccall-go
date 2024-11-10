package handlers

import (
	"net/http"

	"github.com/elkcityhazard/andrew-mccall-go/internal/forms"
	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
	"github.com/elkcityhazard/andrew-mccall-go/internal/render"
)

func (hr *HandlerRepo) HandleGetAdminCategories(w http.ResponseWriter, r *http.Request) {

	cats, err := hr.conn.ListCategories()

	if err != nil {
		returnErr(w, err)
		return
	}

	stringMap := make(map[string]string)

	stringMap["PageTitle"] = "Categories - Admin"

	data := make(map[string]any)
	data["Categories"] = cats

	render.RenderTemplate(w, r, "admin-categories.gohtml", &models.TemplateData{
		StringMap: stringMap,
		Data:      data,
		Form:      forms.New(nil),
	})

}
