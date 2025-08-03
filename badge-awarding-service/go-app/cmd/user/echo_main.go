package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	get_handler "hello-world/adapter/handler/User/get"
	upsert_handler "hello-world/adapter/handler/User/upsert"
	infra "hello-world/infra/echo"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	getUserHandler := get_handler.NewUserHandler()
	upsertUserHandler := upsert_handler.NewUserHandler()
	router := infra.NewRouter(e, getUserHandler, upsertUserHandler, nil, nil)
	config := infra.NewEchoConfig("1323", router)
	echoRepo := infra.NewEchoRepository(config)
	echoRepo.Run(e)
}
