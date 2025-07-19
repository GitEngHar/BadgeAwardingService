package dynamo

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

type Config struct {
	TableName string
	Client    *dynamodb.Client
}
