package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) createHeroHandler(w http.ResponseWriter, r *http.Request) {
	// POST /heroes

	fmt.Fprintln(w, "create a new hero")
}

func (app *application) showHeroHandler(w http.ResponseWriter, r *http.Request) {
	// GET /heroes/:id

	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details for hero %d\n", id)
}