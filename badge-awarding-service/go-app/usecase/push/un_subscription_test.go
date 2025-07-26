package push

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"hello-world/domain/notification"
)

type mockSubscriberRepo struct {
	endpoint string
}

func (m *mockSubscriberRepo) SubscribeEmail(ctx context.Context, endpoint string) error { return nil }
func (m *mockSubscriberRepo) UnSubscribeByEndpoint(ctx context.Context, endpoint string) error {
	m.endpoint = endpoint
	return nil
}
func (m *mockSubscriberRepo) SendMessageToEmail(ctx context.Context, publisher notification.Publisher) error {
	return nil
}

func TestUnSubscriptionUseCase_Do(t *testing.T) {
	repo := &mockSubscriberRepo{}
	uc := NewUnSubscriptionUseCase(repo)

	if err := uc.Do(context.Background(), "addr@example.com"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if diff := cmp.Diff("addr@example.com", repo.endpoint); diff != "" {
		t.Errorf("endpoint mismatch (-want +got):\n%s", diff)
	}
}
