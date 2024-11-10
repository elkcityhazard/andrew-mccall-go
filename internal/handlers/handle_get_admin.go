package handlers

import (
	"log"
	"net/http"

	"github.com/elkcityhazard/andrew-mccall-go/internal/forms"
	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
	"github.com/elkcityhazard/andrew-mccall-go/internal/render"
)

func (hr *HandlerRepo) HandleGetAdmin(w http.ResponseWriter, r *http.Request) {

	if !hr.app.SessionManager.Exists(r.Context(), "id") {
		http.Redirect(w, r.WithContext(r.Context()), "/login", http.StatusSeeOther)
		return
	}

	userID := hr.app.SessionManager.GetInt64(r.Context(), "id")

	posts, err := hr.conn.GetPaginatedPosts(userID, 0, 20)

	if err != nil {
		log.Println(err)
	}

	data := make(map[string]any)

	data["Posts"] = posts

	stringMap := make(map[string]string)

	stringMap["PageTitle"] = "Admin Home"
	render.RenderTemplate(w, r, "admin-home.gohtml", &models.TemplateData{
		StringMap: stringMap,
		Data:      data,
		Form:      forms.New(nil),
	})
}
