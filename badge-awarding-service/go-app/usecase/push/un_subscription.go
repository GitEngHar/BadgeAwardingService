package push

import (
	"context"
	"hello-world/domain/notification"
)

type UnSubscriptionUseCase struct {
	repo notification.SubscriberRepository
}

func NewUnSubscriptionUseCase(repo notification.SubscriberRepository) *UnSubscriptionUseCase {
	return &UnSubscriptionUseCase{repo: repo}
}

func (u *UnSubscriptionUseCase) Do(ctx context.Context, endpoint string) error {
	return u.repo.UnSubscribeByEndpoint(ctx, endpoint)
}
