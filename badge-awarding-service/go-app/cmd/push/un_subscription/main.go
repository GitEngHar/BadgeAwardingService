package main

import (
	"context"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"hello-world/adapter/handler/Push"
	"hello-world/domain/notification"
	"hello-world/infra/sns"
	"hello-world/usecase/push"
)

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	// sqsからendpointを生成
	for _, record := range sqsEvent.Records {
		endpoint, err := notification.SqsMessageAttributesToEndpoint(record)
		if err != nil {
			return err
		}
		repoConfig := sns.NewConfig(ctx)
		repo := sns.NewSubscription(repoConfig)
		uc := push.NewUnSubscriptionUseCase(repo)
		UnsubscriptionHandler := Push.NewUnSubscriptionHandler(*uc)
		return UnsubscriptionHandler.Do(ctx, *endpoint)
	}

	return errors.New("no sqs event")
}

func main() {
	lambda.Start(handler)
}
