package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"sync"

	"github.com/wbaker85/tacklebox/pkg/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type newUser struct {
	Email    string
	Password string
}

type emailAlreadyRegistered struct {
	Error string `json:"error"`
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from home"))
}

func (app *application) postHook(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")

	if contentType != "application/json" {
		app.clientError(w, http.StatusUnsupportedMediaType)
		return
	}

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

func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")

	if contentType != "application/json" {
		app.clientError(w, http.StatusUnsupportedMediaType)
		return
	}

	var u newUser

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
	}

	err = app.users.Insert(u.Email, u.Password)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotAcceptable)
			json.NewEncoder(w).Encode(emailAlreadyRegistered{"Email already registered"})

		} else {
			app.serverError(w, err)
		}
	}
}
