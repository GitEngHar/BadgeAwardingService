package user

import (
	"context"
	"github.com/google/go-cmp/cmp/cmpopts"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/go-cmp/cmp"
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

	idAttr, ok := repo.item["id"].(*types.AttributeValueMemberS)
	if !ok || idAttr.Value == "" {
		t.Fatalf("id attribute not set: %+v", repo.item["id"])
	}

	expected := map[string]types.AttributeValue{
		"id":   idAttr,
		"name": &types.AttributeValueMemberS{Value: "Alice"},
		"mail": &types.AttributeValueMemberS{Value: "user@example.com"},
	}
	if diff := cmp.Diff(expected, repo.item,
		cmpopts.IgnoreUnexported(
			types.AttributeValueMemberS{},
		),
	); diff != "" {
		t.Errorf("item mismatch (-want +got):\n%s", diff)
	}
}
