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
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errJSON{"invalid user"})
			return
		}

		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(infoJSON{binID})
}

func (app *application) getBinHooks(w http.ResponseWriter, r *http.Request) {

}

func (app *application) deleteBin(w http.ResponseWriter, r *http.Request) {

}
