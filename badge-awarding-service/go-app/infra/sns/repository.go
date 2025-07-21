package sns

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"hello-world/domain/notification"
	"os"
	"strings"
)

type Subscription struct {
	config Config
}

func NewSubscription(config Config) notification.SubscriberRepository {
	return &Subscription{
		config: config,
	}
}

func NewConfig(ctx context.Context) Config {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic(err)
	}

	topicArn := os.Getenv("SNS_TOPIC_ARN")
	if topicArn == "" {
		panic("environment variable SNS_TOPIC_ARN is not set")
	}
	return Config{
		topicArn: topicArn,
		client:   sns.NewFromConfig(cfg),
	}
}

func (s Subscription) SubscribeEmail(ctx context.Context, endpoint string) error {
	_, err := s.config.client.Subscribe(ctx, &sns.SubscribeInput{
		TopicArn: aws.String(s.config.topicArn),
		Protocol: aws.String("email"),
		Endpoint: aws.String(endpoint),
	})
	return err
}

func (s Subscription) UnSubscribeEmail(ctx context.Context, subscription notification.Subscriber) error {
	return nil
}

func (s Subscription) SendMessageToEmail(ctx context.Context, publisher notification.Publisher) error {

	isExistsEndpoint, err := s.isEmailSubscribed(ctx, publisher.Address)

	if err != nil {
		return err
	}

	// 送信前に確認メールが走るので、作られていない場合はDeadLetterQueueに送信して確認メールを送る
	if !isExistsEndpoint {
		if err := s.SubscribeEmail(ctx, publisher.Address); err != nil {
			return err
		}
		// DeadLetterQueueに送り返す
		return errors.New("email subscribed but message not exists")
	}

	_, err = s.config.client.Publish(ctx, &sns.PublishInput{
		TopicArn: aws.String(s.config.topicArn),
		Message:  aws.String(publisher.Message),
		Subject:  aws.String(publisher.MessageBody),
	})

	if err != nil {
		return err
	}
	return nil
}

func (s Subscription) isEmailSubscribed(ctx context.Context, endpoint string) (bool, error) {
	paginator := sns.NewListSubscriptionsByTopicPaginator(s.config.client, &sns.ListSubscriptionsByTopicInput{
		TopicArn: aws.String(s.config.topicArn),
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return false, err
		}
		for _, sub := range page.Subscriptions {
			if strings.EqualFold(aws.ToString(sub.Endpoint), endpoint) {
				return true, nil
			}
		}
	}
	return false, nil
}
