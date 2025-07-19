package notification

import (
	"context"
)

type SubscriberService interface {
	Subscribe(ctx context.Context, subscription Subscriber) error
	UnSubscribe(ctx context.Context, subscription Subscriber) error
	ListSubscriber(ctx context.Context) ([]Subscriber, error)
}

type SubscriberRepository interface {
	Save(ctx context.Context, s *Subscriber) error
	Delete(ctx context.Context, id string) error
	FindAll(ctx context.Context) ([]*Subscriber, error)
	FindByID(ctx context.Context, id string) (*Subscriber, error)
}
