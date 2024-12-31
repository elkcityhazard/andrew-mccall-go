package amrouter

import (
	"context"
	"embed"
	"errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/elkcityhazard/andrew-mccall-go/internal/config"
)

var (
	app            *config.AppConfig
	RouteParamsKey = &CtxKey{name: "routeParams"}
)

// Update CtxKey to include a name for debugging
type CtxKey struct {
	name string
}

// String makes CtxKey implement the Stringer interface
func (k *CtxKey) String() string {
	return "amrouter context key " + k.name
}

type AMRouter struct {
	PathToStaticDir   string
	EmbeddedStaticDir embed.FS
	IsProduction      bool
	Routes            []AMRoute
	Middleware        []MiddleWareFunc
	GlobalMiddleware  []MiddleWareFunc
}

func NewRouter(cfg *config.AppConfig) *AMRouter {
	app = cfg
	return &AMRouter{
		Routes:           []AMRoute{},
		Middleware:       []MiddleWareFunc{},
		GlobalMiddleware: []MiddleWareFunc{},
	}
}

type AMRoute struct {
	Method     string
	Path       *regexp.Regexp
	Handler    http.Handler
	Middleware []MiddleWareFunc
}

// MiddleWareFunc is an alias for func(http.Handler) http.Handler
type MiddleWareFunc func(http.Handler) http.Handler

// AddRoute takes a method, pattern, handler, and middleware and adds it to an instance of AMRouter.Routes
// It can return a regex compile error
func (rtr *AMRouter) AddRoute(method string, pattern string, handler http.HandlerFunc, mware ...MiddleWareFunc) error {

	var mwareToAdd = []MiddleWareFunc{}

	if len(mware) > 0 {

		for _, mw := range mware {
			mwareToAdd = append(mwareToAdd, mw)
		}

	}

	re, err := regexp.Compile("^" + pattern + "$")
	if err != nil {
		return err
	}
	rtr.Routes = append(rtr.Routes, AMRoute{
		Method:     method,
		Path:       re,
		Handler:    handler,
		Middleware: mwareToAdd,
	})

	return nil
}

func (rtr *AMRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Don't create new context unnecessarily
	isStatic := rtr.ServeStaticDirectory(w, r)
	if isStatic {
		return
	}

	var allow []string

	for _, route := range rtr.Routes {
		matches := route.Path.FindStringSubmatch(r.URL.Path)

		if len(matches) > 0 {
			if r.Method != route.Method {
				allow = append(allow, route.Method)
				continue
			}
			// Store route parameters in context if needed

			ctx := context.WithValue(r.Context(), RouteParamsKey, matches[1:])
			r = r.WithContext(ctx)

			var handler http.Handler = route.Handler

			// middleware gets handled outside in, so add route based first, then global
			if len(route.Middleware) > 0 {
				handler = rtr.AddMiddlewareToHandler(handler, route.Middleware...)
			}

			if len(rtr.GlobalMiddleware) > 0 {
				handler = rtr.AddMiddlewareToHandler(handler, rtr.GlobalMiddleware...)
			}

			handler.ServeHTTP(w, r)
			return
		}
	}

	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		w.WriteHeader(405)
		err := errors.New("405 method not allowed")
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
	} else {
		rtr.Custom404Handler(w, r)
	}
}

// ServeStaticDirectory accepts an http.ResponseWriter, and a *http.Request and determins if
// the current r.URL.Path is to a static file.  It returns a bool to indicate if the rest of the
// ServeHTTP function shoulbe be short circuited
func (rtr *AMRouter) ServeStaticDirectory(w http.ResponseWriter, r *http.Request) bool {
	// handle static directory
	if strings.HasPrefix(r.URL.Path, rtr.PathToStaticDir) {
		// if not in prod, load static resources from disk, else embed
		if !rtr.IsProduction {
			fileServer := http.FileServer(http.Dir(rtr.PathToStaticDir))
			http.StripPrefix("/static/", fileServer).ServeHTTP(w, r)

		} else {
			fileServer := http.FileServer(http.FS(rtr.EmbeddedStaticDir))
			http.StripPrefix("/static/", fileServer).ServeHTTP(w, r)
		}

		return true
	}
	return false

}

// Use adds global middleware to all routes
func (rtr *AMRouter) Use(mw func(http.Handler) http.Handler) {

	rtr.GlobalMiddleware = append(rtr.GlobalMiddleware, mw)

}

// AddMiddlewareToHandler applies middleware in reverse order
func (rtr *AMRouter) AddMiddlewareToHandler(handler http.Handler, middleware ...MiddleWareFunc) http.Handler {
	// Apply middleware in reverse order to maintain correct execution order
	for i := len(middleware) - 1; i >= 0; i-- {
		currentMiddleware := middleware[i]
		handler = currentMiddleware(handler)
	}
	return handler
}

func (rtr *AMRouter) Custom404Handler(w http.ResponseWriter, r *http.Request) {
	notFoundHandler := http.NotFoundHandler()

	if len(rtr.GlobalMiddleware) > 0 {
		notFoundHandler = rtr.AddMiddlewareToHandler(notFoundHandler, rtr.GlobalMiddleware...)
	}

	notFoundHandler.ServeHTTP(w, r)
}
