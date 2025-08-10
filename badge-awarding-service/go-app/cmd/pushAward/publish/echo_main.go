package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	gethandler "hello-world/adapter/handler/PublishAward/get"
	infra "hello-world/infra/echo"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	getPublishHandler := gethandler.NewPublisherHandler()
	router := infra.NewRouter(e, getPublishHandler, nil, nil, nil)
	config := infra.NewEchoConfig("1323", router)
	echoRepo := infra.NewEchoRepository(config)
	echoRepo.Run(e)
}
