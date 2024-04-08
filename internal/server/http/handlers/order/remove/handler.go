package remove

import (
	"errors"

	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
	"github.com/krivenkov/order/internal/model"
	orderModel "github.com/krivenkov/order/internal/model/order"
	"github.com/krivenkov/order/internal/server/http/models"
	"github.com/krivenkov/order/internal/server/http/operations/order"
	"github.com/krivenkov/pkg/mlog"
	"github.com/krivenkov/pkg/ptr"
	"go.uber.org/zap"
)

type Handler struct {
	service orderModel.Service
}

func New(service orderModel.Service) order.DeleteOrderHandler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Handle(params order.DeleteOrderParams, i interface{}) middleware.Responder {
	if _, err := uuid.Parse(params.ID); err != nil {
		return order.NewDeleteOrderNotFound().WithPayload(&models.Error{
			Error:            ptr.Pointer(models.ErrorErrorInvalidRequest),
			ErrorDescription: ptr.Pointer("Not Found"),
		})
	}

	userID := i.(string)

	ctx := params.HTTPRequest.Context()
	l := mlog.FromContext(ctx).With(
		zap.String("userID", userID),
		zap.String("orderID", params.ID),
	)
	ctx = mlog.CtxWithLogger(ctx, l)

	if err := h.service.SoftDelete(ctx, userID, params.ID); err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return order.NewDeleteOrderNotFound().WithPayload(&models.Error{
				Error:            ptr.Pointer(models.ErrorErrorInvalidRequest),
				ErrorDescription: ptr.Pointer(err.Error()),
			})
		}

		if errors.Is(err, model.ErrPermissionDenied) {
			return order.NewDeleteOrderForbidden().WithPayload(&models.Error{
				Error:            ptr.Pointer(models.ErrorErrorAccessDenied),
				ErrorDescription: ptr.Pointer(err.Error()),
			})
		}

		l.Error("delete order failed", zap.Error(err))

		return order.NewDeleteOrderInternalServerError().WithPayload(&models.Error{
			Error:            ptr.Pointer(models.ErrorErrorServerError),
			ErrorDescription: ptr.Pointer("Delete order failed"),
		})
	}

	return order.NewDeleteOrderNoContent()
}
