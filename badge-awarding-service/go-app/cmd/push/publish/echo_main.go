package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	handler "hello-world/adapter/handler/Push"
	infra "hello-world/infra/echo"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	userHandler := handler.NewPublisherHandler()
	router := infra.NewRouter(e, userHandler, "POST")
	config := infra.NewEchoConfig("1323", router)
	echoRepo := infra.NewEchoRepository(config)
	echoRepo.Run(e)
}
