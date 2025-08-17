//go:build wireinject
// +build wireinject

package dipendency_injection

//go:generate go run github.com/google/wire/cmd/wire@latest

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"hello-world/domain/management"
	"hello-world/infra/lambda"
	"net/http"
)
import "github.com/google/wire"
import badgeGetHandler "hello-world/adapter/handler/Badge/get"
import badgePostHandler "hello-world/adapter/handler/Badge/upsert"

var Set = wire.NewSet(
	//wire.Struct(new(badgeGetHandler.Handler), "*"), // wireが自動でfieldを埋めてくれる
	badgeGetHandler.NewBadgeHandler,
	badgePostHandler.NewBadgeHandler,
	NewBadgeGetController, NewBadgePostController,
	NewGetController, NewPostController,
	lambda.NewRouter,
)

type BadgeGetController struct{ BadgeGet badgeGetHandler.Handler }

func NewBadgeGetController(badgeGet *badgeGetHandler.Handler) *BadgeGetController {
	return &BadgeGetController{BadgeGet: *badgeGet}
}
func (c *BadgeGetController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.URL.Query().Get("id")
	value, err := c.BadgeGet.Do(ctx, id)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"val": err.Error()})
	}
	writeJSON(w, http.StatusOK, value)

}

type BadgePostController struct{ BadgePost badgePostHandler.Handler }

func NewBadgePostController(badgePost *badgePostHandler.Handler) *BadgePostController {
	return &BadgePostController{
		BadgePost: *badgePost,
	}
}

func (c *BadgePostController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var badgeDTO management.BadgeDTO

	if err := json.NewDecoder(r.Body).Decode(&badgeDTO); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"val": err.Error()})
		return
	}
	badge, err := management.NewBadge(badgeDTO.ID, badgeDTO.Name, badgeDTO.Reason)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"val": err.Error()})
	}
	value, err := c.BadgePost.Do(ctx, *badge)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"val": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, value)
}

type GetController struct {
	BadgeGetController *BadgeGetController
}

func NewGetController(badgeGetController *BadgeGetController) lambda.GetController {
	return &GetController{
		BadgeGetController: badgeGetController,
	}
}

func (g GetController) GetController() map[string]http.Handler {
	getController := map[string]http.Handler{}
	getController["/badge"] = g.BadgeGetController
	return getController
}

type PostController struct {
	BadgePostController *BadgePostController
}

func NewPostController(badgePostController *BadgePostController) lambda.PostController {
	return &PostController{BadgePostController: badgePostController}
}

func (p PostController) PostController() map[string]http.Handler {
	postController := map[string]http.Handler{}
	postController["/badge"] = p.BadgePostController
	return postController
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func InitializeRouter() (*chi.Mux, error) {
	wire.Build(Set)
	return nil, nil
}
