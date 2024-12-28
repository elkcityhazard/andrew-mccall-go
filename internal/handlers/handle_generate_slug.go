package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"net/http"

	"github.com/elkcityhazard/andrew-mccall-go/pkg/utils"
	"github.com/justinas/nosurf"
)

func (hr *HandlerRepo) HandleGenerateSlug(w http.ResponseWriter, r *http.Request) {

	if !nosurf.VerifyToken(r.Header.Get("X-CSRF-Token"), nosurf.Token(r)) {
		returnErr(w, errors.New("unauthorized"))
		return
	}

	u := utils.NewUtil()

	type slug struct {
		Value string `json:"value"`
	}

	var incoming slug

	err := json.NewDecoder(r.Body).Decode(&incoming)

	if err != nil {
		fmt.Println(err)
		returnErr(w, err)
		return
	}

	outgoing := u.Slugify(html.EscapeString(incoming.Value))

	type Payload struct {
		Status int    `json:"status"`
		Slug   string `json:"slug"`
	}

	p := Payload{
		Status: http.StatusOK,
		Slug:   fmt.Sprintf("%s", outgoing),
	}

	err = json.NewEncoder(w).Encode(p)

	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
