package get

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
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
func (h SubscriptionHandler) Do(ctx context.Context) ([]map[string]types.MessageAttributeValue, error) {
	snsConfig := sns.NewConfig(ctx)
	sqsConfig := queue.NewConfig(ctx)
	subRepo := sns.NewSubscription(snsConfig)
	pubRepo := queue.NewPublisher(*sqsConfig)
	uc := usecase.NewSubscriptionUseCase(subRepo, pubRepo)
	subscribedAwards, err := uc.Do(ctx)
	if err != nil {
		return nil, err
	}
	return subscribedAwards, nil
}

func (h SubscriptionHandler) Hub(ctx echo.Context) error {
	subscribedAwards, err := h.Do(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, subscribedAwards)
}
