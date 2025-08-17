package lambda

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type GetController interface {
	GetController() map[string]http.Handler
}
type PostController interface {
	PostController() map[string]http.Handler
}

func NewRouter(getController GetController, postController PostController) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("ok")) })
	r = setGetRouter(r, getController.GetController())
	r = setPostRouter(r, postController.PostController())
	return r
}

func setPostRouter(router *chi.Mux, routes map[string]http.Handler) *chi.Mux {
	for path := range routes {
		router.Method("POST", path, routes[path])
	}
	return router
}

func setGetRouter(router *chi.Mux, routes map[string]http.Handler) *chi.Mux {
	for path := range routes {
		router.Method("GET", path, routes[path])
	}
	return router
}
