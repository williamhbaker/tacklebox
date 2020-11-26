package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/wbaker85/tacklebox/pkg/models"

	"github.com/google/uuid"
)

func (app *application) createBin(w http.ResponseWriter, r *http.Request) {
	var bin binInfo

	err := json.NewDecoder(r.Body).Decode(&bin)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	bin.ID = uuid.New().String()
	_, err = app.bins.Insert(bin.ID, bin.UserID)
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
	err = json.NewEncoder(w).Encode(&bin)
}

func (app *application) getBinHooks(w http.ResponseWriter, r *http.Request) {

}

func (app *application) deleteBin(w http.ResponseWriter, r *http.Request) {

}
