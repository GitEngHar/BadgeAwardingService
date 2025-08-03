package infra

import (
	"github.com/labstack/echo/v4"
)

type GetHandler interface {
	Hub(ctx echo.Context) error
}

type PostHandler interface {
	Hub(ctx echo.Context) error
}

type PutHandler interface {
	Hub(ctx echo.Context) error
}

type DeleteHandler interface {
	Hub(ctx echo.Context) error
}

type Router struct {
	server        *echo.Echo
	getHandler    GetHandler
	postHandler   PostHandler
	putHandler    PutHandler
	deleteHandler DeleteHandler
	methodType    string
}

func NewRouter(server *echo.Echo, getHandler GetHandler, postHandler PostHandler, putHandler PutHandler, deleteHandler DeleteHandler) *Router {
	return &Router{
		server:        server,
		getHandler:    getHandler,
		postHandler:   postHandler,
		putHandler:    putHandler,
		deleteHandler: deleteHandler,
	}
}

func (r *Router) Do() {
	server := r.server
	if r.getHandler != nil {
		server.GET("/", r.getHandler.Hub)
	}
	if r.deleteHandler != nil {
		server.DELETE("/", r.deleteHandler.Hub)
	}
	if r.postHandler != nil {
		server.POST("/", r.postHandler.Hub)
	}
	if r.putHandler != nil {
		server.PUT("/", r.putHandler.Hub)
	}
}
