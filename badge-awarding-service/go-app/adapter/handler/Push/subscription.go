package Push

import (
	"github.com/labstack/echo/v4"
	infra "hello-world/infra/echo"
	"hello-world/infra/queue"
	"hello-world/infra/sns"
	usecase "hello-world/usecase/push"
	"net/http"
)

type SubscriptionHandler struct{}

func NewSubscriptionHandler() infra.Handler {
	return SubscriptionHandler{}
}

// Do TODO: echoの依存性を解放する
func (h SubscriptionHandler) Do(ctx echo.Context) error {
	// repo実体化
	snsConfig := sns.NewConfig(ctx.Request().Context())
	sqsConfig := queue.NewConfig(ctx.Request().Context())
	subRepo := sns.NewSubscription(snsConfig)
	pubRepo := queue.NewPublisher(*sqsConfig)
	// useCase実体化
	uc := usecase.NewSubscriptionUseCase(subRepo, pubRepo)

	// sqsにメッセージをパブリッシュ
	err := uc.Do(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
