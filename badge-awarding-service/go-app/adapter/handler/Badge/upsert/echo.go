package upsert

import (
	"context"
	"github.com/labstack/echo/v4"
	"hello-world/domain/management"
	"hello-world/infra/db/dynamo"
	usecase "hello-world/usecase/badge"
	"net/http"
	"os"
)

type Handler struct{}

func NewBadgeHandler() *Handler {
	return &Handler{}
}

func (h Handler) Do(ctx context.Context, badge management.Badge) (string, error) {
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
	BadgeID, err := uc.Do(ctx, badge.ID, badge.Name, badge.Reason)
	if err != nil {
		return "", echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return BadgeID, nil
}

func (h Handler) Hub(ctx echo.Context) error {
	var badge management.Badge
	if err := ctx.Bind(&badge); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	badgeID, err := h.Do(ctx.Request().Context(), badge)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, badgeID)
}
