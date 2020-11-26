package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.logRequest)
	dynamicMiddleware := alice.New(app.session.Enable)

	mux := pat.New()

	mux.Get("/hook/:binID", dynamicMiddleware.ThenFunc(app.getHooks))
	mux.Post("/hook/:binID", dynamicMiddleware.ThenFunc(app.postHook))
	mux.Post("/user", dynamicMiddleware.ThenFunc(app.createUser))
	mux.Post("/login", dynamicMiddleware.ThenFunc(app.login))
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))

	return standardMiddleware.Then(mux)
}
