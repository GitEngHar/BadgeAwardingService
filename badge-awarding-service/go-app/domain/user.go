package domain

import (
	"context"
	"time"
)

type User struct {
	id          string
	mailAddress string
	name        string
	batchID     string
	getBatchAt  time.Time
}

type UserRepository interface {
	Create(ctx context.Context, user User) error
	Get(ctx context.Context, user User) (User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, user User) error
}

func NewUser(u User) *User {
	return &User{
		id:          u.id,
		mailAddress: u.mailAddress,
		name:        u.name,
		batchID:     u.batchID,
		getBatchAt:  u.getBatchAt,
	}
}
