package domain

import "context"

type MailMessage struct {
	To      string
	Subject string
	Body    string
}

type PublishMessage interface {
	SendMessage(ctx context.Context, message string) error
}
