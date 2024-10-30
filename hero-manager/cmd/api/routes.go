package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/heroes", app.createHeroHandler)
	router.HandlerFunc(http.MethodGet, "/heroes/:id", app.showHeroHandler)

	return router
}
