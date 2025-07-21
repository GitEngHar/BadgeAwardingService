package handler

import (
	"context"
	"hello-world/domain/notification"
)

type UnSubscriptionHandler interface {
	Do(ctx context.Context, endpoint notification.UnSubscriptionEndpoint) error
}
