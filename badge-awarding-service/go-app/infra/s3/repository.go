package s3

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"hello-world/domain/management"
)

type BadgeImageRepository struct {
	config Config
}

func NewConfig(ctx context.Context, bucketName, bucketKey string) *Config {
	// セッション作成
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	})
	if err != nil {
		panic(err)
	}
	// SQSクライアントの作成
	client := s3.New(sess)
	return &Config{
		client:     client,
		bucketName: bucketName,
		bucketKey:  bucketKey,
	}
}

func NewBadgeImageRepository(config Config) management.BadgeImgRepository {
	return BadgeImageRepository{config: config}
}

func (b BadgeImageRepository) DownloadBadge(ctx context.Context) (*management.Badge, error) {
	return nil, nil
}
