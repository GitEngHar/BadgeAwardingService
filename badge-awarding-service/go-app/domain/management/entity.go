package management

import (
	"errors"
	"hello-world/domain"
	"time"
)

type Badge struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Reason string `json:"reason"`
}

func NewBadge(id, name, reason string) (*Badge, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}
	if id == "" {
		id = NewBadgeID(name)
	}
	return &Badge{
		ID:     id,
		Name:   name,
		Reason: reason,
	}, nil
}

func NewBadgeID(name string) string {
	return "Badge-" + name
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

func NewUserBadge(userID string, badgeRankID string, createdAt time.Time, updatedAt time.Time) *UserBadge {
	return &UserBadge{
		UserID:    userID,
		BadgeID:   badgeRankID,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

// TODO: DTOとわける
type BadgeDetailsByRank struct {
	BadgeRankID string `json:"id"`
	BadgeName   string `json:"name"`
	Rank        string `json:"rank"`
	Message     string `json:"message"`
	Effect      string `json:"effect"`
	Reason      string `json:"reason"`
}

var rankDisplayName = map[string]string{
	"B": "Bronze",
	"A": "Gold",
	"S": "Diamond",
}

func NewBadgeRankID(badgeName string, rankName string) string {
	return "BadgeRank-" + rankName + "-" + badgeName
}

func NewBadgeDetailsByRank(badgeRankID, badgeName, message, effect, reason string, rank string) (*BadgeDetailsByRank, error) {
	name, ok := rankDisplayName[rank]
	if !ok {
		return nil, errors.New("unknown rank")
	}
	if badgeRankID == "" {
		badgeRankID = NewBadgeRankID(badgeName, name)
	}
	return &BadgeDetailsByRank{
		BadgeRankID: badgeRankID,
		BadgeName:   badgeName,
		Rank:        rank,
		Message:     message,
		Effect:      effect,
		Reason:      reason,
	}, nil
}
