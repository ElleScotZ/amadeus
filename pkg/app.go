package pkg

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Application struct {
	router *chi.Mux

	search *Search
}

func NewApplication() *Application {
	application := Application{}

	application.search = &Search{}

	application.setUpRoutes()

	return &application
}

func (a *Application) setUpRoutes() {
	a.router = chi.NewRouter()

	a.router.Mount("/api/v0.1/search", a.search.Routes())
}

func (a *Application) Start() error {
	return http.ListenAndServe("localhost:9000", a.router)
}
