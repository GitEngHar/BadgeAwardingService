package notification

import (
	"context"
	"time"
)

type MailMessage struct {
	To      string
	Subject string
	Body    string
}

func NewMailMessage(to, subject, body string) *MailMessage {
	return &MailMessage{
		To:      to,
		Subject: subject,
		Body:    body,
	}
}

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
