package rest

import (
	"personal-website-server/env"

	"github.com/gorilla/mux"
)

type RestHandler struct {
	env *env.Env
}

func SetupRoutes(env *env.Env) *mux.Router {
	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/api/").Subrouter()

	ch := newContactHandler(env)
	apiRouter.HandleFunc("/contact", ch.receiveContact).Methods("POST")

	vh := newVideosHandler(env)
	apiRouter.HandleFunc("/videos", vh.login).Methods("POST")

	return r
}
