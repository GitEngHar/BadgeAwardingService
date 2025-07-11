package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Hello struct{}

func NewHelloHandler() *Hello {
	return &Hello{}
}

func (h *Hello) Do(c echo.Context) error {
	return c.JSON(http.StatusOK, []string{"ok"})
}
