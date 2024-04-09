package count

import (
	"github.com/go-openapi/runtime/middleware"
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

func New(service orderModel.Service) order.GetOrdersCountHandler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Handle(params order.GetOrdersCountParams, i interface{}) middleware.Responder {
	userID := i.(string)

	ctx := params.HTTPRequest.Context()
	l := mlog.FromContext(ctx).With(
		zap.String("userID", userID),
	)
	ctx = mlog.CtxWithLogger(ctx, l)

	count, err := h.service.Count(ctx, userID, h.prepareCountCondition(params))
	if err != nil {
		l.Error("get order count failed", zap.Error(err))

		return order.NewGetOrdersCountInternalServerError().WithPayload(&models.Error{
			Error:            ptr.Pointer(models.ErrorErrorServerError),
			ErrorDescription: ptr.Pointer("Get order count failed"),
		})
	}

	return order.NewGetOrdersCountOK().WithPayload(&models.GetCountResponse{
		Count: ptr.Pointer(int64(count)),
	})
}

func (h *Handler) prepareCountCondition(_ order.GetOrdersCountParams) *orderModel.GetCountRequest {
	return &orderModel.GetCountRequest{}
}
