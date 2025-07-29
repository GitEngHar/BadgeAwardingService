package push

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"hello-world/domain/management"
	"hello-world/domain/notification"
)

type PublishMessageUseCase struct {
	repo notification.MessagePublisher
}

func NewPublishMessageUseCase(repo notification.MessagePublisher) *PublishMessageUseCase {
	return &PublishMessageUseCase{repo: repo}
}

// TODO: タイトルを設定できるようにする
func (uc PublishMessageUseCase) Do(ctx context.Context, messageBody, userName, address, message string) error {
	imgUrl, err := management.CreatePublicBadgeImgUrl("bucket", "keyname")
	if err != nil {
		return err
	}
	// sqsに送信するinputを生成
	sendMessageWithImgUrl, err := notification.GenerateSendMessageWithImgUrl("aa", *imgUrl)
	sendMessage, err := json.Marshal(sendMessageWithImgUrl)
	if err != nil {
		return err
	}
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
	err = uc.repo.PublishMailMessage(ctx, string(sendMessage), sqsAttributeValues)
	if err != nil {
		return err
	}
	return nil
}
