package user_handler

import (
	"context"
	"fmt"

	"github.com/krivenkov/order/internal/model/order"
	"github.com/krivenkov/order/internal/model/user"
	"github.com/krivenkov/pkg/bus"
)

type Handler struct {
	service order.Service
}

func New(service order.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (p *Handler) Handle(ctx context.Context, record bus.Message[user.User]) *bus.HandleResult {
	if record.Value.Payload == nil {
		return &bus.HandleResult{
			Err:  fmt.Errorf("empty record"),
			Code: bus.StatusWarning,
		}
	}

	if !record.Value.Payload.Disable {
		return &bus.HandleResult{
			Code: bus.StatusOk,
		}
	}

	if err := p.service.Disable(ctx, record.Value.Payload.ID); err != nil {
		return &bus.HandleResult{
			Code: bus.StatusError,
			Err:  err,
		}
	}

	return &bus.HandleResult{
		Code: bus.StatusOk,
	}
}
