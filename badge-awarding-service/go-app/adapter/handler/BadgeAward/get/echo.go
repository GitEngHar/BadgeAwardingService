package get

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/labstack/echo/v4"
	"hello-world/infra/db/dynamo"
	badgeAwardUsecase "hello-world/usecase/userBadgeAward"
	"net/http"
)

type Handler struct{}

func NewBadgeAwardHandler() *Handler {
	return &Handler{}
}

func (h Handler) Do(ctx context.Context) ([]map[string]types.AttributeValue, error) {
	// repo実体化
	dbConf := dynamo.NewConnectionDynamoDBForLocal()
	dynamodbRepo := dynamo.NewUserRepository(dbConf)
	// tableの作成
	if err := dynamodbRepo.CreateTable(ctx); err != nil {
		return nil, err
	}
	// useCase実体化
	badgeAwardGetUseCase := badgeAwardUsecase.NewGetUseCase(dynamodbRepo)
	// badgeAward, err := uc.Do(ctx, id)
	badgeAward, err := badgeAwardGetUseCase.GetAwardTargetUserByDate(ctx)
	if err != nil {
		return nil, err
	}
	return badgeAward, nil
}

func (h Handler) Hub(ctx echo.Context) error {
	badgeAward, err := h.Do(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, badgeAward)
}
