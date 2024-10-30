package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	app.logger.Info("healthcheckHandler called")

	fmt.Fprintln(w, "status: available")
}
