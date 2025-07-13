package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"log"
	"time"
)

type sqsClient struct {
	client    *sqs.Client
	queueName string
	queueURL  string
}

func newClient(queueName string, config aws.Config) *sqsClient {
	return &sqsClient{
		client:    sqs.NewFromConfig(config),
		queueName: queueName,
	}
}

func newClientWithQueueURL(client *sqsClient, queueURL string) *sqsClient {
	return &sqsClient{
		client:    client.client,
		queueName: client.queueName,
		queueURL:  queueURL,
	}
}

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}
	sqsClient := newClient("TestQueue", cfg)
	getUrlResp, err := sqsClient.client.GetQueueUrl(context.TODO(), &sqs.GetQueueUrlInput{
		QueueName: aws.String(sqsClient.queueName),
	})
	if err != nil {
		panic(err)
	}
	queueUrl := *getUrlResp.QueueUrl
	fmt.Println("queueUrl:", queueUrl)
	sqsClientWithQueueURL := newClientWithQueueURL(sqsClient, queueUrl)
	polling(sqsClientWithQueueURL)
}

func polling(sqsClientWithQueueURL *sqsClient) {
	for {
		recvResp, err := sqsClientWithQueueURL.client.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(sqsClientWithQueueURL.queueURL),
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
			// messageに応じて本来は処理する

			// 処理完了したキューを削除する
			_, err := sqsClientWithQueueURL.client.DeleteMessage(context.TODO(), &sqs.DeleteMessageInput{
				QueueUrl:      aws.String(sqsClientWithQueueURL.queueURL),
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
