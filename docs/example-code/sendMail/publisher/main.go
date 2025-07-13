package main

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

// sqs
type sqsClient struct {
	client    *sqs.Client
	queueName string
}
type sqsClientWithQueueURL struct {
	client    *sqs.Client
	queueName string
	queueUrl  string
}

// dynamoDB
type dynamoDBClient struct {
	tableName string
	client    *dynamodb.Client
}

// sns
type snsClient struct {
	client *sns.Client
}

type user struct {
	email string
}

func newSqsClient() *sqsClient {
	return &sqsClient{}
}

func newSqsClientWithQueueURL() *sqsClientWithQueueURL {
	return &sqsClientWithQueueURL{}
}

func newDynamoDBClient() *dynamoDBClient {
	return &dynamoDBClient{}
}

func newSnsClient() *snsClient {
	return &snsClient{}
}

func newUser(email string) *user {
	return &user{email: email}
}

// sqsにメッセージを送信する

// dynamoDBにユーザー情報を登録する

// dynamoDBからユーザー情報を取得する

// pollingしてメッセージを取得する

// userのSNSエンドポイントを生成し、メールを送信する

// エンドポイントが存在しない場合は、ユーザーの情報をDynamoDBから削除する

// Mockでエンドポイントが存在しないエラーを返す

func main() {
	// dynamoDBにユーザー情報を登録
	// TODO: 本物のdynamoDBから取得するように変更する

	// sqsにメッセージを送信

	// dynamoDBにユーザー情報を
}
