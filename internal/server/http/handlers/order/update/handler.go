package update

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

func New(
	service orderModel.Service,
) order.UpdateOrderHandler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Handle(params order.UpdateOrderParams, i interface{}) middleware.Responder {
	if _, err := uuid.Parse(params.ID); err != nil {
		return order.NewUpdateOrderNotFound().WithPayload(&models.Error{
			Error:            ptr.Pointer(models.ErrorErrorInvalidRequest),
			ErrorDescription: ptr.Pointer("Not Found"),
		})
	}

	userID := i.(string)

	ctx := params.HTTPRequest.Context()
	l := mlog.FromContext(ctx).With(zap.String("userID", userID))
	ctx = mlog.CtxWithLogger(ctx, l)

	if params.Body == nil {
		l.Warn("request body is empty")
		return order.NewUpdateOrderBadRequest().WithPayload(&models.Error{
			Error:            ptr.Pointer(models.ErrorErrorInvalidRequest),
			ErrorDescription: ptr.Pointer("request body is empty"),
		})
	}

	item, err := h.service.Update(ctx, userID, params.ID, &orderModel.Form{
		Name:        params.Body.Name,
		Description: params.Body.Description,
	})
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return order.NewUpdateOrderNotFound().WithPayload(&models.Error{
				Error:            ptr.Pointer(models.ErrorErrorInvalidRequest),
				ErrorDescription: ptr.Pointer(err.Error()),
			})
		}

		if errors.Is(err, model.ErrPermissionDenied) {
			return order.NewUpdateOrderForbidden().WithPayload(&models.Error{
				Error:            ptr.Pointer(models.ErrorErrorAccessDenied),
				ErrorDescription: ptr.Pointer(err.Error()),
			})
		}

		l.Error("update order failed", zap.Error(err))

		return order.NewUpdateOrderInternalServerError().WithPayload(&models.Error{
			Error:            ptr.Pointer(models.ErrorErrorServerError),
			ErrorDescription: ptr.Pointer("Update order failed"),
		})
	}

	return order.NewUpdateOrderOK().WithPayload(&models.UpdateOrderResponse{
		Order: convertors.OrderFromModel(item),
	})
}
