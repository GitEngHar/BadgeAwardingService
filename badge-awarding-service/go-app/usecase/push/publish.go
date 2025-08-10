package push

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"hello-world/domain/management"
	"hello-world/domain/notification"
	"strconv"
	"time"
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
func (uc PublishMessageUseCase) Do(ctx context.Context, publishers []notification.Publisher) error {
	//TODO: publishersの依存性を切り離してインフラに効率的なアップロード処理を移す

	// sqsは最大10件を同時送信できるので10にしておく
	var sendEntries = make([]types.SendMessageBatchRequestEntry, 0, 10)
	var sqsAttributeValues = map[string]types.MessageAttributeValue{}
	var isSendSplitTen bool
	var isSendSarPlus bool
	today := time.Now().Format("2006-01-02")
	splitTen := len(publishers) / 10
	sendSarPlus := len(publishers) % 10
	// TODO: IDは日付と回数にしているが、後ほどちゃんと設計する
	// sqsに送信するinputを生成
	for i, publisher := range publishers {
		sqsAttributeValues = map[string]types.MessageAttributeValue{
			"message": makeStringAttr(publisher.Message),
			"sub":     makeStringAttr(publisher.Title),
			"address": makeStringAttr(publisher.Address),
		}
		sendEntries = append(sendEntries, types.SendMessageBatchRequestEntry{
			MessageAttributes: sqsAttributeValues,
			MessageBody:       &publisher.Message,
			Id:                aws.String(fmt.Sprintf("%s-%s", today, strconv.Itoa(i))),
		})
		isSendSplitTen = i+1%10 == 0
		isSendSarPlus = (i/10 == splitTen) && (i+1 == sendSarPlus)
		// sqsは最大10件を一斉送信する
		if isSendSplitTen {
			err := uc.repo.PublishMailMessage(ctx, sendEntries)
			sendEntries = sendEntries[:0] // メモリを再利用したいのでリセットする
			if err != nil {
				return err
			}
			fmt.Printf("Published message of number is : %d\n", i)
		} else if isSendSarPlus {
			err := uc.repo.PublishMailMessage(ctx, sendEntries)
			sendEntries = sendEntries[:0] // メモリを再利用したいのでリセットする
			if err != nil {
				return err
			}
			fmt.Printf("Published message of number is : %d\n", i)
		}
	}
	return nil
}
