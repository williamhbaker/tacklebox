package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/wbaker85/tacklebox/pkg/models"

	"github.com/google/uuid"
)

func (app *application) createBin(w http.ResponseWriter, r *http.Request) {
	userID := app.session.GetInt(r, "authenticatedUserID")
	binID := uuid.New().String()

	_, err := app.bins.Insert(binID, userID)
	if err != nil {
		if errors.Is(err, models.ErrInvalidUser) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errJSON{"invalid user"})
			return
		}

		app.serverError(w, err)
		return
	}

	json.NewEncoder(w).Encode(infoJSON{binID})
}

func (app *application) getBinHooks(w http.ResponseWriter, r *http.Request) {

}

func (app *application) destroyBin(w http.ResponseWriter, r *http.Request) {
	bin := &binInfo{}

	err := json.NewDecoder(r.Body).Decode(bin)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
	}

	bin.UserID = app.session.GetInt(r, "authenticatedUserID")
	err = app.bins.Destroy(bin.ID, bin.UserID)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		app.serverError(w, err)
		return
	}

	json.NewEncoder(w).Encode(infoJSON{"success"})
}
