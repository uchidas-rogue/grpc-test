package handler

import (
	"net/http"

	"github.com/uchidas-rogue/kitchen/services/common/genproto/orders"
	"github.com/uchidas-rogue/kitchen/services/common/util"
	"github.com/uchidas-rogue/kitchen/services/orders/types"
)

type OrdersHttpHandler struct {
	ordersService types.OrderService
}

func NewHttpOrdersHandler(orderService types.OrderService) *OrdersHttpHandler {
	handler := &OrdersHttpHandler{
		ordersService: orderService,
	}

	return handler
}

func (h *OrdersHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	// Parse the JSON request marshaling
	err := util.ParseJSON(r, &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	order := &orders.Order{
		OrderId:    42,
		CustomerId: req.GetCustomerId(),
		ProductId:  req.GetProductId(),
		Quantity:   req.GetQuantity(),
	}

	err = h.ordersService.CreateOrder(r.Context(), order)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &orders.CreateOrderResponse{Status: "success"}
	util.WriteJSON(w, http.StatusOK, res)
}
