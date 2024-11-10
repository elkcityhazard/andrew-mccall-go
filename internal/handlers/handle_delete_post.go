package handlers

import (
	"net/http"
	"strconv"
)

func (hr *HandlerRepo) HandleDeletePost(w http.ResponseWriter, r *http.Request) {
	userID := hr.app.SessionManager.GetInt64(r.Context(), "id")

	if userID == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)

	if err != nil {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	if id == 0 {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	res, err := hr.conn.DeletePostById(id, userID)

	stringMap := make(map[string]string)

	stringMap["PageTitle"] = "Admin"

	if err != nil {
		hr.app.SessionManager.Put(r.Context(), "error", err.Error())
		http.Redirect(w, r.WithContext(r.Context()), "/admin", http.StatusSeeOther)
		return

	}

	if res > 0 {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return

	}

	http.NotFound(w, r)

}
