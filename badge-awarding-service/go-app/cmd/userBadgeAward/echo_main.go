package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	getHandler "hello-world/adapter/handler/Badgeaward/get"
	upsertHandler "hello-world/adapter/handler/Badgeaward/upsert"
	infra "hello-world/infra/echo"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	getUserBadgeAwardHandler := getHandler.NewUserHandler()
	upsertUserBadgeAwardHandler := upsertHandler.NewBadgeHandler()
	router := infra.NewRouter(e, getUserBadgeAwardHandler, upsertUserBadgeAwardHandler, nil, nil)
	config := infra.NewEchoConfig("1323", router)
	echoRepo := infra.NewEchoRepository(config)
	echoRepo.Run(e)
}
