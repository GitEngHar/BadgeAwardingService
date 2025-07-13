package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"log"
	"os"
	"time"
)

// dynamoDB
type dynamoDBClient struct {
	tableName string
	client    *dynamodb.Client
}

type user struct {
	email string
}

// sqs
type sqsClient struct {
	client    *sqs.Client
	queueName string
	queueUrl  string
}

func newSqsClient(config aws.Config, tableName string) *sqsClient {
	return &sqsClient{
		client:    sqs.NewFromConfig(config),
		queueName: tableName,
	}
}

func newSqsClientWithQueueURL(client *sqsClient, queueURL string) *sqsClient {
	return &sqsClient{
		client:    client.client,
		queueName: client.queueName,
		queueUrl:  queueURL,
	}
}

func newDynamoDBClient() *dynamoDBClient {
	return &dynamoDBClient{}
}

// sns
type subscribe struct {
	topicArn string
	client   *sns.Client
}

func newSubscribe(snsClient *sns.Client) *subscribe {
	return &subscribe{
		topicArn: os.Getenv("TOPIC_ARN"),
		client:   snsClient,
	}
}

func newUser(email string) *user {
	return &user{email: email}
}

// snsエンドポイントの生成
func createSubscription(sub subscribe, u user) {
	subOut, err := sub.client.Subscribe(context.TODO(), &sns.SubscribeInput{
		TopicArn: aws.String(sub.topicArn),
		Protocol: aws.String("email"),
		Endpoint: aws.String(u.email),
	})
	if err != nil {
		log.Fatalf("unable to subscribe to topic %s, %v", sub.topicArn, err)
	}
	log.Printf("subscribed to topic %s", *subOut.SubscriptionArn)
}

// snsにメッセージを送信する
func sendMessage(sub subscribe, message string) {
	pubOut, err := sub.client.Publish(context.TODO(), &sns.PublishInput{
		TopicArn: aws.String(sub.topicArn),
		Subject:  aws.String("Test Mail Title"),
		Message:  aws.String(message),
	})
	if err != nil {
		log.Fatalf("unable to publish message to topic %s, %v", sub.topicArn, err)
	}
	fmt.Printf("Published message to topic %s\n", *pubOut.MessageId)
}

func deleteSubscription(subscribe subscribe, endpoint string) error {
	var nextToken *string
	ctx := context.Background()
	for {
		resp, err := subscribe.client.ListSubscriptionsByTopic(ctx, &sns.ListSubscriptionsByTopicInput{
			TopicArn:  aws.String(subscribe.topicArn),
			NextToken: nextToken,
		})
		if err != nil {
			log.Fatalf("unable to list subscriptions in topic %s, %v", subscribe.topicArn, err)
		}

		for _, sub := range resp.Subscriptions {
			if aws.ToString(sub.Endpoint) == endpoint {
				// endpointが一致すれば削除
				if _, err := subscribe.client.Unsubscribe(ctx, &sns.UnsubscribeInput{
					SubscriptionArn: sub.SubscriptionArn,
				}); err != nil {
					return fmt.Errorf("unable to unsubscribe from topic %s, %v", subscribe.topicArn, err)
				}
				log.Printf("Unsubscribed from topic %s\n", *sub.TopicArn)
				return nil
			}
			return nil
		}

		if resp.NextToken == nil {
			break
		}
	}
	return fmt.Errorf("subscription not found in topic")
}

// TODO: dynamoDBにユーザー情報を登録する

// TODO: dynamoDBからユーザー情報を取得する

// エンドポイントが存在しない場合は、ユーザーの情報をDynamoDBから削除する

// Mockでエンドポイントが存在しないエラーを返す

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// pollingしてメッセージを取得する
	sQsClient := newSqsClient(cfg, "TestQueue")
	getUrlResp, err := sQsClient.client.GetQueueUrl(context.TODO(), &sqs.GetQueueUrlInput{
		QueueName: aws.String(sQsClient.queueName),
	})
	if err != nil {
		panic(err)
	}
	queueUrl := *getUrlResp.QueueUrl
	fmt.Println("queueUrl:", queueUrl)
	sQsClientWithQueueURL := newSqsClientWithQueueURL(sQsClient, queueUrl)
	polling(sQsClientWithQueueURL, func(queueMessage string) {
		// snsクライアントの生成
		snsClient := sns.NewFromConfig(cfg)
		snsSubscribe := newSubscribe(snsClient)
		fmt.Printf("snsSubscribe: %v\n", snsSubscribe.topicArn)

		email := os.Getenv("EMAIL")
		u := newUser(email)
		createSubscription(*snsSubscribe, *u)
		time.Sleep(20 * time.Second)

		// snsエンドポイントにメールを送信
		sendMessage(*snsSubscribe, queueMessage)
		time.Sleep(20 * time.Second)

		// snsエンドポイントを削除
		err := deleteSubscription(*snsSubscribe, email)
		if err != nil {
			log.Println("unable to delete subscription")
		}
		time.Sleep(20 * time.Second)
		// TODO: dynamoDBからユーザー情報を取得
	})

}

// pollingしてメッセージを取得する
func polling(sqsClientWithQueueURL *sqsClient, snsManagement func(queueMessage string)) {
	for {
		recvResp, err := sqsClientWithQueueURL.client.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(sqsClientWithQueueURL.queueUrl),
			MaxNumberOfMessages: 5,  //最大件数
			WaitTimeSeconds:     10, //待機秒
			VisibilityTimeout:   30, //処理中は他のコンシューマで表示できないようにする
		})

		if err != nil {
			log.Println("Error receiving messages:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		// 取得メッセージがなければ継続取得
		if len(recvResp.Messages) == 0 {
			log.Println("No messages received")
			continue
		}

		// メッセージを処理する
		for _, message := range recvResp.Messages {
			fmt.Printf("Message: %s\n", *message.Body)
			// 受信したメッセージを送信する
			snsManagement(*message.Body)
			// 処理完了したキューを削除する
			_, err := sqsClientWithQueueURL.client.DeleteMessage(context.TODO(), &sqs.DeleteMessageInput{
				QueueUrl:      aws.String(sqsClientWithQueueURL.queueUrl),
				ReceiptHandle: message.ReceiptHandle,
			})

			if err != nil {
				log.Println("Error deleting message:", err)
			} else {
				log.Println("Deleted message:", *message.MessageId)
			}
		}
	}

}
