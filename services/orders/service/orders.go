package service

import (
	"context"

	"github.com/Sanjaiy/go-grpc/services/common/genproto/orders"
)

var ordersStore = make([]*orders.Order, 0)

type OrderService struct {	
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	ordersStore = append(ordersStore, order)
	return nil
}

func (s *OrderService) GetOrders(ctx context.Context) []*orders.Order {
	return ordersStore
}