package user

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"hello-world/domain/management"
)

type UpsertUseCase struct {
	repo management.Repository
}

func NewUpsertUseCase(repo management.Repository) UpsertUseCase {
	return UpsertUseCase{
		repo: repo,
	}
}

// Do Userを追加し、IDが存在する場合は更新する(少人数で利用することを想定しているのでIDは重複しない)
func (u UpsertUseCase) Do(ctx context.Context, userID string, email string, name string) (string, error) {
	newUser, err := management.NewUser(userID, email, name)
	if err != nil {
		return "", err
	}
	item := map[string]types.AttributeValue{
		"PK":   &types.AttributeValueMemberS{Value: newUser.ID},
		"SK":   &types.AttributeValueMemberS{Value: newUser.Name},
		"mail": &types.AttributeValueMemberS{Value: string(newUser.MailAddress)},
	}
	err = u.repo.Upsert(ctx, item)
	if err != nil {
		return "", err
	}
	return newUser.ID, nil
}
