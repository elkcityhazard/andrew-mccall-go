package handlers

import (
	"errors"
	"net/http"

	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
	"github.com/elkcityhazard/andrew-mccall-go/internal/render"
)

func (hr *HandlerRepo) HandlePostUserActivation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		returnErr(w, err)
		return
	}

	token := r.Form.Get("token")
	id := r.Form.Get("id")

	if token == "" || id == "" {
		returnErr(w, errors.New("invalid request"))
		return
	}

	user, at, err := hr.conn.GetActivationToken(token)

	if err != nil || at == nil {

		if at == nil {
			err = errors.New("invalid token")
		}

		render.RenderTemplate(w, r, "get-user-activation.gohtml", &models.TemplateData{
			Error: err.Error(),
		})
		return

	}

	activatedUserID, err := hr.conn.ActivateUser(user)

	if err != nil {

		render.RenderTemplate(w, r, "get-user-activation.gohtml", &models.TemplateData{
			Error: err.Error(),
		})
		return
	}

	hr.app.SessionManager.Put(r.Context(), "id", activatedUserID)

	http.Redirect(w, r.WithContext(r.Context()), "/admin", http.StatusSeeOther)

}
