package s3

import (
	"net/http"
	"time"
)

type BadgeImageRepository struct {
	client *http.Client
}

func NewBadgeImageRepository() *BadgeImageRepository {
	client := http.Client{
		Timeout: time.Second * 5,
	}

	return &BadgeImageRepository{
		client: &client,
	}
}

func (b BadgeImageRepository) IsValidS3Image(badgeImgUrl string) (bool, error) {
	// Badgeの画像を取得
	resp, err := b.client.Get(badgeImgUrl)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK, nil
}
