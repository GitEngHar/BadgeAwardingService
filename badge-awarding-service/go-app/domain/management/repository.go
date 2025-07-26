package management

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type BadgeUserRepository interface {
	Upsert(ctx context.Context, item map[string]types.AttributeValue) error
	Get(ctx context.Context, item map[string]types.AttributeValue) (map[string]types.AttributeValue, error)
	Del(ctx context.Context, filter map[string]types.AttributeValue) error
	CreateTable(ctx context.Context) error
}

type BadgeImgRepository interface {
	DownloadBadge(ctx context.Context) (*Badge, error)
}
