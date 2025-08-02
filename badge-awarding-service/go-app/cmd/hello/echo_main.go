package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	handler "hello-world/adapter/handler/Hello"
	infra "hello-world/infra/echo"
	"hello-world/infra/profile"
	"log"
	"time"
)

func main() {
	e := echo.New()

	go func() {
		for {
			time.Sleep(3 * time.Second)
			if err := profile.SaveMemoryProfile("mem_profile.prof"); err != nil {
				log.Fatal(err)
			}
		}
	}()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	helloHandler := handler.NewHelloHandler()
	router := infra.NewRouter(e, helloHandler, "GET")
	config := infra.NewEchoConfig("1323", router)
	repository := infra.NewEchoRepository(config)
	repository.Run(e)

}
