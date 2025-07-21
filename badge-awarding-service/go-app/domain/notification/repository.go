package notification

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type SubscriberRepository interface {
	SubscribeEmail(ctx context.Context, endpoint string) error
	UnSubscribeByEndpoint(ctx context.Context, endpoint string) error
	SendMessageToEmail(ctx context.Context, publisher Publisher) error
	//ListSubscriber(ctx context.Context) ([]Subscriber, error)
}

type MessagePublisher interface {
	PublishMailMessage(ctx context.Context, messageBody string, sqsMessageAttributes map[string]types.MessageAttributeValue) error
	GetMailMessage(ctx context.Context) ([]types.Message, error)
}
