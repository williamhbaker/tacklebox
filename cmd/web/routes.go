package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.logRequest)
	sessionMiddleware := alice.New(app.session.Enable, app.authenticate)

	mux := pat.New()

	mux.Post("/hook/:binID", alice.New(app.requireJSON).ThenFunc(app.postHook))
	mux.Del("/hook/:hookID", sessionMiddleware.Append(app.requireAuth, app.checkAccessForHook).ThenFunc(app.destroyHook))
	mux.Get("/hook/:hookID", sessionMiddleware.Append(app.requireAuth, app.checkAccessForHook).ThenFunc(app.getHook))

	mux.Get("/bin/:binID", sessionMiddleware.Append(app.requireAuth, app.checkAccessForBin).ThenFunc(app.getBinHooks))
	mux.Del("/bin/:binID", sessionMiddleware.Append(app.requireAuth, app.checkAccessForBin).ThenFunc(app.destroyBin))
	mux.Post("/bin", sessionMiddleware.Append(app.requireAuth).ThenFunc(app.createBin))

	mux.Get("/user/bins", sessionMiddleware.Append(app.requireAuth).ThenFunc(app.getUserBins))

	mux.Post("/user", sessionMiddleware.Append(app.requireJSON).ThenFunc(app.createUser))
	mux.Post("/login", sessionMiddleware.Append(app.requireJSON).ThenFunc(app.login))
	mux.Post("/logout", sessionMiddleware.Append(app.requireAuth).ThenFunc(app.logout))

	mux.Get("/", sessionMiddleware.ThenFunc(app.home))

	return standardMiddleware.Then(mux)
}
