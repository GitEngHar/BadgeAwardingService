package User

import (
	"context"
	"github.com/labstack/echo/v4"
	"hello-world/domain/management"
	"hello-world/infra/db/dynamo"
	usecase "hello-world/usecase/user"
	"net/http"
)

type Handler struct{}

func NewUserHandler() *Handler {
	return &Handler{}
}

func (h Handler) Do(ctx context.Context, user management.UserDTO) error {
	// repo実体化
	dbConf := dynamo.NewConnectionDynamoDBForLocal()
	repo := dynamo.NewUserRepository(dbConf)
	// tableの作成
	if err := repo.CreateTable(ctx); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	// useCase実体化
	uc := usecase.NewUpsertUseCase(repo)
	err := uc.Do(ctx, user.Mail, user.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (h Handler) Hub(ctx echo.Context) error {
	var user management.UserDTO
	if err := ctx.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return h.Do(ctx.Request().Context(), user)
}
