package domain

import (
	"errors"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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

func GetStringAttr(item map[string]types.AttributeValue, key string) (string, bool) {
	if v, ok := item[key]; ok { // 存在チェック
		if av, ok := v.(*types.AttributeValueMemberS); ok { // S型チェック
			return av.Value, true
		}
	}
	return "", false // 無い or S型じゃない
}
