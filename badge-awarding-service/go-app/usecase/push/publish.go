package push

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"hello-world/domain/notification"
)

type PublishMessageUseCase struct {
	repo notification.MessagePublisher
}

func NewPublishMessageUseCase(repo notification.MessagePublisher) *PublishMessageUseCase {
	return &PublishMessageUseCase{repo: repo}
}

func (uc PublishMessageUseCase) Do(ctx context.Context, messageBody, userName, address, message string) error {
	// sqsに送信するinputを生成
	sqsAttributeValues := map[string]types.MessageAttributeValue{
		"userName": {
			DataType:    aws.String("String"),
			StringValue: aws.String(userName),
		},
		"message": {
			DataType:    aws.String("String"),
			StringValue: aws.String(message),
		},
		"address": {
			DataType:    aws.String("String"),
			StringValue: aws.String(address),
		},
	}
	// sqsにメッセージを送信する
	err := uc.repo.PublishMailMessage(ctx, messageBody, sqsAttributeValues)
	if err != nil {
		return err
	}
	return nil
}
