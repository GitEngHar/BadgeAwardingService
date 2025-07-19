package dynamo

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"hello-world/domain/management"
)

const (
	userTableName = "users"
)

type UserRepositorySt struct {
	config Config
}

func NewUserRepository(ctx context.Context) management.UserRepository {
	dbConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}
	client := dynamodb.NewFromConfig(dbConfig)
	conf := Config{
		tableName: userTableName,
		client:    client,
	}
	return UserRepositorySt{
		config: conf,
	}
}

func (u UserRepositorySt) Upsert(ctx context.Context, user management.User) error {
	return nil
}

func (u UserRepositorySt) Create(ctx context.Context, user management.User) error {
	return nil
}

func (u UserRepositorySt) GetByID(ctx context.Context, id string) (management.User, error) {
	// mockで作る
	return management.User{}, nil
}

func (u UserRepositorySt) Delete(ctx context.Context, id string) error {
	return nil
}

// CreateTable テーブルを作成する
func (u UserRepositorySt) CreateTable(ctx context.Context) error {
	return nil
}
