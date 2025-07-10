package domain

import "context"

type Batch struct {
	id          string
	name        string
	imageUrl    string
	description string
}

type BatchFactory struct {
	b Batch
}

type BatchRepository interface {
	Create(ctx context.Context, batch Batch) error
	Get(ctx context.Context, batch Batch) (Batch, error)
	Update(ctx context.Context, batch Batch) error
	Delete(ctx context.Context, batch Batch) error
}
