// main.go
package main

import (
	"log"

	"hello-world/adapter/dipendency_injection"

	"github.com/aws/aws-lambda-go/lambda"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
)

var adapter *chiadapter.ChiLambda

func init() {
	r, err := dipendency_injection.InitializeRouter()
	if err != nil {
		log.Fatal(err)
	}
	adapter = chiadapter.New(r)
}

func main() {
	lambda.Start(adapter.ProxyWithContext)
}
