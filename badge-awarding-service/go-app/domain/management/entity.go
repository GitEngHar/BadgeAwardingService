package management

import (
	"time"
)

type Badge struct {
	ID          string
	Name        string
	Image       ImageUrl
	Description string
}

func NewBadge(id, name, url, description string) (*Badge, error) {
	imageUrl, err := NewImageUrl(url)
	if err != nil {
		return nil, err
	}
	return &Badge{
		ID:          id,
		Name:        name,
		Image:       imageUrl,
		Description: description,
	}, nil
}

type User struct {
	ID          string
	MailAddress string
	Name        string
	BatchID     string
	GetBatchAt  time.Time
}

func NewUser(id, mailAddress, name string) *User {
	return &User{
		ID:          id,
		MailAddress: mailAddress,
		Name:        name,
	}
}

type UserBadge struct {
	UserID    string
	BadgeID   string
	GrantedAt time.Time
}

func NewUserBadge(userID string, badgeID string, grantedAt time.Time) *UserBadge {
	return &UserBadge{
		UserID:    userID,
		BadgeID:   badgeID,
		GrantedAt: grantedAt,
	}
}
