package infra

import (
	"github.com/labstack/echo/v4"
	"log"
)

type Handler interface {
	Hub(ctx echo.Context) error
}
type Router struct {
	server     *echo.Echo
	handler    Handler
	methodType string
}

func NewRouter(server *echo.Echo, handler Handler, methodType string) *Router {
	switch methodType {
	case "GET", "POST", "PUT", "DELETE":
		break
	default:
		log.Fatalf("Unsupported method: %s", methodType)
		return nil
	}
	return &Router{
		server:     server,
		handler:    handler,
		methodType: methodType,
	}
}

func (r *Router) Do() {
	server := r.server
	switch r.methodType {
	case "GET":
		server.GET("/", r.handler.Hub)
	case "POST":
		server.POST("/", r.handler.Hub)
	case "PUT":
		server.PUT("/", r.handler.Hub)
	case "DELETE":
		server.DELETE("/", r.handler.Hub)
	}

}
