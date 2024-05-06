package main

import "net/http"

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", app.neuter(fileServer)))

	mux.HandleFunc("/", app.home)

	mux.HandleFunc("/user/signup", app.userSignup)
	mux.HandleFunc("/user/login", app.userLogin)
	mux.HandleFunc("/user/logout", app.requireAuth(app.userLogout))

	mux.HandleFunc("/user/profile", app.requireAuth(app.userProfile)) // not finished
	mux.HandleFunc("/user/favourites", app.requireAuth(app.userFavourites))

	mux.HandleFunc("/post/create", app.requireAuth(app.postCreate))
	mux.HandleFunc("/post/view", app.postView)

	mux.HandleFunc("/post/like", app.requireAuth(app.postLike))
	mux.HandleFunc("/post/dislike", app.requireAuth(app.postDislike))

	mux.HandleFunc("/post/comment", app.requireAuth(app.postComment))
	mux.HandleFunc("/post/commentLike", app.requireAuth(app.commentLike))
	mux.HandleFunc("/post/commentDislike", app.requireAuth(app.commentDislike))

	return app.recoverPanic(app.AuthMiddleware(app.logRequest(secureHeaders(mux))))
}
