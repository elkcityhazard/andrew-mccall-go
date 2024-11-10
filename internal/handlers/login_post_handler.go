package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/elkcityhazard/andrew-mccall-go/internal/forms"
	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
	"github.com/elkcityhazard/andrew-mccall-go/internal/render"
)

func (hr *HandlerRepo) LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	var stringMap = map[string]string{}

	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	form := forms.New(r.Form)

	form.Required("email", "password")

	form.MinLength("email", 2)

	form.MinLength("password", 8)

	isValidEmail := form.IsEmail(form.Get("email"))

	if !isValidEmail {
		form.Errors.Add("email", "invalid email provided")
	}

	if !form.Valid() {

		render.RenderTemplate(w, r, "login.gohtml", &models.TemplateData{
			StringMap: stringMap,
			Form:      form,
		})
		return
	}

	u, err := hr.conn.GetUserByEmail(form.Get("email"))

	if err != nil {

		form.Errors.Add("email", "something went wrong")
		form.Errors.Add("password", "something went wrong")

		render.RenderTemplate(w, r, "login.gohtml", &models.TemplateData{
			StringMap: stringMap,
			Form:      form,
		})
		return
	}

	if !u.IsActive {
		returnErr(w, errors.New("something went wrong"))
		return
	}

	passwordIsValid, err := argon2id.ComparePasswordAndHash(form.Get("password"), u.Password.Hash)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !passwordIsValid {
		form.Errors.Add("password", "invalid credentials, try again")
		form.Errors.Add("email", "invalid credentials, try again.")
	}

	if !form.Valid() {

		render.RenderTemplate(w, r, "login.gohtml", &models.TemplateData{
			StringMap: stringMap,
			Form:      form,
		})
		return
	}

	hr.app.SessionManager.Put(r.Context(), "id", u.ID)
	hr.app.SessionManager.Put(r.Context(), "flash", fmt.Sprintf("Welcome, %s", u.Email))

	http.Redirect(w, r.WithContext(r.Context()), "/admin", http.StatusSeeOther)

}
