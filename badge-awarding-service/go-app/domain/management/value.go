package management

import (
	"errors"
	"github.com/segmentio/ksuid"
	"image"
	"net/url"
)

type ImageUrl string

var (
	InvalidImageUrl = errors.New("invalid image url")
)

type BadgeImg struct {
	Name string
	Type string
	// TODO: 想定する使われ方によってはImageURLだけでもいいかもしれない
	Image    image.Image
	ImageUrl url.URL
}

func NewBadgeImg(name, imageType string, image image.Image, imageUrl url.URL) (*BadgeImg, error) {
	//TODO: imageの検査とURLを検査する
	return &BadgeImg{
		Name:     name,
		Type:     imageType,
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
