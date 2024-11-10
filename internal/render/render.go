package render

import (
	"bytes"
	"errors"
	"log"
	"net/http"

	"github.com/elkcityhazard/andrew-mccall-go/internal/config"
	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
	"github.com/justinas/nosurf"
)

var app *config.AppConfig

func NewRenderer(a *config.AppConfig) {
	app = a
}

func AddDefaultTemplateData(td *models.TemplateData, r *http.Request) *models.TemplateData {

	if td == nil {
		td = &models.TemplateData{}
	}

	td.SiteTitle = app.SiteTitle
	td.IsAuthenticated = app.SessionManager.Exists(r.Context(), "id")
	td.CSRFToken = nosurf.Token(r)
	td.Error = app.SessionManager.PopString(r.Context(), "error")
	td.Flash = app.SessionManager.PopString(r.Context(), "flash")
	td.Warning = app.SessionManager.PopString(r.Context(), "warning")

	td.PopulateAdminMenu()

	return td

}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {

	td = AddDefaultTemplateData(td, r)

	var buf = new(bytes.Buffer)

	t, ok := app.TemplateCache[tmpl]

	if !ok {
		log.Panic(errors.New("error with template cache"))
		return
	}

	err := t.Execute(buf, td)

	if err != nil {
		log.Fatalln(err)
	}

	_, err = buf.WriteTo(w)

	if err != nil {
		log.Panic(err)
	}

}
