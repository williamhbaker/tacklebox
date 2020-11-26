package main

import (
	"context"
	"errors"
	"net/http"

	"github.com/wbaker85/tacklebox/pkg/models"
)

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}

func (app *application) requireJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		if contentType != "application/json" {
			app.clientError(w, http.StatusUnsupportedMediaType)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := app.session.GetInt(r, "authenticatedUserID")
		if id == 0 {
			app.session.Remove(r, "authenticatedUserID")
			next.ServeHTTP(w, r)
			return
		}

		_, err := app.users.Get(id)
		if err != nil {
			if errors.Is(err, models.ErrInvalidUser) {
				app.session.Remove(r, "authenticatedUserID")
				next.ServeHTTP(w, r)
				return
			}
			app.serverError(w, err)
			return
		}

		ctx := context.WithValue(r.Context(), contextKeyIsAuthenticated, true)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) requireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !app.isAuthenticated(r) {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		w.Header().Add("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}
