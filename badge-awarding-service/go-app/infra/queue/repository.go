package queue

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"hello-world/domain/notification"
	"os"
)

type Publisher struct {
	config Config
}

var (
	notFoundMessage = errors.New("publisher(sqs queue) not exists")
)

func NewConfig(ctx context.Context) *Config {
	// queueのURLを取得
	queueURL := os.Getenv("QUEUE_URL")
	if queueURL == "" {
		panic("environment variable QUEUE_URL is not set")
	}

	// AWS設定をロード
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {

		panic("unable to load SDK config, " + err.Error())
	}

	// SQSクライアントの作成
	client := sqs.NewFromConfig(cfg)

	return &Config{
		client:   client,
		queueUrl: queueURL,
	}
}

func NewPublisher(config Config) notification.MessagePublisher {
	return Publisher{config: config}
}

func (q Publisher) PublishMailMessage(ctx context.Context, messageBody string, sqsMessageAttributes map[string]types.MessageAttributeValue) error {
	input := &sqs.SendMessageInput{
		QueueUrl:          &q.config.queueUrl,
		MessageBody:       aws.String(messageBody),
		MessageAttributes: sqsMessageAttributes,
	}
	_, err := q.config.client.SendMessage(ctx, input)
	if err != nil {
		return err
	}
	fmt.Println("Successfully published message")
	return nil
}

func (q Publisher) GetMailMessage(ctx context.Context) ([]types.Message, error) {
	revResp, err := q.config.client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:              aws.String(q.config.queueUrl),
		MaxNumberOfMessages:   10, //最大件数
		WaitTimeSeconds:       1,
		VisibilityTimeout:     5,
		MessageAttributeNames: []string{"All"},
	})

	if err != nil {
		return nil, err
	}

	if len(revResp.Messages) == 0 {
		return nil, notFoundMessage
	}
	return revResp.Messages, nil
}
