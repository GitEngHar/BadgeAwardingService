package notification

import (
	"errors"
	"log"
	"net/mail"
)

type Mail string

var (
	MailInvalidErr = errors.New("invalid mail address")
)

func isValidMail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func NewMail(email string) (Mail, error) {
	isMail := isValidMail(email)
	if isMail {
		return Mail(email), nil
	}
	return "", MailInvalidErr
}

type MailMessage struct {
	To      Mail
	Subject string
	Body    string
}

func NewMailMessage(to, subject, body string) *MailMessage {
	toEmail, err := NewMail(to)
	if err != nil {
		log.Fatalf("Error creating email: %v", err)
	}
	return &MailMessage{
		To:      toEmail,
		Subject: subject,
		Body:    body,
	}
}
