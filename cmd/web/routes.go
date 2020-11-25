package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.logRequest)

	mux := pat.New()

	mux.Get("/hook/:binID", http.HandlerFunc(app.getHooks))
	mux.Post("/hook/:binID", http.HandlerFunc(app.postHook))
	mux.Post("/user", http.HandlerFunc(app.createUser))
	mux.Get("/", http.HandlerFunc(app.home))

	return standardMiddleware.Then(mux)
}
