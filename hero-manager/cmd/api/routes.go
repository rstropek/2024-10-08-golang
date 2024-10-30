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
	router.HandlerFunc(http.MethodGet, "/user", app.userHandler)

	return app.recoverPanic(app.authenticate(router))
}

func (app *application) panicingHandler(w http.ResponseWriter, r *http.Request) {
	x := 10
	y := 0
	z := x / y
	fmt.Fprintln(w, z)
}

func (app *application) userHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)
	fmt.Fprintf(w, "user: %s\n", user)
}
