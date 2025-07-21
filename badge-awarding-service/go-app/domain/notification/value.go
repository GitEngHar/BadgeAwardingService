package notification

import (
	"errors"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type Publisher struct {
	UserName    string `json:"username"`
	Message     string `json:"message"`
	Address     string `json:"address"`
	MessageBody string `json:"message_body"`
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
	if isEmpty(publisher) {
		return nil, errors.New("missing required field 'address'")
	}
	// userNameは不要なのでコメントアウトにしておく
	//if v, ok := attrs["userName"]; ok && v.StringValue != nil {
	//	p.UserName = *v.StringValue
	//}
	return &publisher, nil
}

func isEmpty(publisher Publisher) bool {
	return publisher.Address == "" && publisher.Message == "" && publisher.MessageBody == ""
}
