package main

import (
	"context"
	"net/http"
)

type contextKey string
const userContextKey = contextKey("user")

func (app *application) contextSetUser(r *http.Request, user string) *http.Request {
	newContext := context.WithValue(r.Context(), userContextKey, user)
	r = r.WithContext(newContext)
	return r
}

func (app *application) contextGetUser(r *http.Request) string {
	user, ok := r.Context().Value(userContextKey).(string)
	if !ok {
		return "Anonymous"
	}

	return user
}
