package management

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/segmentio/ksuid"
	"hello-world/domain"
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
	ID   string `json:"id"`
	Name string `json:"name"`
	Mail string `json:"mail"`
}

type UserID struct {
	ID string `json:"id"`
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

func MapUserToDomain(userMap map[string]types.AttributeValue) (*User, error) {
	userID, isExistsKey := domain.GetStringAttr(userMap, "PK")
	if !isExistsKey {
		return nil, errors.New("OK does not exist key")
	}
	name, isExistsKey := domain.GetStringAttr(userMap, "SK")
	if !isExistsKey {
		return nil, errors.New("SK does not exist key")
	}
	mail, isExistsKey := domain.GetStringAttr(userMap, "mail")
	if !isExistsKey {
		return nil, errors.New("mail does not exist key")
	}
	userDomain, err := NewUser(userID, mail, name)
	if err != nil {
		return nil, err
	}
	return userDomain, nil
}

func MapBadgeRankToDomain(badgeRankIDMap map[string]types.AttributeValue) (*BadgeDetailsByRank, error) {
	badgeRankID, isExistsKey := domain.GetStringAttr(badgeRankIDMap, "PK")
	if !isExistsKey {
		return nil, errors.New("PK does not exist key")
	}
	name, isExistsKey := domain.GetStringAttr(badgeRankIDMap, "name")
	if !isExistsKey {
		return nil, errors.New("name does not exist key")
	}
	effect, isExistsKey := domain.GetStringAttr(badgeRankIDMap, "effect")
	if !isExistsKey {
		return nil, errors.New("effect does not exist key")
	}
	message, isExistsKey := domain.GetStringAttr(badgeRankIDMap, "message")
	if !isExistsKey {
		return nil, errors.New("message does not exist key")
	}
	reason, isExistsKey := domain.GetStringAttr(badgeRankIDMap, "reason")
	if !isExistsKey {
		return nil, errors.New("reason does not exist key")
	}
	rank, isExistsKey := domain.GetStringAttr(badgeRankIDMap, "SK")
	if !isExistsKey {
		return nil, errors.New("SK does not exist key")
	}
	badgeRankDomain, err := NewBadgeDetailsByRank(badgeRankID, name, message, effect, reason, rank)
	if err != nil {
		return nil, err
	}
	return badgeRankDomain, nil
}
