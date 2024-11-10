package handlers

import (
	"fmt"
	"html"
	"net/http"
	"time"

	"github.com/elkcityhazard/andrew-mccall-go/internal/forms"
	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
	"github.com/elkcityhazard/andrew-mccall-go/internal/render"
)

func (hr *HandlerRepo) HandlePostCompose(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		returnErr(w, err)
		return
	}

	formVals := forms.New(r.Form)

	formVals.Required("title")
	formVals.Required("slug")
	formVals.Required("description")
	formVals.Required("editorContent")
	formVals.Required("editorDelta")

	formVals.MinLength("editorContent", 1)
	formVals.MinLength("description", 1)
	formVals.MinLength("editorDelta", 1)
	if !formVals.Valid() {
		fmt.Println(formVals.Errors)
		var stringMap = map[string]string{}
		stringMap["pageTitle"] = "Compose"

		var data = map[string]interface{}{}

		data["EditorContent"] = formVals.Get("editorContent")

		render.RenderTemplate(w, r, "compose.gohtml", &models.TemplateData{
			Form:      formVals,
			StringMap: stringMap,
			Data:      data,
		})
		return
	}

	var content = &models.Content{}

	content.Title = html.EscapeString(formVals.Get("title"))
	content.UserId = hr.app.SessionManager.GetInt64(r.Context(), "id")
	content.Slug = html.EscapeString(formVals.Get("slug"))
	content.Description = html.EscapeString(formVals.Get("description"))
	content.CreatedAt = time.Now()
	content.UpdatedAt = time.Now()
	content.Status = "published"
	content.Version = 1
	content.Content = formVals.Get("editorContent")
	content.Delta = formVals.Get("editorDelta")

	_, err = hr.conn.InsertEditorContent(content)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r.WithContext(r.Context()), "/admin", http.StatusSeeOther)

}
