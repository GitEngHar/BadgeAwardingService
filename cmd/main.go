package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
)

type Batch struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IconUri     string `json:"icon_uri"`
}

type ApiGatewayRequestHandler interface {
	Handle(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}

type Config struct {
	apiGatewayHandler ApiGatewayRequestHandler
}

func (c *Config) newConfig(handler ApiGatewayRequestHandler) *Config {
	c.apiGatewayHandler = handler
	return c
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//proxyPath := req.PathParameters["proxy"]
	//var p Batch

}

func main() {

}
