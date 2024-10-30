package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rstropek/hero-manager/internal/data"
)

func (app *application) createHeroHandler(w http.ResponseWriter, r *http.Request) {
	// POST /heroes

	var input struct {
		Name      string    `json:"name"`
		CanFly    bool      `json:"canFly"`
		FirstSeen time.Time `json:"firstSeen"`
		RealName  string    `json:"realName,omitempty"`
		Abilities []string  `json:"abilities,omitempty"`
	}

	err := app.readJSON(r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	//body, err := io.ReadAll(r.Body)
	//if err != nil {
	//	app.badRequestResponse(w, r, err)
	//	return
	//}
	//err = json.Unmarshal(body, &input)
	//if err != nil {
	//	app.badRequestResponse(w, r, err)
	//	return
	//}

	fmt.Fprintf(w, "create a new hero: %+v\n", input)
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
