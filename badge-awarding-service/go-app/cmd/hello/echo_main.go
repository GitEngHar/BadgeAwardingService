package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	handler "hello-world/adapter/handler/Hello"
	infra "hello-world/infra/echo"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	helloHandler := handler.NewHelloHandler()
	router := infra.NewRouter(e, helloHandler, "GET")
	config := infra.NewEchoConfig("1323", router)
	repository := infra.NewEchoRepository(config)
	repository.Run(e)
}
