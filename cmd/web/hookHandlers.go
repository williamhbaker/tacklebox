package main

import (
	"bytes"
	"io"
	"net/http"
	"sync"

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

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		err := app.hooks.InsertOne(&id, buf.String())
		if err != nil {
			app.errorLog.Println(err)
		}
	}()

	go func() {
		defer wg.Done()
		err := app.hookRecords.InsertOne(binID, id.Hex())
		if err != nil {
			app.errorLog.Println(err)
		}
	}()

	wg.Wait()
	w.WriteHeader(http.StatusOK)
}

func (app *application) getHooks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from get hooks"))
}
