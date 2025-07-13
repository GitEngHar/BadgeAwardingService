package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"log"
	"os"
)

type subscribe struct {
	topicArn string
	client   *sns.Client
}

type user struct {
	email string
}

func newClient(topicArn string, client *sns.Client) *subscribe {
	return &subscribe{
		topicArn: topicArn,
		client:   client,
	}
}

func newUser(email string) *user {
	return &user{
		email: email,
	}
}

func main() {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// subscriberを生成する
	snsClient := sns.NewFromConfig(cfg)
	topicArn := os.Getenv("SNS_TOPIC_ARN")
	//endpoint := os.Getenv("SNS_ENDPOINT")
	//u := newUser(endpoint)
	client := newClient(topicArn, snsClient)
	//createSubscription(*client, *u)

	// subscriberにメッセージを送信する
	message := "Hello!! This is Test Mail"
	sendMessage(*client, message)

}

func createSubscription(sub subscribe, u user) {
	subOut, err := sub.client.Subscribe(context.TODO(), &sns.SubscribeInput{
		TopicArn: aws.String(sub.topicArn),
		Protocol: aws.String("email"),
		Endpoint: aws.String(u.email),
	})
	if err != nil {
		log.Fatalf("unable to subscribe to topic %s, %v", sub.topicArn, err)
	}
	fmt.Printf("Subscribed to topic %s\n", *subOut.SubscriptionArn)
}

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
