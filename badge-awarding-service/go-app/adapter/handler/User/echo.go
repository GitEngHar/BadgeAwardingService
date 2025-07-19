package User

import (
	"github.com/labstack/echo/v4"
	"hello-world/infra/db/dynamo"
	infra "hello-world/infra/echo"
	usecase "hello-world/usecase/user"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Mail string `json:"mail"`
}

type Handler struct{}

func NewUserHandler() infra.Handler {
	return Handler{}
}

func (h Handler) Do(ctx echo.Context) error {
	var user User
	if err := ctx.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// repo実体化
	dbConf := dynamo.NewConnectionDynamoDBForLocal()
	repo := dynamo.NewUserRepository(dbConf)
	// tableの作成
	if err := repo.CreateTable(ctx.Request().Context()); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	// useCase実体化
	uc := usecase.NewUpsertUseCase(repo)
	err := uc.Do(ctx.Request().Context(), user.Mail, user.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
