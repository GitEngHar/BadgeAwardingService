package dynamo

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

type Config struct {
	tableName string
	client    *dynamodb.Client
}
