package management

import "context"

type BadgeRepository interface {
	Upsert(ctx context.Context, badge *Badge) error
	GetByID(ctx context.Context, id string) (Badge, error)
	Delete(ctx context.Context, id string) error
}

type UserRepository interface {
	Upsert(ctx context.Context, user User) error
	Create(ctx context.Context, user User) error
	GetByID(ctx context.Context, id string) (User, error)
	Delete(ctx context.Context, id string) error
}

type UserBadgeRepository interface {
	Upsert(ctx context.Context, userBadge UserBadge) error
	Create(ctx context.Context, userBadge UserBadge) error
	GetByID(ctx context.Context, id string) (User, error)
	Delete(ctx context.Context, id string) error
}
