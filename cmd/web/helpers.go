package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
)

type userInfo struct {
	Email    string
	Password string
}

type binInfo struct {
	ID     string
	UserID int
}

type errJSON struct {
	Error string `json:"error"`
}

type infoJSON struct {
	Info string `json:"message"`
}

type userIDJSON struct {
	ID string `json:"id"`
}

func validJSONBytes(b []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(b, &js) == nil
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
