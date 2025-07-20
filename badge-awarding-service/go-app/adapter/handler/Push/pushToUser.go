package Push

import (
	"github.com/labstack/echo/v4"
	infra "hello-world/infra/echo"
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

func NewPublisherHandler() infra.Handler {
	return Handler{}
}

func (h Handler) Do(ctx echo.Context) error {
	var publisher Publisher
	if err := ctx.Bind(&publisher); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// repo実体化
	sqsConfig := queue.NewConnection(ctx.Request().Context())
	repo := queue.NewMailQueue(*sqsConfig)

	// useCase実体化
	uc := usecase.NewToUserUseCase(repo)

	// sqsにメッセージをパブリッシュ
	err := uc.Do(ctx.Request().Context(), publisher.MessageBody, publisher.UserName, publisher.Address, publisher.Message)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
