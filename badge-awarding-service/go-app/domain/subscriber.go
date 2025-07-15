package domain

import (
	"context"
	"time"
)

type Subscriber struct {
	ID         string
	Endpoint   string
	Protocol   string
	Subscribed bool
	CreatedAt  time.Time
}

type SubscriptionService interface {
	Subscribe(ctx context.Context, subscription Subscriber) error
	UnSubscribe(ctx context.Context, subscription Subscriber) error
	ListSubscriber(ctx context.Context) ([]Subscriber, error)
}
