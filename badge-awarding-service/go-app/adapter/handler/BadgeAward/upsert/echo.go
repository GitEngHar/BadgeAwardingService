package upsert

import (
	"context"
	"github.com/labstack/echo/v4"
	"hello-world/domain/management"
	"hello-world/infra/db/dynamo"
	usecase "hello-world/usecase/userBadgeAward"
	"net/http"
	"os"
)

type Handler struct{}

func NewBadgeAwardHandler() *Handler {
	return &Handler{}
}

func (h Handler) Do(ctx context.Context, userBadgeAward management.UserBadgeAward) (string, error) {
	// repo実体化
	dbConf := dynamo.NewConnectionDynamoDBForLocal()
	repo := dynamo.NewUserRepository(dbConf)
	// PROD環境ではCFが作成するためアプリでは作成しない
	if os.Getenv("ENVIRONMENT") != "PROD" {
		if err := repo.CreateTable(ctx); err != nil {
			return "", err
		}
	}
	// useCase実体化
	uc := usecase.NewUpsertUseCase(repo)
	BadgeID, err := uc.Do(ctx, userBadgeAward.UserBadgeAwardID, userBadgeAward.UserID, userBadgeAward.BadgeRankID)
	if err != nil {
		return "", echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return BadgeID, nil
}

func (h Handler) Hub(ctx echo.Context) error {
	var userBadgeAward management.UserBadgeAward
	if err := ctx.Bind(&userBadgeAward); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userBadgeAwardID, err := h.Do(ctx.Request().Context(), userBadgeAward)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, userBadgeAwardID)
}
