package entity

import "context"

type Badge struct {
	id          string
	name        string
	imageUrl    string
	description string
}

type BadgeRepository interface {
	Upsert(ctx context.Context, badge *Badge) error
	Get(ctx context.Context, badge Badge) (Badge, error)
	Delete(ctx context.Context, badge Badge) error
}

func NewBadge(b Badge) *Badge {
	return &Badge{
		id:          b.id,
		name:        b.name,
		imageUrl:    b.imageUrl,
		description: b.description,
	}
}
