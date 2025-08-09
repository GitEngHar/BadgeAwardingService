package user

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"hello-world/domain/management"
	"time"
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
	badgeAward, err := u.repo.GetByPK(ctx, id)
	if err != nil {
		return nil, err
	}
	return badgeAward, nil
}

// AwardTargetUserByDate 今日の日付で報酬を通知する情報を取得する
func (c GetUseCase) AwardTargetUserByDate(ctx context.Context) ([]map[string]types.AttributeValue, error) {
	awardTargetDate := time.Now().Format("2006-01-02")
	badgeAwardTarget, err := c.repo.GetsByPK(ctx, awardTargetDate)
	if err != nil {
		return nil, err
	}
	return badgeAwardTarget, nil
}

func getStringAttr(item map[string]types.AttributeValue, key string) (string, bool) {
	if v, ok := item[key]; ok { // 存在チェック
		if av, ok := v.(*types.AttributeValueMemberS); ok { // S型チェック
			return av.Value, true
		}
	}
	return "", false // 無い or S型じゃない
}
