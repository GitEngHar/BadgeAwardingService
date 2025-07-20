package push

import (
	"context"
	"hello-world/domain/notification"
)

type SubscriptionUseCase struct {
	repo notification.SubscriberRepository
}

func NewSubscriptionUseCase(repo notification.SubscriberRepository) *SubscriptionUseCase {
	return &SubscriptionUseCase{repo: repo}
}

func (uc SubscriptionUseCase) Do(ctx context.Context, endpoint string) error {
	// ユーザーのメールアドレスにメッセージを送信する
	err := uc.repo.SendMessageToEmail(ctx, endpoint)
	if err != nil {
		return err
	}
	return nil
}
