package Push

import (
	"context"
	"github.com/labstack/echo/v4"
	"hello-world/infra/queue"
	usecase "hello-world/usecase/push"
	"net/http"
)

type Publisher struct {
	UserName    string `json:"username"`
	Message     string `json:"message"`
	Address     string `json:"address"`
	MessageBody string `json:"message_body"`
}

type Handler struct{}

func NewPublisherHandler() *Handler {
	return &Handler{}
}

// Do TODO: echoの依存性を解放する
func (h Handler) Do(ctx context.Context, publisher Publisher) error {
	// repo実体化
	sqsConfig := queue.NewConfig(ctx)
	repo := queue.NewPublisher(*sqsConfig)

	// useCase実体化
	uc := usecase.NewPublishMessageUseCase(repo)

	// sqsにメッセージをパブリッシュ
	err := uc.Do(ctx, publisher.MessageBody, publisher.UserName, publisher.Address, publisher.Message)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (h Handler) Hub(ctx echo.Context) error {
	var publisher Publisher
	if err := ctx.Bind(&publisher); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return h.Do(ctx.Request().Context(), publisher)
}
