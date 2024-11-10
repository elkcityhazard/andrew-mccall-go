package main

import (
	"context"
	"errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/elkcityhazard/andrew-mccall-go/internal/handlers"
	"github.com/elkcityhazard/andrew-mccall-go/internal/models"
	"github.com/elkcityhazard/andrew-mccall-go/internal/render"
	"github.com/elkcityhazard/andrew-mccall-go/internal/static"
)

type router struct {
	routes []*route
}

func (m *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Handle static files first
	if strings.HasPrefix(r.URL.Path, "/static/") {
		// if not in prod, load static resources from disk, else embed
		if !app.IsProduction {
			fileServer := http.FileServer(http.Dir("./internal/static/"))
			http.StripPrefix("/static/", fileServer).ServeHTTP(w, r)

		} else {
			fileServer := http.FileServer(http.FS(static.GetStaticDir()))
			http.StripPrefix("/static/", fileServer).ServeHTTP(w, r)
		}
		return
	}

	if strings.HasPrefix(r.URL.Path, "/uploads/") {
		fileServer := http.FileServer(http.Dir("./uploads/"))
		http.StripPrefix("/uploads/", fileServer).ServeHTTP(w, r)

		return
	}

	var allow []string

	for _, v := range m.routes {
		matches := v.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != v.method {
				allow = append(allow, v.method)
				continue
			}

			ctx := context.WithValue(r.Context(), handlers.CtxKey{}, matches[1:])

			// handle middleware
			// last in first out algo

			if len(v.middleware) > 0 {
				for i := len(v.middleware) - 1; i >= 0; i-- {
					// pass handler into middleware and update handler with new middleware
					v.handler = v.middleware[i](v.handler).ServeHTTP
				}
			}

			v.handler(w, r.WithContext(ctx))
			return
		}
	}
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		w.WriteHeader(405)
		err := errors.New("405 method not allowed")
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(404)
	render.RenderTemplate(w, r, "404.gohtml", &models.TemplateData{})
	return

}

// New accepts a method, pattern, http.HandlerFunc, and an array of middlewares and adds it to the list of allowed routes
func (m *router) New(method, pattern string, handler http.HandlerFunc, mware ...func(next http.Handler) http.Handler) {
	m.routes = append(m.routes, &route{
		method:     method,
		regex:      regexp.MustCompile("^" + pattern + "$"),
		handler:    handler,
		middleware: mware,
	})
}

type route struct {
	method     string
	regex      *regexp.Regexp
	handler    http.HandlerFunc
	middleware []func(http.Handler) http.Handler
}

func routes() http.Handler {

	rtr := &router{}

	rtr.New("GET", "/", handlers.Repo.HomeGetHandler)
	rtr.New("GET", "/success", handlers.Repo.HandleGetSuccess)

	rtr.New("GET", "/signup", handlers.Repo.SignupGetHandler)
	rtr.New("POST", "/signup", handlers.Repo.HandlePostSignup)
	rtr.New("GET", "/signup/success", handlers.Repo.HandleGetSignupSuccess)

	rtr.New("GET", "/login", handlers.Repo.LoginGetHandler)
	rtr.New("POST", "/login", handlers.Repo.LoginPostHandler)

	rtr.New("GET", "/admin", handlers.Repo.HandleGetAdmin, requiresAuth)
	rtr.New("GET", "/admin/compose", handlers.Repo.HandleGetCompose, requiresAuth)
	rtr.New("POST", "/admin/compose", handlers.Repo.HandlePostCompose, requiresAuth)
	rtr.New("POST", "/admin/delete", handlers.Repo.HandleDeletePost, requiresAuth)

	rtr.New("GET", "/admin/compose/edit/([0-9]+)", handlers.Repo.HandleGetEditCompose, requiresAuth)
	rtr.New("PUT", "/admin/compose/edit/([0-9]+)", handlers.Repo.HandlePutCompose, requiresAuth, returnsJSON)

	// Admin Manage Categories
	rtr.New("GET", "/admin/categories", handlers.Repo.HandleGetAdminCategories, requiresAuth)

	rtr.New("GET", "/blog", handlers.Repo.HandleGetBlog)
	rtr.New("GET", `/blog/([\w-\/]+)`, handlers.Repo.HandleGetPost)

	// handle activation
	rtr.New("GET", "/users/activation", handlers.Repo.HandleGetUserActivation)
	rtr.New("POST", "/users/activation", handlers.Repo.HandlePostUserActivation)

	// Contact
	rtr.New("GET", "/contact", handlers.Repo.HandleGetContact)
	rtr.New("POST", "/contact", handlers.Repo.HandlePostContact)

	//API routes

	rtr.New("POST", "/api/v1/upload/image", handlers.Repo.HandlePostUploadImage, requiresAuth, returnsJSON)
	rtr.New("POST", "/api/v1/category", handlers.Repo.HandlePostCategory, requiresAuth, returnsJSON)

	return rtr

}
