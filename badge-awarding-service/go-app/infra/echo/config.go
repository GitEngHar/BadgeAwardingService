package echo

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"strconv"
)

type Config struct {
	port   string
	server *echo.Echo
	router *Router
}

func NewConfig(port string, router *Router) *Config {
	if _, err := strconv.Atoi(port); err != nil {
		log.Fatalf("port must be an integer")
		return nil
	}

	return &Config{
		port:   fmt.Sprintf(":%s", port),
		router: router,
	}
}
