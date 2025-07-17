package management

import (
	"errors"
	"net/url"
)

type ImageUrl string

var (
	InvalidImageUrl = errors.New("invalid image url")
)

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
