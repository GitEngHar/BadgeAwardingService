package user

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"hello-world/domain/management"
)

type GetUseCase struct {
	repo management.Repository
}

func NewGetUseCase(repo management.Repository) GetUseCase {
	return GetUseCase{
		repo: repo,
	}
}

func (u GetUseCase) Do(ctx context.Context, id string) (map[string]types.AttributeValue, error) {
	user, err := u.repo.GetByPK(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
