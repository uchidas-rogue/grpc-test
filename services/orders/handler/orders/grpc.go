package handler

import (
	"context"

	"github.com/uchidas-rogue/kitchen/services/common/genproto/orders"
	"github.com/uchidas-rogue/kitchen/services/orders/types"
	"google.golang.org/grpc"
)

type OrderGrpcHandler struct {
	ordersService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrdersService(grpc *grpc.Server, orderService types.OrderService) {
	gRPCHandler := &OrderGrpcHandler{
		ordersService: orderService,
	}
	orders.RegisterOrderServiceServer(grpc, gRPCHandler)
}

func (h *OrderGrpcHandler) GetOrders(ctx context.Context, req *orders.GetOrderRequest) (*orders.GetOrderResponse, error) {
	o := h.ordersService.GetOrders(ctx)
	res := &orders.GetOrderResponse{
		Orders: o,
	}
	return res, nil
}

func (h *OrderGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {

	oList := h.ordersService.GetOrders(ctx)
	// 最新のorderIdを取得してインクリメント
	var maxOrderID int32 = 0
	for _, o := range oList {
		if o.OrderId > maxOrderID {
			maxOrderID = o.OrderId
		}
	}
	newOrderID := maxOrderID + 1

	// Assuming CreateOrderRequest has fields: OrderId, Items, Amount, Status
	order := &orders.Order{
		OrderId:    newOrderID,
		CustomerId: req.GetCustomerId(),
		ProductId:  req.GetProductId(),
		Quantity:   req.GetQuantity(),
	}

	err := h.ordersService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return &orders.CreateOrderResponse{
		Status: "success",
	}, nil
}
