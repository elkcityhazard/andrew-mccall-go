package main

import (
	"net/http"
	"strings"

	"github.com/justinas/nosurf"
)

func LoadSession(next http.Handler) http.Handler {
	return app.SessionManager.LoadAndSave(next)
}

func returnsJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r.WithContext(r.Context()))
	})
}

func requiresAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isValidSession := app.SessionManager.Exists(r.Context(), "id")
		if !isValidSession {
			http.Redirect(w, r.WithContext(r.Context()), "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r.WithContext(r.Context()))
	})

}

func csrfToken(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.IsProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

func stripTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Path
		switch true {
		case p == "/" && len(p) < 2:
			next.ServeHTTP(w, r.WithContext(r.Context()))
		// handle dir browsing
		case strings.HasPrefix(p, "/static/"):
			if p[len(p)-1:] == "/" {
				http.NotFound(w, r)
				return
			}
			next.ServeHTTP(w, r.WithContext(r.Context()))
		case p[len(p)-1:] == "/":
			http.Redirect(w, r.WithContext(r.Context()), p[:len(p)-1], http.StatusSeeOther)
		default:
			next.ServeHTTP(w, r.WithContext(r.Context()))
		}
	})
}
