package handlers

import (
	"errors"
	"html"
	"net/http"
	"time"

	"github.com/elkcityhazard/andrew-mccall-go/internal/forms"
	"github.com/elkcityhazard/andrew-mccall-go/internal/mailer"
	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
	"github.com/elkcityhazard/andrew-mccall-go/internal/render"
)

func (hr *HandlerRepo) HandlePostContact(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["PageTitle"] = "Contact"
	err := r.ParseForm()

	if err != nil {
		returnErr(w, err)
		return
	}

	f := forms.New(r.Form)

	username := f.Get("username")

	if username != "" {
		http.NotFound(w, r)
		return

	}

	f.Required("email", "message")

	f.MinLength("email", 3)
	f.MinLength("message", 3)

	isEmail := f.IsEmail(f.Values.Get("email"))

	if !isEmail {
		f.Errors.Add("email", "invalid email")
	}

	if !f.Valid() {
		render.RenderTemplate(w, r, "contact.gohtml", &models.TemplateData{
			StringMap: stringMap,
			Form:      f,
		})
		return
	}

	msg := models.ContactMsg{}

	msg.Email = html.EscapeString(f.Get("email"))
	msg.Message = html.EscapeString(f.Get("message"))
	msg.CreatedAt = time.Now()
	msg.Version = 1

	_, err = hr.conn.InsertMessage(&msg)

	if err != nil {
		hr.app.SessionManager.Put(r.Context(), "error", errors.New("error handling the message").Error())
		render.RenderTemplate(w, r, "contact.gohtml", &models.TemplateData{
			StringMap: stringMap,
			Form:      f,
		})

	}

	mailData := make(map[string]any)

	mailData["Email"] = msg.Email
	mailData["Message"] = msg.Message

	mail := mailer.NewMailMsgPayload()

	mail.Recipient = "andrew@andrew-mccall.com"
	mail.Data = mailData
	mail.Template = "contact-form.gohtml"

	hr.app.Mailer.MsgChan <- mail

	hr.app.SessionManager.Put(r.Context(), "flash", "Your message was submitted.  I will respond as soon as possible")

	http.Redirect(w, r.WithContext(r.Context()), "/success", http.StatusSeeOther)

}
