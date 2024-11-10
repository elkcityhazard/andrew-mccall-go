package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
	"github.com/elkcityhazard/andrew-mccall-go/internal/render"
)

func (hr *HandlerRepo) HandleGetUserActivation(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	id := r.URL.Query().Get("id")

	if token == "" || id == "" {
		err := errors.New("there has been an error")
		returnErr(w, err)
		return
	}

	user, actTok, err := hr.conn.GetActivationToken(token)

	if err != nil {
		returnErr(w, err)
		return
	}

	stringMap := make(map[string]string)
	stringMap["PageTitle"] = fmt.Sprintf("Welcome, %s!", user.Email)

	data := make(map[string]any)

	data["User"] = user
	data["ActivationToken"] = actTok

	render.RenderTemplate(w, r, "get-user-activation.gohtml", &models.TemplateData{
		StringMap: stringMap,
		Data:      data,
	})

}
