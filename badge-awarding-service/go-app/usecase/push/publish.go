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
	repo   notification.MessagePublisher
	imgUrl string
}

func NewPublishMessageUseCase(repo notification.MessagePublisher) *PublishMessageUseCase {
	imgUrl, err := management.CreatePublicBadgeImgUrl("bucket", "keyname")
	if err != nil && imgUrl == nil {
		panic(err)
	}
	return &PublishMessageUseCase{
		repo:   repo,
		imgUrl: *imgUrl,
	}
}

var stringType = aws.String("String")

func makeStringAttr(val string) types.MessageAttributeValue {
	return types.MessageAttributeValue{
		DataType:    stringType,
		StringValue: aws.String(val),
	}
}

// TODO: タイトルを設定できるようにする
func (uc PublishMessageUseCase) Do(ctx context.Context, messageBody, userName, address, message string) error {
	// sqsに送信するinputを生成
	sendMessageWithImgUrl, err := notification.GenerateSendMessageWithImgUrl("aa", uc.imgUrl)
	sendMessage, err := json.Marshal(sendMessageWithImgUrl)
	if err != nil {
		return err
	}

	sqsAttributeValues := map[string]types.MessageAttributeValue{
		"userName": makeStringAttr(userName),
		"message":  makeStringAttr(message),
		"address":  makeStringAttr(address),
	}
	// sqsにメッセージを送信する
	err = uc.repo.PublishMailMessage(ctx, string(sendMessage), sqsAttributeValues)
	if err != nil {
		return err
	}
	return nil
}
