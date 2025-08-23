package management

import (
	"errors"
	"hello-world/domain"
	"time"
)

// Badge ユーザーに付与するバッチ
type Badge struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Reason string `json:"reason"`
}

type BadgeDTO struct {
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

// User ユーザー情報
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

// UserBadgeAward ユーザーが取得しているバッチの情報
type UserBadgeAwardDTO struct {
	UserBadgeAwardID string `json:"id"`
	UserID           string `json:"user_id"`
	BadgeRankID      string `json:"badge_rank_id"`
	UpdateAt         time.Time
}

type UserBadgeAward struct {
	UserBadgeAwardID string `json:"id"`
	UserID           string `json:"user_id"`
	BadgeRankID      string `json:"badge_rank_id"`
	UpdateAt         time.Time
}

func NewUserBadgeAward(userBadgeAwardID string, userID string, badgeRankID string) (*UserBadgeAward, error) {
	// IDを指定するには直接指定されているかIDを作成するのに必要なuserIDとbadgeRankIDの両方の値が必要
	if userBadgeAwardID == "" && (userID == "" || badgeRankID == "") {
		return nil, errors.New("userID or userBadgeAwardID is empty")
	}
	if userBadgeAwardID == "" {
		userBadgeAwardID = NewBadgeAwardID(userID, badgeRankID)
	}
	return &UserBadgeAward{
		UserBadgeAwardID: userBadgeAwardID,
		UserID:           userID,
		BadgeRankID:      badgeRankID,
		UpdateAt:         time.Now(),
	}, nil
}

func NewBadgeAwardID(userID string, badgeRankID string) string {
	return "Badge-Award-" + userID + "-" + badgeRankID
}

// TODO: DTOとわける

// BadgeDetailsByRank バッチのランク別情報
type BadgeDetailsByRankDTO struct {
	BadgeRankID string `json:"id"`
	BadgeName   string `json:"name"`
	Rank        string `json:"rank"`
	Message     string `json:"message"`
	Effect      string `json:"effect"`
	Reason      string `json:"reason"`
}

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
