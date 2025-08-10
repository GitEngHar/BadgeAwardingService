package notification

import (
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"hello-world/domain"
)

type UnSubscriptionEndpoint struct {
	Address string `json:"address"`
}

type AwardComment struct {
	title         string
	praiseComment string
	reason        string
	effect        string
}

type Publisher struct {
	Title   string `json:"title"`
	Message string `json:"message"`
	Address string `json:"address"`
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

func NewAwardComment(title, praiseComment, reason, effect string) (*AwardComment, error) {
	if title == "" && praiseComment == "" && reason == "" {
		return nil, errors.New("message or public image url is empty")
	}
	return &AwardComment{
		title:         title,
		praiseComment: praiseComment,
		reason:        reason,
		effect:        effect,
	}, nil
}

func NewPublisher(awardComment *AwardComment, address string) (*Publisher, error) {
	if awardComment == nil {
		return nil, errors.New("award comment is empty")
	}
	if _, err := domain.NewMail(address); err != nil {
		return nil, err
	}
	publishMessage := generatePublishMessage(awardComment)
	return &Publisher{
		Title:   awardComment.title,
		Message: publishMessage,
		Address: address,
	}, nil
}

func generatePublishMessage(awardComment *AwardComment) string {
	awardMessage := fmt.Sprintf("<ノルマを達成！>, 「%s！」\n%s\n%s\n偉大な業績に敬意を表します。\n\n【報酬】\n%s", awardComment.title, awardComment.reason, awardComment.praiseComment, awardComment.effect)
	return fmt.Sprintf("%s", awardMessage)
}

func isEmptyUnsubscriptionEndpoint(endpoint UnSubscriptionEndpoint) bool {
	return endpoint.Address == ""
}

func SubscribedMessageToPublish(message types.Message) (*Publisher, error) {
	var publisher Publisher
	if v, ok := message.MessageAttributes["title"]; ok && v.StringValue != nil {
		publisher.Title = *v.StringValue
	}
	if v, ok := message.MessageAttributes["address"]; ok && v.StringValue != nil {
		publisher.Address = *v.StringValue
	}
	if v, ok := message.MessageAttributes["message"]; ok && v.StringValue != nil {
		publisher.Message = *v.StringValue
	}
	return &publisher, nil
}
