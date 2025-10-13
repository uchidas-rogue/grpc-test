package types

import (
	"context"

	"github.com/uchidas-rogue/kitchen/services/common/genproto/orders"
)

type OrderService interface {
	// Define methods for the OrderService here
	CreateOrder(context.Context, *orders.Order) error
	GetOrders(context.Context) []*orders.Order
}
