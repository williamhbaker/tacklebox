package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"sync"

	"github.com/wbaker85/tacklebox/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (app *application) postHook(w http.ResponseWriter, r *http.Request) {
	binID := r.URL.Query().Get(":binID")
	if binID == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var buf bytes.Buffer
	io.Copy(&buf, r.Body)
	bytes := buf.Bytes()

	if !validJSONBytes(bytes) {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	id := primitive.NewObjectID()

	_, err := app.hookRecords.Insert(binID, id.Hex())
	if err != nil {
		if err == models.ErrInvalidBin {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errJSON{"invalid bin"})
			return
		}
		app.serverError(w, err)
		return
	}

	_, err = app.hooks.Insert(buf.String(), &id)
	if err != nil {
		app.hookRecords.Destroy(id.Hex())
		app.serverError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(infoJSON{"success"})
}

func (app *application) getHook(w http.ResponseWriter, r *http.Request) {
	hookID := r.URL.Query().Get(":hookID")

	record, err := app.hookRecords.GetOne(hookID)
	if err != nil {
		if err == models.ErrInvalidHook {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(errJSON{"hook not found"})
			return
		}
		app.serverError(w, err)
		return
	}

	doc, err := app.hooks.GetOne(hookID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	output := assembleHookJSON([]*models.HookRecord{record}, []*models.HookDocument{doc})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output[0])
}

func (app *application) destroyHook(w http.ResponseWriter, r *http.Request) {
	hookID := r.URL.Query().Get(":hookID")

	var foundErr error
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		err := app.hookRecords.Destroy(hookID)
		if err != nil {
			foundErr = err
		}
	}()

	go func() {
		defer wg.Done()
		_, err := app.hooks.DestroyOne(hookID)
		if err != nil {
			foundErr = err
		}
	}()

	wg.Wait()

	if foundErr != nil {
		app.serverError(w, foundErr)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(infoJSON{"success"})
}
