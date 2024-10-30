package main

import (
	"fmt"
	"net/http"
)

func (app *application) createHeroHandler(w http.ResponseWriter, r *http.Request) {
	// POST /heroes

	fmt.Fprintln(w, "create a new hero")
}

func (app *application) showHeroHandler(w http.ResponseWriter, r *http.Request) {
	// GET /heroes/:id

	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details for hero %d\n", id)
}
