package main

import (
	"fmt"
	"net/http"
	"strconv"
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

	hero := &data.Hero{
		Name:      input.Name,
		CanFly:    input.CanFly,
		FirstSeen: input.FirstSeen,
		RealName:  input.RealName,
		Abilities: input.Abilities,
	}

	err = hero.Validate()
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

	err = app.repos.Heroes.Insert(hero)
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/heroes/%d", hero.ID))

	err = app.writeJSON(w, http.StatusCreated, hero, headers)
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}

func (app *application) showHeroHandler(w http.ResponseWriter, r *http.Request) {
	// GET /heroes/:id

	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	hero, err := app.repos.Heroes.Get(id)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.writeJSON(w, http.StatusOK, hero, nil)
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}

func (app *application) listHeroesHandler(w http.ResponseWriter, r *http.Request) {
	// GET /heroes

	qs := r.URL.Query()
	page, err := strconv.Atoi(qs.Get("page"))
	if err != nil {
		page = 0
	}

	pageSize, err := strconv.Atoi(qs.Get("pageSize"))
	if err != nil {
		pageSize = 5
	}

	sort := qs.Get("sort")

	heroes, err := app.repos.Heroes.GetAll(page, pageSize, sort)
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, heroes, nil)
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}
