package item

import (
	"errors"

	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"
	"github.com/krivenkov/order/internal/model"
	orderModel "github.com/krivenkov/order/internal/model/order"
	"github.com/krivenkov/order/internal/server/http/convertors"
	"github.com/krivenkov/order/internal/server/http/models"
	"github.com/krivenkov/order/internal/server/http/operations/order"
	"github.com/krivenkov/pkg/mlog"
	"github.com/krivenkov/pkg/ptr"
	"go.uber.org/zap"
)

type Handler struct {
	service orderModel.Service
}

func New(service orderModel.Service) order.GetOrderHandler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Handle(params order.GetOrderParams, i interface{}) middleware.Responder {
	userID := i.(string)

	if _, err := uuid.Parse(params.ID); err != nil {
		return order.NewGetOrderNotFound().WithPayload(&models.Error{
			Error:            ptr.Pointer(models.ErrorErrorInvalidRequest),
			ErrorDescription: ptr.Pointer("Not Found"),
		})
	}

	ctx := params.HTTPRequest.Context()
	l := mlog.FromContext(ctx).With(
		zap.String("userID", userID),
		zap.String("orderID", params.ID),
	)
	ctx = mlog.CtxWithLogger(ctx, l)

	orderItem, err := h.service.GetItem(ctx, userID, params.ID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return order.NewGetOrderNotFound().WithPayload(&models.Error{
				Error:            ptr.Pointer(models.ErrorErrorInvalidRequest),
				ErrorDescription: ptr.Pointer(err.Error()),
			})
		}

		if errors.Is(err, model.ErrPermissionDenied) {
			return order.NewGetOrderForbidden().WithPayload(&models.Error{
				Error:            ptr.Pointer(models.ErrorErrorAccessDenied),
				ErrorDescription: ptr.Pointer(err.Error()),
			})
		}

		l.Error("get order failed", zap.Error(err))

		return order.NewGetOrderInternalServerError().WithPayload(&models.Error{
			Error:            ptr.Pointer(models.ErrorErrorServerError),
			ErrorDescription: ptr.Pointer("Get Order Failed"),
		})
	}

	return order.NewGetOrderOK().WithPayload(&models.GetOrderResponse{
		Order: convertors.OrderFromModel(orderItem),
	})
}
