package push

import (
	"context"
	"github.com/google/go-cmp/cmp/cmpopts"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/google/go-cmp/cmp"
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

	if diff := cmp.Diff("body", repo.body); diff != "" {
		t.Errorf("body mismatch (-want +got):\n%s", diff)
	}

	expected := map[string]types.MessageAttributeValue{
		"userName": {DataType: aws.String("String"), StringValue: aws.String("Bob")},
		"message":  {DataType: aws.String("String"), StringValue: aws.String("hello")},
		"address":  {DataType: aws.String("String"), StringValue: aws.String("bob@example.com")},
	}
	if diff := cmp.Diff(expected, repo.attrs,
		cmpopts.IgnoreUnexported(types.MessageAttributeValue{}),
	); diff != "" {
		t.Errorf("message attributes mismatch (-want +got):\n%s", diff)
	}

}
