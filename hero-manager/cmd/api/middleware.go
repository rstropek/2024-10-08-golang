package main

import (
	"fmt"
	"net/http"
	"strings"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.internalServerErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")

		if authorizationHeader == "" {
			next.ServeHTTP(w, r)
			return
		}

		parts := strings.Split(authorizationHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			app.unauthorizedResponse(w, r)
			return
		}

		userName := parts[1]
		// jwtToke := parts[1]
		// Validate jwt Token (I would NOT implement this myself!!!!!)

		r = app.contextSetUser(r, userName)

		next.ServeHTTP(w, r)
	})
}
