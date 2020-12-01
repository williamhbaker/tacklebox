package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/wbaker85/tacklebox/pkg/models"
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

func assembleHookJSON(r []*models.HookRecord, d []*models.HookDocument) []models.HookData {
	output := []models.HookData{}

	for idx := range r {
		h := models.HookData{}
		h.ID = r[idx].HookID
		h.BinID = r[idx].BinID
		h.Created = r[idx].Created

		var content string
		for idx := 0; idx < len(d); idx++ {
			if d[idx].ID.Hex() == h.ID {
				content = d[idx].Content
				break
			}
		}

		h.Content = content
		output = append(output, h)
	}

	return output
}

func docIDsFromRecords(recs []*models.HookRecord) []string {
	var docIDs []string
	for _, r := range recs {
		docIDs = append(docIDs, r.HookID)
	}
	return docIDs
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

func (app *application) isAuthenticated(r *http.Request) bool {
	isAuthenticated, ok := r.Context().Value(contextKeyIsAuthenticated).(bool)
	if !ok {
		return false
	}

	return isAuthenticated
}
