package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rstropek/hero-manager/internal/data"
)

func (app *application) createHeroHandler(w http.ResponseWriter, r *http.Request) {
	// POST /heroes

	fmt.Fprintln(w, "create a new hero")
}

func (app *application) showHeroHandler(w http.ResponseWriter, r *http.Request) {
	// GET /heroes/:id

	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	hero := data.Hero{
		ID:        id,
		Name:      "Deadpool",
		CanFly:    false,
		FirstSeen: time.Now(),
		Version:   1,
		RealName:  "Wade Wilson",
		Abilities: []string{"Accelerated Healing", "Super Strong"},
	}

	err = app.writeJSON(w, http.StatusOK, hero, nil)
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}
