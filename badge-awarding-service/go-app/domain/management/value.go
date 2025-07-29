package management

import (
	"errors"
	"fmt"
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

func CreatePublicBadgeImgUrl(bucket, key string) (*string, error) {
	if bucket == "" && key == "" {
		return nil, ImageRequireParam
	}
	badgeImgUrl := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, "ap-northeast-1", url.PathEscape(key))
	return &badgeImgUrl, nil
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
