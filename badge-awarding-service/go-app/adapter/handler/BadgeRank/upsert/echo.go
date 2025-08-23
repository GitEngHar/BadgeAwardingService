package upsert

import (
	"context"
	"github.com/labstack/echo/v4"
	"hello-world/domain/management"
	"hello-world/infra/db/dynamo"
	usecase "hello-world/usecase/badgeRank"
	"net/http"
	"os"
)

type Handler struct{}

func NewBadgeRankHandler() *Handler {
	return &Handler{}
}

func (h Handler) Do(ctx context.Context, badgeRank management.BadgeDetailsByRank) (string, error) {
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
	BadgeID, err := uc.Do(ctx, badgeRank.BadgeRankID, badgeRank.BadgeName, badgeRank.Rank, badgeRank.Message, badgeRank.Effect, badgeRank.Reason)
	if err != nil {
		return "", echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return BadgeID, nil
}

func (h Handler) Hub(ctx echo.Context) error {
	var badgeRank management.BadgeDetailsByRank
	if err := ctx.Bind(&badgeRank); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	badgeRankID, err := h.Do(ctx.Request().Context(), badgeRank)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, badgeRankID)
}
