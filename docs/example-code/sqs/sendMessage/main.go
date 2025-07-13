package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
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
	sendMessage(*sqsClientWithQueueURL, "Hello!! This is Test Mail")
}

func sendMessage(sqsClientWithQueueURL sqsClient, message string) {
	send, err := sqsClientWithQueueURL.client.SendMessage(context.TODO(), &sqs.SendMessageInput{
		QueueUrl:    &sqsClientWithQueueURL.queueURL,
		MessageBody: aws.String(message),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Send Message: %s\n", *send.MessageId)
}
