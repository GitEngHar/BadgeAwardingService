package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"hello-world/domain/management"
	"net/url"
)

type BadgeImageRepository struct {
	config Config
}

func NewConfig(bucketName, bucketKey string) *Config {
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

// DownloadBadge S3から画像を取得してオブジェクトに画像情報を格納して渡す。画像を配信する際に利用する
func (b BadgeImageRepository) DownloadBadge() (*management.BadgeImg, error) {
	// Badgeの画像を取得
	result, err := b.config.client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(b.config.bucketName),
		Key:    aws.String(b.config.bucketKey),
	})
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()
	img, err := management.S3BodyConvertToImage(result)
	if err != nil {
		return nil, err
	}
	// TODO: ユースケースがはっきりしたら精査する
	imgUrl := url.URL{}
	badgeImg, err := management.NewBadgeImg(img, imgUrl)
	return badgeImg, err
}
