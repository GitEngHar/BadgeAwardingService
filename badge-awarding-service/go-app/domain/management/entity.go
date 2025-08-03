package management

import (
	"errors"
	"hello-world/domain"
	"time"
)

type Badge struct {
	ID     string
	Name   string
	Reason string
}

func NewBadge(id, name, reason string) (*Badge, error) {
	return &Badge{
		ID:     id,
		Name:   name,
		Reason: reason,
	}, nil
}

type User struct {
	ID          string
	MailAddress domain.Mail
	Name        string
}

func NewUser(userID, email, name string) (*User, error) {
	userMail, err := domain.NewMail(email)
	if err != nil {
		return nil, err
	}
	if userID == "" {
		userID = NewUserID()
	}
	return &User{
		ID:          userID,
		MailAddress: userMail,
		Name:        name,
	}, nil
}

type UserBadge struct {
	UserID    string
	BadgeID   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserBadge(userID string, badgeID string, createdAt time.Time, updatedAt time.Time) *UserBadge {
	return &UserBadge{
		UserID:    userID,
		BadgeID:   badgeID,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

type BadgeDetailsByRank struct {
	BadgeID   string
	Rank      string
	Message   string
	Effect    string
	Reason    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewBadgeDetailsByRank(badgeID, Message, Effect, Reason string, rank Rank, createdAt time.Time, updatedAt time.Time) (*BadgeDetailsByRank, error) {
	rankString := rank.String()
	if rankString == "Unknown" {
		return nil, errors.New("unknown rank")
	}

	return &BadgeDetailsByRank{
		BadgeID:   badgeID,
		Rank:      rankString,
		Message:   Message,
		Effect:    Effect,
		Reason:    Reason,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}
