package main

import (
	"bytes"
	"fmt"
	"io"
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

	var bodyBytes bytes.Buffer
	io.Copy(&bodyBytes, r.Body)

	if !validJSONBytes(bodyBytes.Bytes()) {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// do something with the resulting string of JSON, maybe put it in a database
	jsonString := bodyBytes.String()

	binID := r.URL.Query().Get(":binID")
	if binID == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Bin ID: %s\n", binID)))
	w.Write([]byte(fmt.Sprintf("Posted JSON: %s", jsonString)))
}

func (app *application) getHooks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from get hooks"))
}
