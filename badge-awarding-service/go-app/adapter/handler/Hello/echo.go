package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Hello struct{}

func NewHelloHandler() *Hello {
	return &Hello{}
}

func (h *Hello) Do() (map[string]string, error) {
	return map[string]string{"hello": "world"}, nil
}

func (h *Hello) Hub(c echo.Context) error {
	returnStr, _ := h.Do()
	return c.JSON(http.StatusOK, returnStr)
}
