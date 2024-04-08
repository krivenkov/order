package create

import (
	"github.com/go-openapi/runtime/middleware"
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
) order.CreateOrderHandler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Handle(params order.CreateOrderParams, i interface{}) middleware.Responder {
	userID := i.(string)

	ctx := params.HTTPRequest.Context()
	l := mlog.FromContext(ctx).With(zap.String("userID", userID))
	ctx = mlog.CtxWithLogger(ctx, l)

	if params.Body == nil {
		l.Warn("request body is empty")
		return order.NewCreateOrderBadRequest().WithPayload(&models.Error{
			Error:            ptr.Pointer(models.ErrorErrorInvalidRequest),
			ErrorDescription: ptr.Pointer("request body is empty"),
		})
	}

	form := &orderModel.Form{
		Name:        params.Body.Name,
		Description: params.Body.Description,
	}

	item, err := h.service.Create(ctx, userID, form)
	if err != nil {
		l.Error("create order failed", zap.Error(err))

		return order.NewCreateOrderInternalServerError().WithPayload(&models.Error{
			Error:            ptr.Pointer(models.ErrorErrorServerError),
			ErrorDescription: ptr.Pointer("Create order failed"),
		})
	}

	return order.NewCreateOrderOK().WithPayload(&models.CreateOrderResponse{
		Order: convertors.OrderFromModel(item),
	})
}
