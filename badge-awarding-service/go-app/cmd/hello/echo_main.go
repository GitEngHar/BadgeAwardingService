package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	handler "hello-world/adapter/handler/Hello"
	infra "hello-world/infra/echo"
	"hello-world/infra/newRelic"
)

func main() {
	e := echo.New()
	nrApp := newRelic.InitializeNewRelic()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(nrecho.Middleware(&nrApp.App))
	//e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	//	return func(c echo.Context) error {
	//		req := c.Request()
	//		txn := newRelicApp.App.StartTransaction(req.URL.Path)
	//		defer txn.End()
	//		txn.SetWebRequestHTTP(req)
	//		c.SetRequest(req.WithContext(newrelic.NewContext(req.Context(), txn)))
	//		// related response
	//		c.Response().Writer = txn.SetWebResponse(c.Response())
	//		return next(c)
	//	}
	//})
	helloHandler := handler.NewHelloHandler()
	router := infra.NewRouter(e, helloHandler, "GET")
	config := infra.NewEchoConfig("1323", router)
	repository := infra.NewEchoRepository(config)
	repository.Run(e)

}
