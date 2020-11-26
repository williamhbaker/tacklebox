package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/wbaker85/tacklebox/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from home"))
}

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

func (app *application) getHooks(w http.ResponseWriter, r *http.Request) {
	binID := r.URL.Query().Get(":binID")
	if binID == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

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

	output := []models.HookData{}

	for idx := range records {
		h := models.HookData{}
		h.ID = records[idx].HookID
		h.BinID = records[idx].BinID
		h.Created = records[idx].Created

		var content string
		for idx := 0; idx < len(hooks); idx++ {
			if hooks[idx].ID.Hex() == h.ID {
				content = hooks[idx].Content
				break
			}
		}

		h.Content = content
		output = append(output, h)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
