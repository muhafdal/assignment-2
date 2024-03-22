package order_repository

import (
	"h8-assignment-2/entity"
	"h8-assignment-2/pkg/errs"
)

type Repository interface {
	ReadOrderById(orderId int) (*entity.Order, errs.Error)
	CreateOrderWithItems(orderPayload entity.Order, itemPayload []entity.Item) errs.Error
	ReadOrders() ([]OrderItemMapped, errs.Error)
	UpdateOrder(orderPayload entity.Order, itemPayload []entity.Item) errs.Error
	DeleteOrder(orderId int) errs.Error
}
