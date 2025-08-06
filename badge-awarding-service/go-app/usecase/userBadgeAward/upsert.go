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

func (u UpsertUseCase) Do(ctx context.Context, userBadgeAwardID, userID, badgeRankID string) (string, error) {
	newUserBadgeAward, err := management.NewUserBadgeAward(userBadgeAwardID, userID, badgeRankID)
	if err != nil {
		return "", err
	}
	item := map[string]types.AttributeValue{
		"id":            &types.AttributeValueMemberS{Value: newUserBadgeAward.UserBadgeAwardID},
		"user_id":       &types.AttributeValueMemberS{Value: newUserBadgeAward.UserID},
		"badge_rank_id": &types.AttributeValueMemberS{Value: newUserBadgeAward.BadgeRankID},
		"update_at":     &types.AttributeValueMemberS{Value: newUserBadgeAward.UpdateAt.Format("2006-01-02")},
	}
	err = u.repo.Upsert(ctx, item)
	if err != nil {
		return "", err
	}
	itemForPush := map[string]types.AttributeValue{
		"id":            &types.AttributeValueMemberS{Value: newUserBadgeAward.UpdateAt.Format("2006-01-02")},
		"badge_id":      &types.AttributeValueMemberS{Value: newUserBadgeAward.UserID},
		"user_id":       &types.AttributeValueMemberS{Value: newUserBadgeAward.UserID},
		"badge_rank_id": &types.AttributeValueMemberS{Value: newUserBadgeAward.BadgeRankID},
	}
	err = u.repo.Upsert(ctx, itemForPush)
	if err != nil {
		return "", err
	}
	return newUserBadgeAward.UserBadgeAwardID, nil
}
