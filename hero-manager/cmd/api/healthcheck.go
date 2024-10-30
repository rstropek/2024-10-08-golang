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
		app.logger.Error(err.Error())
		http.Error(w, "the server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
