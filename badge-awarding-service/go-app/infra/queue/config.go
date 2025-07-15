package queue

import "github.com/aws/aws-sdk-go-v2/service/sqs"

type config struct {
	client    *sqs.Client
	queueName string
	queueUrl  string
}
