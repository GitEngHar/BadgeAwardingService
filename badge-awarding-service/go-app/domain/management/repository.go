package management

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// dynamoDBはtableを統一するため注入するRepositoryは同一

type Repository interface {
	Upsert(ctx context.Context, item map[string]types.AttributeValue) error
	GetByPK(ctx context.Context, pk string) (map[string]types.AttributeValue, error)
	Del(ctx context.Context, filter map[string]types.AttributeValue) error
	CreateTable(ctx context.Context) error
}
