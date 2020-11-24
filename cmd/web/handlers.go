package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from home"))
}

func (app *application) postHook(w http.ResponseWriter, r *http.Request) {

	contentType := r.Header.Get("Content-Type")

	if contentType != "application/json" {
		msg := "Content-Type header is not application/json"
		http.Error(w, msg, http.StatusUnsupportedMediaType)
		return
	}

	var rb interface{}
	err := json.NewDecoder(r.Body).Decode(&rb)
	if err != nil {
		app.errorLog.Println(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	binID := r.URL.Query().Get(":binID")
	if binID == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v", rb)))
}

func (app *application) getHooks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from get hooks"))
}
