package management

import (
	"hello-world/domain"
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
	MailAddress domain.Mail
	Name        string
}

func NewUser(email, name string) (*User, error) {
	userMail, err := domain.NewMail(email)
	if err != nil {
		return nil, err
	}
	id := NewUserID()
	return &User{
		ID:          id,
		MailAddress: userMail,
		Name:        name,
	}, nil
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
