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
	mux.Post("/hook/:binID", dynamicMiddleware.Append(app.requireJSON).ThenFunc(app.postHook))
	mux.Post("/user", dynamicMiddleware.Append(app.requireJSON).ThenFunc(app.createUser))
	mux.Post("/login", dynamicMiddleware.Append(app.requireJSON).ThenFunc(app.login))
	mux.Post("/bin", dynamicMiddleware.Append(app.requireJSON).ThenFunc(app.createBin))
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))

	return standardMiddleware.Then(mux)
}
