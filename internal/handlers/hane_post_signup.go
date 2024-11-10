package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/elkcityhazard/andrew-mccall-go/internal/forms"
	"github.com/elkcityhazard/andrew-mccall-go/internal/mailer"
	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
	"github.com/elkcityhazard/andrew-mccall-go/internal/render"
	"github.com/elkcityhazard/andrew-mccall-go/pkg/utils"
)

func (hr *HandlerRepo) HandlePostSignup(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	const (
		minPWLength = 8
	)

	if err != nil {
		returnErr(w, err)
		return
	}

	f := forms.New(r.Form)

	f.Required("email", "username", "password1", "password2")

	f.MinLength("email", 4)
	f.MinLength("username", 6)
	f.MinLength("password1", minPWLength)
	f.MinLength("password2", minPWLength)

	isEmail := f.IsEmail(f.Get("email"))

	if !isEmail {
		f.Errors.Add("email", "invalid email signature")
	}

	if !strings.EqualFold(f.Get("password1"), f.Get("password2")) {
		f.Errors.Add("password1", "passwords do not match")
		f.Errors.Add("password2", "passwords do not match")
	}

	var util = utils.NewUtil()

	isComplexPW := util.CheckPWStrength(f.Get("password1"), minPWLength)

	if !isComplexPW {
		f.Errors.Add("password1", "invalid password")
		f.Errors.Add("password2", "invalid password")
	}

	user, _ := hr.conn.GetUserByEmail(f.Get("email"))

	if user != nil {
		f.Errors.Add("email", "invalid entry")
	}

	if !f.Valid() {

		stringMap := make(map[string]string)
		stringMap["PageTitle"] = "Singup"
		render.RenderTemplate(w, r, "signup.gohtml", &models.TemplateData{
			StringMap: stringMap,
			Form:      f,
		})
		return
	}

	// handle user creation

	var newUser = &models.User{}

	var newPW = &models.Password{}

	newUser.Email = f.Get("email")
	newUser.Username = f.Get("username")
	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()
	newUser.IsActive = false
	newUser.Role = "user"
	newUser.Version = 1

	hash, err := argon2id.CreateHash(f.Get("password1"), argon2id.DefaultParams)

	if err != nil {
		returnErr(w, err)
		return
	}

	newPW.Hash = hash
	newPW.PlainText = f.Get("password1")
	newPW.CreatedAt = time.Now()
	newPW.UpdatedAt = time.Now()
	newPW.IsActive = false
	newPW.Version = 1
	newUser.Password = newPW

	insertedUser, err := hr.conn.InsertUser(newUser)

	if err != nil {
		returnErr(w, err)
		return
	}

	newUser.ID = insertedUser

	at := models.NewActivationToken()

	err = at.GenerateToken(insertedUser, time.Hour*48, "user")

	if err != nil {
		returnErr(w, err)
		return
	}

	_, err = hr.conn.InsertActivationToken(at)

	if err != nil {
		returnErr(w, err)
		return
	}

	var mailMsg = &mailer.MailMsgPayload{}

	mailMsg.Recipient = newUser.Email
	mailMsg.Template = "welcome.gohtml"

	mailData := make(map[string]any)

	mailData["Email"] = newUser.Email

	if !hr.app.IsProduction {
		mailData["ActivationLink"] = template.HTML(fmt.Sprintf("http://localhost:8080/users/activation?token=%s&id=%d", at.Plaintext, at.UserID))

	} else {
		mailData["ActivationLink"] = template.HTML(fmt.Sprintf("https://www.andrew-mccall.com/users/activation?token=%s&id=%d", at.Plaintext, at.UserID))
	}

	mailMsg.Data = mailData

	hr.app.Mailer.MsgChan <- mailMsg
	hr.app.SessionManager.Put(r.Context(), "id", insertedUser)

	http.Redirect(w, r.WithContext(r.Context()), "/signup/success", http.StatusSeeOther)

}
