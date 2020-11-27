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
	binID := r.URL.Query().Get(":binID")

	records, err := app.hookRecords.Get(binID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	var docIDs []string
	for _, record := range records {
		docIDs = append(docIDs, record.HookID)
	}

	hooks, err := app.hooks.GetMany(docIDs)
	if err != nil {
		app.serverError(w, err)
		return
	}

	output := assembleHookJSON(records, hooks)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func (app *application) destroyBin(w http.ResponseWriter, r *http.Request) {
	binID := r.URL.Query().Get(":binID")

	_, err := app.bins.Destroy(binID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(infoJSON{"success"})
}
