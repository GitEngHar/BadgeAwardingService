package get

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/labstack/echo/v4"
	"hello-world/infra/db/dynamo"
	usecase "hello-world/usecase/userBadgeAward"
	"net/http"
)

type Handler struct{}

func NewBadgeAwardHandler() *Handler {
	return &Handler{}
}

func (h Handler) Do(ctx context.Context, id string) ([]map[string]types.AttributeValue, error) {
	// repo実体化
	dbConf := dynamo.NewConnectionDynamoDBForLocal()
	repo := dynamo.NewUserRepository(dbConf)
	// tableの作成
	if err := repo.CreateTable(ctx); err != nil {
		return nil, err
	}
	// useCase実体化
	uc := usecase.NewGetUseCase(repo)
	// badgeAward, err := uc.Do(ctx, id)
	badgeAward, err := uc.AwardTargetUserByDate(ctx)
	if err != nil {
		return nil, err
	}
	return badgeAward, nil
}

func (h Handler) Hub(ctx echo.Context) error {
	id := ctx.QueryParam("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "id not found")
	}
	badgeAward, err := h.Do(ctx.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, badgeAward)
}
