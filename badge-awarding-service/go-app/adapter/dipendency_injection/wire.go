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
import userGetHandler "hello-world/adapter/handler/User/get"
import userPostHandler "hello-world/adapter/handler/User/upsert"
import badgeAwardGetHandler "hello-world/adapter/handler/BadgeAward/get"
import badgeAwardPostHandler "hello-world/adapter/handler/BadgeAward/upsert"
import badgeRankGetHandler "hello-world/adapter/handler/BadgeRank/get"
import badgeRankPostHandler "hello-world/adapter/handler/BadgeRank/upsert"
import publishAwardGetHandler "hello-world/adapter/handler/PublishAward/get"
import subscribeAwardGetHandler "hello-world/adapter/handler/SubscribeAward/get"
import badgeGetHandler "hello-world/adapter/handler/Badge/get"
import badgePostHandler "hello-world/adapter/handler/Badge/upsert"

var Set = wire.NewSet(
	//wire.Struct(new(badgeGetHandler.Handler), "*"), // wireが自動でfieldを埋めてくれる
	userGetHandler.NewUserHandler,
	userPostHandler.NewUserHandler,
	badgeAwardGetHandler.NewBadgeAwardHandler,
	badgeAwardPostHandler.NewBadgeAwardHandler,
	badgeRankGetHandler.NewBadgeRankHandler,
	badgeRankPostHandler.NewBadgeRankHandler,
	publishAwardGetHandler.NewPublisherHandler,
	subscribeAwardGetHandler.NewSubscriptionHandler,
	badgeGetHandler.NewBadgeHandler,
	badgePostHandler.NewBadgeHandler,
	NewBadgeGetController, NewUserGetController, NewSubscribeAwardGetController, NewPublishAwardGetController, NewBadgeRankGetController, NewBadgeAwardGetController,
	NewBadgePostController, NewBadgeRankPostController, NewBadgeAwardPostController, NewUserPostController,
	NewGetController, NewPostController,
	lambda.NewRouter,
)

type BadgeGetController struct{ BadgeGet badgeGetHandler.Handler }

func NewBadgeGetController(
	badgeGet *badgeGetHandler.Handler,
) *BadgeGetController {
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

/**
* GET Method
 */
// UserGetController ユーザー情報を取得する
type UserGetController struct{ UserGet userGetHandler.Handler }

func NewUserGetController(
	userGet *userGetHandler.Handler,
) *UserGetController {
	return &UserGetController{UserGet: *userGet}
}
func (c *UserGetController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.URL.Query().Get("id")
	value, err := c.UserGet.Do(ctx, id)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"val": err.Error()})
	}
	writeJSON(w, http.StatusOK, value)

}

// BadgeAwardGetController 成果通知(バッチ付与・更新)する対象者を抽出する
type BadgeAwardGetController struct{ BadgeAwardGet badgeAwardGetHandler.Handler }

func NewBadgeAwardGetController(
	badgeAwardGet *badgeAwardGetHandler.Handler,
) *BadgeAwardGetController {
	return &BadgeAwardGetController{BadgeAwardGet: *badgeAwardGet}
}
func (c *BadgeAwardGetController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	value, err := c.BadgeAwardGet.Do(ctx)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"val": err.Error()})
	}
	writeJSON(w, http.StatusOK, value)

}

// BadgeRankGetController 指定したIDのバッチRank情報を取得する
type BadgeRankGetController struct{ BadgeRankGet badgeRankGetHandler.Handler }

func NewBadgeRankGetController(
	badgeRankGet *badgeRankGetHandler.Handler,
) *BadgeRankGetController {
	return &BadgeRankGetController{BadgeRankGet: *badgeRankGet}
}
func (c *BadgeRankGetController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.URL.Query().Get("id")
	value, err := c.BadgeRankGet.Do(ctx, id)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"val": err.Error()})
	}
	writeJSON(w, http.StatusOK, value)

}

// PublishAwardGetController 生活うち対象のユーザーを抽出し通知対象者としてキューにパブリッシュする
type PublishAwardGetController struct {
	PublishAwardGet publishAwardGetHandler.Handler
}

func NewPublishAwardGetController(
	publishAwardGet *publishAwardGetHandler.Handler,
) *PublishAwardGetController {
	return &PublishAwardGetController{PublishAwardGet: *publishAwardGet}
}
func (c *PublishAwardGetController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := c.PublishAwardGet.Do(ctx)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"val": err.Error()})
	}
	writeJSON(w, http.StatusOK, "success")

}

// PublishAwardGetController 配信したいユーザー情報をキューに入れておく
type SubscribeAwardGetController struct {
	SubscribeAwardGet subscribeAwardGetHandler.SubscriptionHandler
}

func NewSubscribeAwardGetController(
	subscribeAwardGet *subscribeAwardGetHandler.SubscriptionHandler,
) *SubscribeAwardGetController {
	return &SubscribeAwardGetController{SubscribeAwardGet: *subscribeAwardGet}
}
func (c *SubscribeAwardGetController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	value, err := c.SubscribeAwardGet.Do(ctx)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"val": err.Error()})
	}
	writeJSON(w, http.StatusOK, value)

}

/**
* POST Method
 */
type UserPostController struct{ UserPost userPostHandler.Handler }

func NewUserPostController(userPost *userPostHandler.Handler) *UserPostController {
	return &UserPostController{
		UserPost: *userPost,
	}
}
func (c *UserPostController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var userDTO management.UserDTO

	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"val": err.Error()})
		return
	}
	value, err := c.UserPost.Do(ctx, userDTO)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"val": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, value)
}

type BadgeAwardPostController struct{ BadgeAwardPost badgeAwardPostHandler.Handler }

func NewBadgeAwardPostController(badgeAwardPost *badgeAwardPostHandler.Handler) *BadgeAwardPostController {
	return &BadgeAwardPostController{
		BadgeAwardPost: *badgeAwardPost,
	}
}
func (c *BadgeAwardPostController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var userBadgeAwardDTO management.UserBadgeAwardDTO

	if err := json.NewDecoder(r.Body).Decode(&userBadgeAwardDTO); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"val": err.Error()})
		return
	}
	userBadgeAward, err := management.NewUserBadgeAward(userBadgeAwardDTO.UserBadgeAwardID, userBadgeAwardDTO.UserID, userBadgeAwardDTO.BadgeRankID)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"val": err.Error()})
	}
	value, err := c.BadgeAwardPost.Do(ctx, *userBadgeAward)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"val": err.Error()})
		return
	}
	writeJSON(w, http.StatusOK, value)
}

type BadgeRankPostController struct{ BadgeRankPost badgeRankPostHandler.Handler }

func NewBadgeRankPostController(badgeRankPost *badgeRankPostHandler.Handler) *BadgeRankPostController {
	return &BadgeRankPostController{BadgeRankPost: *badgeRankPost}
}
func (c *BadgeRankPostController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var badgeDetailsByRankDTO management.BadgeDetailsByRankDTO

	if err := json.NewDecoder(r.Body).Decode(&badgeDetailsByRankDTO); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"val": err.Error()})
		return
	}
	userBadgeAward, err := management.NewBadgeDetailsByRank(
		badgeDetailsByRankDTO.BadgeRankID,
		badgeDetailsByRankDTO.BadgeName,
		badgeDetailsByRankDTO.Message,
		badgeDetailsByRankDTO.Effect,
		badgeDetailsByRankDTO.Reason,
		badgeDetailsByRankDTO.Rank,
	)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"val": err.Error()})
	}
	value, err := c.BadgeRankPost.Do(ctx, *userBadgeAward)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"val": err.Error()})
		return
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
	BadgeGetController          *BadgeGetController
	UserGetController           *UserGetController
	BadgeRankGetController      *BadgeRankGetController
	BadgeAwardGetController     *BadgeAwardGetController
	PublishAwardGetController   *PublishAwardGetController
	SubscribeAwardGetController *SubscribeAwardGetController
}

func NewGetController(
	badgeGetController *BadgeGetController,
	userGetController *UserGetController,
	badgeRankGetController *BadgeRankGetController,
	badgeAwardGetController *BadgeAwardGetController,
	publishAwardGetController *PublishAwardGetController,
	subscribeAwardGetController *SubscribeAwardGetController) lambda.GetController {
	return &GetController{
		BadgeGetController:          badgeGetController,
		UserGetController:           userGetController,
		BadgeRankGetController:      badgeRankGetController,
		BadgeAwardGetController:     badgeAwardGetController,
		PublishAwardGetController:   publishAwardGetController,
		SubscribeAwardGetController: subscribeAwardGetController,
	}
}

func (g GetController) GetController() map[string]http.Handler {
	getController := map[string]http.Handler{}
	getController["/badge"] = g.BadgeGetController
	getController["/user"] = g.UserGetController
	getController["/rank"] = g.BadgeRankGetController
	getController["/award"] = g.BadgeAwardGetController
	getController["/publish"] = g.PublishAwardGetController
	getController["/subscribe"] = g.SubscribeAwardGetController
	return getController
}

type PostController struct {
	BadgePostController      *BadgePostController
	BadgeRankPostController  *BadgeRankPostController
	BadgeAwardPostController *BadgeAwardPostController
	UserPostController       *UserPostController
}

func NewPostController(
	badgePostController *BadgePostController,
	badgeRankPostController *BadgeRankPostController,
	badgeAwardPostController *BadgeAwardPostController,
	userPostController *UserPostController) lambda.PostController {
	return &PostController{
		BadgePostController:      badgePostController,
		BadgeRankPostController:  badgeRankPostController,
		BadgeAwardPostController: badgeAwardPostController,
		UserPostController:       userPostController,
	}
}

func (p PostController) PostController() map[string]http.Handler {
	postController := map[string]http.Handler{}
	postController["/badge"] = p.BadgePostController
	postController["/rank"] = p.BadgeRankPostController
	postController["/award"] = p.BadgeAwardPostController
	postController["/user"] = p.UserPostController
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
