package push

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"hello-world/domain/notification"
)

type SubscriptionUseCase struct {
	subRepo notification.SubscriberRepository
	pubRepo notification.MessagePublisher
}

func NewSubscriptionUseCase(subRepo notification.SubscriberRepository, pubRepo notification.MessagePublisher) *SubscriptionUseCase {
	return &SubscriptionUseCase{subRepo: subRepo, pubRepo: pubRepo}
}

func (uc SubscriptionUseCase) PushEmail(ctx context.Context) ([]map[string]types.MessageAttributeValue, error) {
	// queueからメッセージをポーリングする
	messages, err := uc.pubRepo.GetMailMessage(ctx)
	if err != nil {
		return nil, err
	}
	subscribedAwards := make([]map[string]types.MessageAttributeValue, 0)
	// 全てのメッセージを送信する
	for _, message := range messages {
		subscribedPublish, err := notification.SubscribedMessageToPublish(message)
		if err != nil {
			fmt.Println("publisher Unmarshal Error :", err)
			continue
		}

		// ユーザーのメールアドレスにメッセージを送信する
		err = uc.subRepo.SendMessageToEmail(ctx, *subscribedPublish)
		if err != nil {
			fmt.Println("Publisher Error :", err)
			continue
		}

		subscribedAwards = append(subscribedAwards, message.MessageAttributes)
	}
	return subscribedAwards, nil
}

func (uc SubscriptionUseCase) Do(ctx context.Context) ([]map[string]types.MessageAttributeValue, error) {
	// queueからメッセージをポーリングする
	messages, err := uc.pubRepo.GetMailMessage(ctx)
	if err != nil {
		return nil, err
	}
	subscribedAwards := make([]map[string]types.MessageAttributeValue, 0)
	// 全てのメッセージを送信する
	for _, message := range messages {
		subscribedPublish, err := notification.SubscribedMessageToPublish(message)
		if err != nil {
			fmt.Println("publisher Unmarshal Error :", err)
			continue
		}

		// ユーザーのメールアドレスにメッセージを送信する
		err = uc.subRepo.SendMessageToEmail(ctx, *subscribedPublish)
		if err != nil {
			fmt.Println("Publisher Error :", err)
			continue
		}

		subscribedAwards = append(subscribedAwards, message.MessageAttributes)
	}
	return subscribedAwards, nil
}
