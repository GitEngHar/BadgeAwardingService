package push

import (
	"context"
	"fmt"
	"hello-world/domain/notification"
)

type SubscriptionUseCase struct {
	subRepo notification.SubscriberRepository
	pubRepo notification.MessagePublisher
}

func NewSubscriptionUseCase(subRepo notification.SubscriberRepository, pubRepo notification.MessagePublisher) *SubscriptionUseCase {
	return &SubscriptionUseCase{subRepo: subRepo, pubRepo: pubRepo}
}

func (uc SubscriptionUseCase) Do(ctx context.Context) error {
	//var execPublisherCount = 0
	// queueからメッセージをポーリングする
	messages, err := uc.pubRepo.GetMailMessage(ctx)
	if err != nil {
		return err
	}

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
		//execPublisherCount++
	}
	return nil
}
