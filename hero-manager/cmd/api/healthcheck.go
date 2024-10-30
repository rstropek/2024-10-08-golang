package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]any{
		"status":      "available",
		"environment": app.config.env,
	}

	err := app.writeJSON(w, http.StatusOK, response, nil)
	if err != nil {
		app.internalServerErrorResponse(w, r, err)
	}
}
