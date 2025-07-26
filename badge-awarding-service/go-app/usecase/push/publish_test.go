package push

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type mockPublisher struct {
	body  string
	attrs map[string]types.MessageAttributeValue
}

func (m *mockPublisher) PublishMailMessage(ctx context.Context, body string, attrs map[string]types.MessageAttributeValue) error {
	m.body = body
	m.attrs = attrs
	return nil
}
func (m *mockPublisher) GetMailMessage(ctx context.Context) ([]types.Message, error) { return nil, nil }

func TestPublishMessageUseCase_Do(t *testing.T) {
	repo := &mockPublisher{}
	uc := NewPublishMessageUseCase(repo)

	err := uc.Do(context.Background(), "body", "Bob", "bob@example.com", "hello")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if repo.body != "body" {
		t.Errorf("body mismatch: %s", repo.body)
	}

	if v := repo.attrs["userName"].StringValue; v == nil || *v != "Bob" {
		t.Errorf("unexpected userName attr: %v", repo.attrs["userName"])
	}
	if v := repo.attrs["address"].StringValue; v == nil || *v != "bob@example.com" {
		t.Errorf("unexpected address attr: %v", repo.attrs["address"])
	}
	if v := repo.attrs["message"].StringValue; v == nil || *v != "hello" {
		t.Errorf("unexpected message attr: %v", repo.attrs["message"])
	}
}
