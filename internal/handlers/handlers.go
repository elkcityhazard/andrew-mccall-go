package handlers

import (
	"net/http"

	"github.com/elkcityhazard/andrew-mccall-go/internal/config"
	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
	"github.com/elkcityhazard/andrew-mccall-go/internal/render"
	"github.com/elkcityhazard/andrew-mccall-go/internal/repository"
)

type CtxKey struct{}

func GetField(r *http.Request, index int) string {
	fields := r.Context().Value(CtxKey{}).([]string)
	return fields[index]
}

type HandlerRepo struct {
	app  *config.AppConfig
	conn repository.DBServicer
}

var Repo *HandlerRepo

func NewHandlerRepo(a *config.AppConfig, servicer repository.DBServicer) *HandlerRepo {
	return &HandlerRepo{
		app:  a,
		conn: servicer,
	}
}

func SetHandlerRepo(hr *HandlerRepo) {
	Repo = hr
}

func (hr *HandlerRepo) HomeGetHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{}

	posts, err := hr.conn.ListPosts(3, 0)

	if err != nil {
		returnErr(w, err)
		return
	}

	data["Posts"] = posts

	stringMap := make(map[string]string)

	stringMap["PageTitle"] = "Home"

	var td = &models.TemplateData{
		StringMap: stringMap,
		Data:      data,
	}
	render.RenderTemplate(w, r, "home.gohtml", td)
	return

}
