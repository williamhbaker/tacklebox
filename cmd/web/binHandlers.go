package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"sync"
	"time"

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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(infoJSON{binID})
}

func (app *application) getBinHooks(w http.ResponseWriter, r *http.Request) {
	binID := r.URL.Query().Get(":binID")

	records, err := app.hookRecords.Get(binID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	docIDs := docIDsFromRecords(records)

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

	records, err := app.hookRecords.Get(binID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	docIDs := docIDsFromRecords(records)

	var foundErr error
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		_, err = app.hooks.DestroyMany(docIDs)
		if err != nil {
			foundErr = err
		}
	}()

	go func() {
		defer wg.Done()
		_, err = app.bins.Destroy(binID)
		if err != nil {
			foundErr = err
		}
	}()

	wg.Wait()

	if foundErr != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(infoJSON{"success"})
}

func (app *application) getUserBins(w http.ResponseWriter, r *http.Request) {
	userID := app.session.GetInt(r, "authenticatedUserID")

	bins, err := app.bins.GetUserBins(userID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	type resultBin struct {
		ID      string
		Created time.Time
	}

	res := make([]*resultBin, len(bins))
	for idx := range res {
		res[idx] = &resultBin{
			ID:      bins[idx].ID,
			Created: bins[idx].Created,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
