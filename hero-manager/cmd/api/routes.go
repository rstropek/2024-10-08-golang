package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.HandlerFunc(http.MethodGet, "/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/heroes", app.createHeroHandler)
	router.HandlerFunc(http.MethodGet, "/heroes/:id", app.showHeroHandler)
	router.HandlerFunc(http.MethodGet, "/error", app.panicingHandler)

	return app.recoverPanic(router)
}

func (app *application) panicingHandler(w http.ResponseWriter, r *http.Request) {
	x := 10
	y := 0
	z := x / y
	fmt.Fprintln(w, z)
}
