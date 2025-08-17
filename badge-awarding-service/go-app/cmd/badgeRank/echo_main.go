package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	getHandler "hello-world/adapter/handler/BadgeRank/get"
	upsertHandler "hello-world/adapter/handler/BadgeRank/upsert"
	infra "hello-world/infra/echo"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	getUserHandler := getHandler.NewBadgeRankHandler()
	upsertBadgeHandler := upsertHandler.NewBadgeHandler()
	router := infra.NewRouter(e, getUserHandler, upsertBadgeHandler, nil, nil)
	config := infra.NewEchoConfig("1323", router)
	echoRepo := infra.NewEchoRepository(config)
	echoRepo.Run(e)
}
