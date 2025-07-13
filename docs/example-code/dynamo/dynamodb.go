package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"time"
)

type dynamoDB struct {
	tableName string
	client    *dynamodb.Client
}

func connectionDynamoDB(ctx context.Context) *dynamoDB {
	dynamodbDefaultConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}
	dynamodbClient := dynamodb.NewFromConfig(dynamodbDefaultConfig)
	return &dynamoDB{
		tableName: "test",
		client:    dynamodbClient,
	}
}

func connectionDynamoDBForLocal(ctx context.Context) *dynamoDB {
	region := "ap-northeast-1"
	endpoint := "http://localhost:8000"
	dynamodbDefaultConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	// クライアント生成
	dynamodbClient := dynamodb.NewFromConfig(dynamodbDefaultConfig, func(o *dynamodb.Options) {
		o.BaseEndpoint = aws.String(endpoint)
	})

	return &dynamoDB{
		tableName: "Users",
		client:    dynamodbClient,
	}
}

func getTableList(ctx context.Context, client *dynamodb.Client) []string {
	out, err := client.ListTables(ctx, &dynamodb.ListTablesInput{})
	if err != nil {
		panic("unable to list DynamoDB tables, " + err.Error())
	}
	return out.TableNames
}

func createTable(ctx context.Context, dynamoDB *dynamoDB) {
	_, err := dynamoDB.client.CreateTable(ctx, &dynamodb.CreateTableInput{
		TableName: aws.String(dynamoDB.tableName),
		KeySchema: []types.KeySchemaElement{
			{AttributeName: aws.String("id"), KeyType: types.KeyTypeHash},
		},
		AttributeDefinitions: []types.AttributeDefinition{
			{AttributeName: aws.String("id"), AttributeType: types.ScalarAttributeTypeS},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
	})
	if err != nil {
		//すでに作成済みの場合はスキップ
		var re *types.ResourceInUseException
		if ok := errors.As(err, &re); ok {
			fmt.Println("already table exists")
		} else {
			panic(err)
		}
	} else {
		fmt.Printf("Creating table %q...\n", dynamoDB.tableName)
		waiter := dynamodb.NewTableExistsWaiter(dynamoDB.client)
		if err := waiter.Wait(ctx, &dynamodb.DescribeTableInput{
			TableName: aws.String(dynamoDB.tableName),
		}, 5*time.Minute); err != nil {
			panic(err)
		}
		fmt.Println("table created")
	}

	fmt.Println("Tables:", getTableList(ctx, dynamoDB.client))

}

func main() {
	ctx := context.Background()
	db := connectionDynamoDBForLocal(ctx)
	createTable(ctx, db)
}
