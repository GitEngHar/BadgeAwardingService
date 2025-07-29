package notification

import (
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type Publisher struct {
	UserName    string `json:"username"`
	Message     string `json:"message"`
	Address     string `json:"address"`
	MessageBody string `json:"message_body"`
}

type UnSubscriptionEndpoint struct {
	Address string `json:"address"`
}

type SendUserMessage struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

func SqsMessageAttributesToEndpoint(record events.SQSMessage) (*UnSubscriptionEndpoint, error) {
	var endpoint UnSubscriptionEndpoint
	if v, ok := record.MessageAttributes["address"]; ok && v.StringValue != nil {
		endpoint.Address = *v.StringValue
	}
	if isEmptyUnsubscriptionEndpoint(endpoint) {
		return nil, errors.New("unsubscription address is empty")
	}
	return &endpoint, nil
}

func SqsMessageAttributesToPublisher(message types.Message) (*Publisher, error) {
	var publisher Publisher
	if message.Body != nil {
		publisher.MessageBody = *message.Body
	}
	if v, ok := message.MessageAttributes["address"]; ok && v.StringValue != nil {
		publisher.Address = *v.StringValue
	}
	if v, ok := message.MessageAttributes["message"]; ok && v.StringValue != nil {
		publisher.Message = *v.StringValue
	}
	if isEmptyPublisher(publisher) {
		return nil, errors.New("missing required field 'address'")
	}
	// userNameは不要なのでコメントアウトにしておく
	//if v, ok := attrs["userName"]; ok && v.StringValue != nil {
	//	p.UserName = *v.StringValue
	//}
	return &publisher, nil
}

func GenerateSendMessageWithImgUrl(title, publicImgUrl string) (*SendUserMessage, error) {
	if title == "" && publicImgUrl == "" {
		return nil, errors.New("message or public image url is empty")
	}
	return &SendUserMessage{
		Title:   title,
		Message: fmt.Sprintf("<!DOCTYPE html><html><body><h2>%s</h2><img src='%s' /></body></html>", title, publicImgUrl),
	}, nil
}

func isEmptyUnsubscriptionEndpoint(endpoint UnSubscriptionEndpoint) bool {
	return endpoint.Address == ""
}

func isEmptyPublisher(publisher Publisher) bool {
	return publisher.Address == "" && publisher.Message == "" && publisher.MessageBody == ""
}
