package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
	"github.com/elkcityhazard/andrew-mccall-go/internal/render"
)

func (hr *HandlerRepo) HandleGetSignupSuccess(w http.ResponseWriter, r *http.Request) {

	id := hr.app.SessionManager.GetInt64(r.Context(), "id")

	if id == 0 {
		err := errors.New("invalid request")
		returnErr(w, err)
		return
	}
	user, err := hr.conn.GetUserByID(id)

	if err != nil {
		returnErr(w, err)
		return
	}

	stringMap := make(map[string]string)

	stringMap["PageTitle"] = fmt.Sprintf("%s - success", user.Email)

	data := make(map[string]any)
	data["User"] = user

	render.RenderTemplate(w, r, "success.gohtml", &models.TemplateData{
		StringMap: stringMap,
		Data:      data,
	})
}
