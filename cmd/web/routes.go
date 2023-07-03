package main

import (
	"net/http"

	"git.sr.ht/~edwlarkey/cardamom/ui"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	fileServer := http.FileServer(http.FS(ui.Files))
	router.Handler(http.MethodGet, "/static/*filepath", fileServer)

	router.HandlerFunc(http.MethodGet, "/ping", ping)

	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)

	// Public routes
	router.Handler(http.MethodGet, "/user/signup", dynamic.ThenFunc(app.userSignup))
	router.Handler(http.MethodPost, "/user/signup", dynamic.ThenFunc(app.userSignupPost))
	router.Handler(http.MethodGet, "/user/login", dynamic.ThenFunc(app.userLogin))
	router.Handler(http.MethodPost, "/user/login", dynamic.ThenFunc(app.userLoginPost))

	// Routes protected via auth
	protected := dynamic.Append(app.requireAuthentication)

	router.Handler(http.MethodGet, "/", protected.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/bookmark/view/:id", protected.ThenFunc(app.bookmarkView))
	router.Handler(http.MethodPost, "/bookmark/markunread", protected.ThenFunc(app.bookmarkMarkUnread))
	router.Handler(http.MethodPost, "/bookmark/markread", protected.ThenFunc(app.bookmarkMarkRead))
	router.Handler(http.MethodGet, "/bookmark/create", protected.ThenFunc(app.bookmarkCreate))
	router.Handler(http.MethodPost, "/bookmark/create", protected.ThenFunc(app.bookmarkCreatePost))

	router.Handler(http.MethodPost, "/user/logout", protected.ThenFunc(app.userLogoutPost))

	// router.Handler(http.MethodGet, "/tag/view/:id", protected.ThenFunc(app.tagView))
	router.Handler(http.MethodGet, "/tag/list", protected.ThenFunc(app.tagList))
	router.Handler(http.MethodPost, "/tag/create", protected.ThenFunc(app.tagCreatePost))

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	return standard.Then(router)
}
