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
	// Assuming CreateOrderRequest has fields: OrderId, Items, Amount, Status
	order := &orders.Order{
		OrderId:    11,
		CustomerId: 2,
		ProductId:  3,
		Quantity:   14,
	}

	err := h.ordersService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return &orders.CreateOrderResponse{
		Status: "success",
	}, nil
}
