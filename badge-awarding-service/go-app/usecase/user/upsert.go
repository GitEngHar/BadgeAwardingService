package user

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"hello-world/domain/management"
)

type UpsertUseCase struct {
	repo management.UserBadgeRepository
}

func NewUpsertUseCase(repo management.UserBadgeRepository) UpsertUseCase {
	return UpsertUseCase{
		repo: repo,
	}
}

func (u UpsertUseCase) Do(ctx context.Context, email string, name string) error {
	newUser, err := management.NewUser(email, name)
	if err != nil {
		return err
	}
	item := map[string]types.AttributeValue{
		"id":   &types.AttributeValueMemberS{Value: newUser.ID},
		"name": &types.AttributeValueMemberS{Value: newUser.Name},
		"mail": &types.AttributeValueMemberS{Value: string(newUser.MailAddress)},
	}
	return u.repo.Upsert(ctx, item)
}
