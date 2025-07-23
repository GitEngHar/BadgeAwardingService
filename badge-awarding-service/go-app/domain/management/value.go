package management

import (
	"errors"
	"github.com/segmentio/ksuid"
	"net/url"
)

type ImageUrl string

var (
	InvalidImageUrl = errors.New("invalid image url")
)

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
