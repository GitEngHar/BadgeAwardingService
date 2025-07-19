package dynamo

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"hello-world/domain/management"
	"time"
)

const (
	tableName = "badge-service"
)

type DBRepository struct {
	config Config
}

func NewConnectionDynamoDBForAWS(ctx context.Context) *Config {
	dynamodbDefaultConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}
	dynamodbClient := dynamodb.NewFromConfig(dynamodbDefaultConfig)
	return &Config{
		TableName: tableName,
		Client:    dynamodbClient,
	}
}

func NewConnectionDynamoDBForLocal() *Config {
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

	return &Config{
		TableName: tableName,
		Client:    dynamodbClient,
	}
}

func NewUserRepository(config *Config) management.UserBadgeRepository {
	return &DBRepository{
		config: *config,
	}
}

// Upsert 値が存在する場合は更新、しない場合は追加
func (d DBRepository) Upsert(ctx context.Context, item map[string]types.AttributeValue) error {
	_, err := d.config.Client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(d.config.TableName),
		Item:      item,
	})
	return err
}

// Get 指定したKeyのItemを取得する
func (d DBRepository) Get(ctx context.Context, filter map[string]types.AttributeValue) (map[string]types.AttributeValue, error) {
	output, err := d.config.Client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(d.config.TableName),
		Key:       filter,
	})
	if err != nil {
		return nil, err
	}
	if output.Item == nil {
		return nil, errors.New("item does not exist")
	}
	return output.Item, nil
}

// Delete 指定したKeyのItemを取得する
func (d DBRepository) Del(ctx context.Context, filter map[string]types.AttributeValue) error {
	_, err := d.config.Client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(d.config.TableName),
		Key:       filter,
	})
	return err
}

// CreateTable テーブルを作成する
func (d DBRepository) CreateTable(ctx context.Context) error {
	_, err := d.config.Client.CreateTable(ctx, &dynamodb.CreateTableInput{
		TableName: aws.String(d.config.TableName),
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
		var re *types.ResourceInUseException
		// テーブルがすでに存在する場合はエラーではないので通常の出力のみ
		if errors.As(err, &re) {
			fmt.Printf("already exists table %s: %v\n", d.config.TableName, re)
		} else {
			panic(err)
		}
	} else {
		// tableを作成する
		fmt.Printf("table %s created\n", d.config.TableName)
		waiter := dynamodb.NewTableExistsWaiter(d.config.Client)
		if err := waiter.Wait(ctx, &dynamodb.DescribeTableInput{
			TableName: aws.String(d.config.TableName),
		}, 5*time.Minute); err != nil {
			panic(err)
		}
		fmt.Printf("table %s exists\n", d.config.TableName)
	}
	return nil
}
