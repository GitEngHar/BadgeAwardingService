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
	var execPublisherCount = 0
	// queueからメッセージをポーリングする
	messages, err := uc.pubRepo.GetMailMessage(ctx)
	if err != nil {
		return err
	}

	// 全てのメッセージを送信する
	for _, message := range messages {
		// messageのドメインを作成する
		publisher, err := notification.SqsMessageAttributesToPublisher(message)

		// 処理件数をカウント
		if err != nil {
			fmt.Println("publisher Unmarshal Error :", err)
			continue
		}

		// ユーザーのメールアドレスにメッセージを送信する
		err = uc.subRepo.SendMessageToEmail(ctx, *publisher)
		if err != nil {
			fmt.Println("Publisher Error :", err)
			continue
		}
		execPublisherCount++
	}

	fmt.Printf("Plan %d , End %d, Perse: %d%", len(messages), execPublisherCount, execPublisherCount/len(messages))

	return nil
}
