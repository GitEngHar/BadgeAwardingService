package management

import (
	"errors"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/segmentio/ksuid"
	"image"
	"net/url"
)

type ImageUrl string

var (
	InvalidImageUrl   = errors.New("invalid image url")
	ImageRequireParam = errors.New("image require param is empty")
)

type BadgeImg struct {
	// TODO: 想定する使われ方によってはImageURLだけでもいいかもしれない
	Image    image.Image
	ImageUrl url.URL
}

func S3BodyConvertToImage(body *s3.GetObjectOutput) (image.Image, error) {
	if body == nil {
		return nil, ImageRequireParam
	}
	defer body.Body.Close()

	img, _, err := image.Decode(body.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func NewBadgeImg(image image.Image, imageUrl url.URL) (*BadgeImg, error) {
	// TODO: ユースケースが決まったら検査を精査する
	if image == nil && imageUrl.String() == "" {
		return nil, ImageRequireParam
	}
	return &BadgeImg{
		Image:    image,
		ImageUrl: imageUrl,
	}, nil
}

type UserDTO struct {
	Name string `json:"name"`
	Mail string `json:"mail"`
}

func NewImageUrl(imageUrl string) (ImageUrl, error) {
	if isValidImageUrl(imageUrl) {
		return ImageUrl(imageUrl), nil
	}
	return "", InvalidImageUrl
}

func isValidImageUrl(imageUrl string) bool {
	_, err := url.ParseRequestURI(imageUrl)
	return err != nil
}

func NewUserID() string {
	id := ksuid.New().String()
	short := id[:26]
	return short
}
