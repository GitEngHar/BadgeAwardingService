package Push

import (
	"context"
	"hello-world/domain/notification"
	"hello-world/usecase/push"
)

type UnSubscriptionHandler struct {
	uc push.UnSubscriptionUseCase
}

func NewUnSubscriptionHandler(uc push.UnSubscriptionUseCase) *UnSubscriptionHandler {
	return &UnSubscriptionHandler{uc: uc}
}

// Do 依存性解消のため値はendpointで共通
func (h *UnSubscriptionHandler) Do(ctx context.Context, endpoint notification.UnSubscriptionEndpoint) error {
	return h.uc.Do(ctx, endpoint.Address)
}
