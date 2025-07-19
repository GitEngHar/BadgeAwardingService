package dynamo

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"hello-world/domain/management"
)

type BadgeRepository struct {
	config Config
}

const (
	badgeTableName = "badges"
)

func NewBadgeRepository(ctx context.Context) management.BadgeRepository {
	dbConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}
	client := dynamodb.NewFromConfig(dbConfig)
	conf := Config{
		tableName: userTableName,
		client:    client,
	}
	return BadgeRepository{
		config: conf,
	}
}

func (b BadgeRepository) Upsert(ctx context.Context, userBadge *management.Badge) error {
	return nil
}

func (b BadgeRepository) Create(ctx context.Context, userBadge management.UserBadge) error {
	return nil
}

func (b BadgeRepository) GetByID(ctx context.Context, id string) (management.Badge, error) {
	return management.Badge{}, nil
}

func (b BadgeRepository) Delete(ctx context.Context, id string) error {
	return nil
}
