package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.logRequest)
	dynamicMiddleware := alice.New(app.session.Enable, app.authenticate)

	mux := pat.New()

	mux.Post("/hook/:binID", alice.New(app.requireJSON).ThenFunc(app.postHook))

	mux.Get("/bin/:binID", dynamicMiddleware.Append(app.requireAuth, app.checkAccessForBin).ThenFunc(app.getBinHooks))
	mux.Del("/bin/:binID", dynamicMiddleware.Append(app.requireAuth, app.checkAccessForBin).ThenFunc(app.destroyBin))

	mux.Post("/bin", dynamicMiddleware.Append(app.requireJSON, app.requireAuth).ThenFunc(app.createBin))

	mux.Post("/user", dynamicMiddleware.Append(app.requireJSON).ThenFunc(app.createUser))
	mux.Post("/login", dynamicMiddleware.Append(app.requireJSON).ThenFunc(app.login))
	mux.Post("/logout", dynamicMiddleware.Append(app.requireAuth).ThenFunc(app.logout))

	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))

	return standardMiddleware.Then(mux)
}
