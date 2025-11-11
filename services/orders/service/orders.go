package service

import (
	"context"
	"log"

	"github.com/uchidas-rogue/kitchen/services/common/genproto/orders"
)

var ordersDb = make([]*orders.Order, 0)

type OrderService struct{}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	log.Printf("Creating order: %+v\n", order)
	ordersDb = append(ordersDb, order)
	return nil
}

func (s *OrderService) GetOrders(ctx context.Context) []*orders.Order {
	return ordersDb
}
