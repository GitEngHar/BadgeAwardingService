package infra

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

type Repository struct {
	config *EchoConfig
}

func NewEchoRepository(config *EchoConfig) *Repository {
	return &Repository{config: config}
}

func (er Repository) Run(e *echo.Echo) {
	config := er.config
	server := e
	config.router.Do()
	fmt.Println(config)
	server.Logger.Fatal(server.Start(config.port))
}
