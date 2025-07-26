package user

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type mockUserBadgeRepo struct {
	item map[string]types.AttributeValue
}

func (m *mockUserBadgeRepo) Upsert(ctx context.Context, item map[string]types.AttributeValue) error {
	m.item = item
	return nil
}
func (m *mockUserBadgeRepo) Get(ctx context.Context, item map[string]types.AttributeValue) (map[string]types.AttributeValue, error) {
	return nil, nil
}
func (m *mockUserBadgeRepo) Del(ctx context.Context, filter map[string]types.AttributeValue) error {
	return nil
}
func (m *mockUserBadgeRepo) CreateTable(ctx context.Context) error { return nil }

func TestUpsertUseCase_Do(t *testing.T) {
	repo := &mockUserBadgeRepo{}
	uc := NewUpsertUseCase(repo)

	if err := uc.Do(context.Background(), "user@example.com", "Alice"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if repo.item == nil {
		t.Fatal("Upsert was not called")
	}

	nameAttr, ok := repo.item["name"].(*types.AttributeValueMemberS)
	if !ok || nameAttr.Value != "Alice" {
		t.Errorf("unexpected name attribute: %+v", repo.item["name"])
	}
	mailAttr, ok := repo.item["mail"].(*types.AttributeValueMemberS)
	if !ok || mailAttr.Value != "user@example.com" {
		t.Errorf("unexpected mail attribute: %+v", repo.item["mail"])
	}
	idAttr, ok := repo.item["id"].(*types.AttributeValueMemberS)
	if !ok || idAttr.Value == "" {
		t.Errorf("id attribute not set: %+v", repo.item["id"])
	}
}
