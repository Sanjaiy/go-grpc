package handler

import (
	"context"

	"github.com/Sanjaiy/go-grpc/services/common/genproto/orders"
	"github.com/Sanjaiy/go-grpc/services/orders/types"
	"google.golang.org/grpc"
)

type OrderGprcHandler struct {
	orderService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrderService(grpc *grpc.Server, orderService types.OrderService) {
	gRPCHandler := &OrderGprcHandler{
		orderService: orderService,
	}

	orders.RegisterOrderServiceServer(grpc, gRPCHandler)
}


func (h *OrderGprcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID: 42,
		CustomerID: 2,
		ProductID: 1,
		Quantity: 1,
	}

	err := h.orderService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "Sucess",
	}

	return res, nil
}


func (h *OrderGprcHandler) GetOrders(ctx context.Context, req *orders.GetOrdersRequest) (*orders.GetOrderResponse, error) {
	o := h.orderService.GetOrders(ctx)

	res := &orders.GetOrderResponse{
		Orders: o,	
	}

	return res, nil
}