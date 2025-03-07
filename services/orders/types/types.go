package types

import (
	"context"

	"github.com/Sanjaiy/go-grpc/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
	GetOrders(context.Context) []*orders.Order
}