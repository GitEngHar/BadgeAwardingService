package Push

import (
	"github.com/labstack/echo/v4"
	infra "hello-world/infra/echo"
	"hello-world/infra/sns"
	usecase "hello-world/usecase/push"
	"net/http"
)

type Subscription struct {
	Endpoint string `json:"endpoint"`
}

type SubscriptionHandler struct{}

func NewSubscriptionHandler() infra.Handler {
	return SubscriptionHandler{}
}

func (h SubscriptionHandler) Do(ctx echo.Context) error {
	var subscription Subscription
	if err := ctx.Bind(&subscription); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// repo実体化
	snsConfig := sns.NewConfig(ctx.Request().Context())
	repo := sns.NewSubscription(snsConfig)

	// useCase実体化
	uc := usecase.NewSubscriptionUseCase(repo)

	// sqsにメッセージをパブリッシュ
	err := uc.Do(ctx.Request().Context(), subscription.Endpoint)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
