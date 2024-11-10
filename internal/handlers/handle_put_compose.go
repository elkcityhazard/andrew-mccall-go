package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/justinas/nosurf"
)

func (hr *HandlerRepo) HandlePutCompose(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-CSRF-TOKEN")
	if !nosurf.VerifyToken(token, nosurf.Token(r)) {
		err := errors.New("invalid token")
		returnErr(w, err)
		return
	}

	if !hr.app.IsProduction {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}

	// create a pointer type to check for what has updated

	type UpdatedContent struct {
		ID            *int64  `json:"ID,omitempty"`
		Title         *string `json:"title,omitempty"`
		Slug          *string `json:"slug,omitempty"`
		Description   *string `json:"description,omitempty"`
		EditorContent *string `json:"editorContent,omitempty"`
		EditorDelta   *string `json:"editorDelta,omitempty"`
		Status        *string `json:"status,omitempty"`
	}

	var c UpdatedContent

	if r.Body != nil {
		defer r.Body.Close()

		err := json.NewDecoder(r.Body).Decode(&c)

		if err != nil {
			w.WriteHeader(400)
			returnErr(w, err)
			return
		}

		toUpdate, err := hr.conn.GetBlogPostByID(*c.ID)

		if err != nil {
			w.WriteHeader(400)
			returnErr(w, err)
			return
		}

		// validate changes

		if c.ID != nil {
			toUpdate.ID = *c.ID
		}

		if c.Title != nil {
			toUpdate.Title = *c.Title
		}

		if c.Slug != nil {
			toUpdate.Slug = *c.Slug
		}

		if c.Description != nil {
			toUpdate.Description = *c.Description
		}

		if c.EditorContent != nil {
			toUpdate.Content = *c.EditorContent
		}

		if c.EditorDelta != nil {
			toUpdate.Delta = *c.EditorDelta
		}

		if c.Status != nil {
			toUpdate.Status = *c.Status
		}

		toUpdate.UpdatedAt = time.Now()

		// handle update op

		affected, err := hr.conn.UpdatePost(toUpdate)

		if err != nil {
			returnErr(w, err)
			return
		}

		if affected > 0 {
			toUpdate.Version = toUpdate.Version + 1
		}

		w.WriteHeader(200)
		err = json.NewEncoder(w).Encode(toUpdate)

		if err != nil {
			w.WriteHeader(500)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	} else {
		returnErr(w, errors.New("no op"))
	}

}
