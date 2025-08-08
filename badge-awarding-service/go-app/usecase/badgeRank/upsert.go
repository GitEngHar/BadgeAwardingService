package badgeRank

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
func (u UpsertUseCase) Do(ctx context.Context, badgeRankID, badgeName, rank, message, effect, reason string) (string, error) {
	newBadgeRank, err := management.NewBadgeDetailsByRank(badgeRankID, badgeName, message, effect, reason, rank)
	if err != nil {
		return "", err
	}
	item := map[string]types.AttributeValue{
		"PK":      &types.AttributeValueMemberS{Value: newBadgeRank.BadgeRankID},
		"SK":      &types.AttributeValueMemberS{Value: newBadgeRank.Rank},
		"name":    &types.AttributeValueMemberS{Value: newBadgeRank.BadgeName},
		"message": &types.AttributeValueMemberS{Value: newBadgeRank.Message},
		"effect":  &types.AttributeValueMemberS{Value: newBadgeRank.Effect},
		"reason":  &types.AttributeValueMemberS{Value: newBadgeRank.Reason},
	}
	err = u.repo.Upsert(ctx, item)
	if err != nil {
		return "", err
	}
	return newBadgeRank.BadgeRankID, nil
}
