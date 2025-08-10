package PushAward

import (
	"github.com/labstack/echo/v4"
	"hello-world/infra/queue"
	"hello-world/infra/sns"
	usecase "hello-world/usecase/push"
	"net/http"
)

type SubscriptionHandler struct{}

func NewSubscriptionHandler() *SubscriptionHandler {
	return &SubscriptionHandler{}
}

// Do 配信対象のメッセージをユーザーへ全て送信しする
func (h SubscriptionHandler) Do(ctx echo.Context) error {
	snsConfig := sns.NewConfig(ctx.Request().Context())
	sqsConfig := queue.NewConfig(ctx.Request().Context())
	subRepo := sns.NewSubscription(snsConfig)
	pubRepo := queue.NewPublisher(*sqsConfig)
	uc := usecase.NewSubscriptionUseCase(subRepo, pubRepo)
	err := uc.Do(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
