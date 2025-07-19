package notification

import (
	"context"
	"time"
)

type MessagePublisher interface {
	Publish(ctx context.Context, message MailMessage) error
}

type Subscriber struct {
	ID         string
	Endpoint   string
	Protocol   string
	Subscribed bool
	CreatedAt  time.Time
}

func NewSubscriber(id, endpoint, protocol string, subscribed bool, createdAt time.Time) *Subscriber {
	return &Subscriber{
		ID:         id,
		Endpoint:   endpoint,
		Protocol:   protocol,
		Subscribed: subscribed,
		CreatedAt:  createdAt,
	}
}
