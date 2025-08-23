package upsert

import (
	"context"
	"github.com/labstack/echo/v4"
	"hello-world/domain/management"
	"hello-world/infra/db/dynamo"
	usecase "hello-world/usecase/user"
	"net/http"
	"os"
)

type Handler struct{}

func NewUserHandler() *Handler {
	return &Handler{}
}

func (h Handler) Do(ctx context.Context, user management.UserDTO) (string, error) {
	// repo実体化
	dbConf := dynamo.NewConnectionDynamoDBForLocal()
	repo := dynamo.NewUserRepository(dbConf)
	// PROD環境ではCFが作成するためアプリでは作成しない
	if os.Getenv("ENVIRONMENT") != "PROD" {
		if err := repo.CreateTable(ctx); err != nil {
			return "", echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	// useCase実体化
	uc := usecase.NewUpsertUseCase(repo)
	userID, err := uc.Do(ctx, user.ID, user.Mail, user.Name)
	if err != nil {
		return "", echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return userID, nil
}

func (h Handler) Hub(ctx echo.Context) error {
	var user management.UserDTO
	if err := ctx.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userID, err := h.Do(ctx.Request().Context(), user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, userID)
}
