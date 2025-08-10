package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	getHandler "hello-world/adapter/handler/SubscribeEmailAward/get"
	infra "hello-world/infra/echo"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	getSubscribeEmail := getHandler.NewSubscriptionHandler()
	router := infra.NewRouter(e, getSubscribeEmail, nil, nil, nil)
	config := infra.NewEchoConfig("1323", router)
	echoRepo := infra.NewEchoRepository(config)
	echoRepo.Run(e)
}
