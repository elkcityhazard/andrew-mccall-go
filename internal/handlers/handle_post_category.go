package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
	"github.com/elkcityhazard/andrew-mccall-go/pkg/utils"
	"github.com/justinas/nosurf"
)

func (hr *HandlerRepo) HandlePostCategory(w http.ResponseWriter, r *http.Request) {
	var util = utils.NewUtil()
	csrfToken := r.Header.Get("X-CSRF-TOKEN")

	if !nosurf.VerifyToken(nosurf.Token(r), csrfToken) {
		err := errors.New("you are not authorized")
		pl := util.CreateJSONEnvelope("error", err.Error())
		err = json.NewEncoder(w).Encode(pl)
		if err != nil {
			returnErr(w, err)
		}
		return
	}

	var cat = models.NewCategory()
	err := json.NewDecoder(r.Body).Decode(cat)

	if err != nil {
		pl := util.CreateJSONEnvelope("error", err.Error())
		err = json.NewEncoder(w).Encode(pl)
		if err != nil {
			returnErr(w, err)
		}
		return
	}

	if len(cat.Name) < 1 {

		errPayload := util.CreateJSONEnvelope("error", errors.New("invalid category name").Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(errPayload)
		if err != nil {
			returnErr(w, err)
			return
		}
		return
	}

	catID, err := hr.conn.InsertCategory(cat)

	if err != nil {
		pl := util.CreateJSONEnvelope("error", err.Error())
		err = json.NewEncoder(w).Encode(pl)
		if err != nil {
			returnErr(w, err)
		}
		return
	}

	cat.ID = catID

	err = json.NewEncoder(w).Encode(cat)

	if err != nil {
		pl := util.CreateJSONEnvelope("error", err.Error())
		err = json.NewEncoder(w).Encode(pl)
		if err != nil {
			returnErr(w, err)
		}
		return
	}

}
