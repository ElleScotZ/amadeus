package pkg

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Application struct {
	router *chi.Mux

	search *Search
}

// NewApplication creates a new instance of Application.
// It creates a new Search object as well.
func NewApplication() *Application {
	application := Application{}

	application.search = &Search{}

	application.setUpRoutes()

	return &application
}

// setUpRoutes handles routes.
// It is a good place for potential middleware.
func (a *Application) setUpRoutes() {
	a.router = chi.NewRouter()

	a.router.Mount("/api/v0.1/search", a.search.Routes())
}

// Start initiates ListenAndServe on port 9000 in localhost.
func (a *Application) Start() error {
	return http.ListenAndServe("localhost:9000", a.router)
}
