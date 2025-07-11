package infra

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"strconv"
)

type EchoConfig struct {
	port   string
	server *echo.Echo
	router *Router
}

func NewEchoConfig(port string, router *Router) *EchoConfig {
	if _, err := strconv.Atoi(port); err != nil {
		log.Fatalf("port must be an integer")
		return nil
	}

	return &EchoConfig{
		port:   fmt.Sprintf(":%s", port),
		router: router,
	}
}
