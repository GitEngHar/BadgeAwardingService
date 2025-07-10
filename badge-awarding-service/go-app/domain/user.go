package domain

import (
	"context"
	"time"
)

type User struct {
	id         string
	name       string
	batchID    string
	getBatchAt time.Time
}

type UserFactory struct {
	u User
}

type UserRepository interface {
	Create(ctx context.Context, user User) error
	Get(ctx context.Context, user User) (User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, user User) error
}

func NewUserFactory(u User) *UserFactory {
	return &UserFactory{u}
}
