package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/wbaker85/tacklebox/pkg/models"
)

func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
	var u userInfo

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
	}

	_, err = app.users.Insert(u.Email, u.Password)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errJSON{"email already registered"})
			return
		}

		app.serverError(w, err)
		return
	}

	json.NewEncoder(w).Encode(infoJSON{"success"})
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	var u userInfo
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
	}

	id, err := app.users.Authenticate(u.Email, u.Password)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(errJSON{"invalid credentials"})
			return
		}
		app.serverError(w, err)
		return
	}

	app.session.Put(r, "authenticatedUserID", id)

	json.NewEncoder(w).Encode(infoJSON{u.Email})
}

func (app *application) logout(w http.ResponseWriter, r *http.Request) {
	app.session.Remove(r, "authenticatedUserID")
}

func (app *application) loggedInUser(w http.ResponseWriter, r *http.Request) {
	id := app.session.GetInt(r, "authenticatedUserID")

	u, err := app.users.Get(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(infoJSON{u.Email})
}
