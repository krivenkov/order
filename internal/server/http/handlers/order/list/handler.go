package list

import (
	"github.com/go-openapi/runtime/middleware"
	orderModel "github.com/krivenkov/order/internal/model/order"
	"github.com/krivenkov/order/internal/server/http/convertors"
	"github.com/krivenkov/order/internal/server/http/models"
	orderOperation "github.com/krivenkov/order/internal/server/http/operations/order"
	"github.com/krivenkov/pkg/mlog"
	"github.com/krivenkov/pkg/option"
	"github.com/krivenkov/pkg/order"
	"github.com/krivenkov/pkg/paginator"
	"github.com/krivenkov/pkg/ptr"
	"go.uber.org/zap"
)

type Handler struct {
	service orderModel.Service
}

func New(service orderModel.Service) orderOperation.GetOrdersHandler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Handle(params orderOperation.GetOrdersParams, i interface{}) middleware.Responder {
	userID := i.(string)

	ctx := params.HTTPRequest.Context()
	l := mlog.FromContext(ctx).With(
		zap.String("userID", userID),
		zap.Float64p("offset", params.Offset),
		zap.Float64p("limit", params.Limit),
		zap.Stringp("sortBy", params.SortBy),
		zap.Stringp("sortDirection", params.SortDirection),
		zap.Stringp("q", params.Q),
	)
	ctx = mlog.CtxWithLogger(ctx, l)

	listReq := h.prepareListCondition(params)
	if !listReq.Pagination.IsSet() {
		l.Error("empty pagination")
		return orderOperation.NewGetOrdersBadRequest().WithPayload(&models.Error{
			Error:            ptr.Pointer(models.ErrorErrorInvalidRequest),
			ErrorDescription: ptr.Pointer("Get list failed"),
		})
	}

	list, err := h.service.GetList(ctx, userID, listReq)
	if err != nil {
		l.Error("order get list failed", zap.Error(err))
		return orderOperation.NewGetOrdersInternalServerError().WithPayload(&models.Error{
			Error:            ptr.Pointer(models.ErrorErrorServerError),
			ErrorDescription: ptr.Pointer("Get list failed"),
		})
	}

	total, err := h.service.Count(ctx, userID, h.prepareCountCondition(params))
	if err != nil {
		l.Error("get order count failed", zap.Error(err))
		return orderOperation.NewGetOrdersInternalServerError().WithPayload(&models.Error{
			Error:            ptr.Pointer(models.ErrorErrorServerError),
			ErrorDescription: ptr.Pointer("Get list failed"),
		})
	}

	return orderOperation.NewGetOrdersOK().WithPayload(&models.GetOrdersResponse{
		Orders: convertors.OrdersFromModel(list),
		Pagination: convertors.Pagination(&paginator.PaginationResult{
			Limit:  listReq.Pagination.Value().Limit,
			Offset: listReq.Pagination.Value().Offset,
			Total:  total,
		}),
	})
}

func (h *Handler) prepareListCondition(params orderOperation.GetOrdersParams) *orderModel.GetListRequest {
	req := &orderModel.GetListRequest{}

	ordering := convertors.Order(params.SortBy, params.SortDirection)
	pagination := convertors.Paginator(params.Limit, params.Offset)

	if ordering != nil {
		req.Orders = option.New([]*order.Order{ordering})
	}

	if pagination != nil {
		req.Pagination = option.New(*pagination)
	}

	if params.Q != nil {
		req.Q = option.New(*params.Q)
	}

	return req
}

func (h *Handler) prepareCountCondition(params orderOperation.GetOrdersParams) *orderModel.GetCountRequest {
	req := &orderModel.GetCountRequest{}

	if params.Q != nil {
		req.Q = option.New(*params.Q)
	}

	return req
}
