package order_pg

import (
	"database/sql"
	"errors"
	"h8-assignment-2/entity"
	"h8-assignment-2/pkg/errs"
	"h8-assignment-2/repository/order_repository"
)

type orderPG struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) order_repository.Repository {
	return &orderPG{db: db}
}

func (orderPG *orderPG) ReadOrderById(orderId int) (*entity.Order, errs.Error) {
	row := orderPG.db.QueryRow(getOrderById, orderId)

	var order entity.Order

	err := row.Scan(&order.OrderId, &order.CustomerName, &order.OrderedAt, &order.CreatedAt, &order.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundError("order not found")
		}

		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &order, nil
}

func (orderPG *orderPG) UpdateOrder(orderPayload entity.Order, itemPayload []entity.Item) errs.Error {
	tx, err := orderPG.db.Begin()

	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	_, err = tx.Exec(updateOrderByIdQuery, orderPayload.OrderId, orderPayload.OrderedAt, orderPayload.CustomerName)

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	for _, eachItem := range itemPayload {
		_, err = tx.Exec(updateItemByCodeQuery, eachItem.ItemCode, eachItem.Description, eachItem.Quantity)

		if err != nil {
			tx.Rollback()
			return errs.NewInternalServerError("something went wrong")
		}
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (orderPG *orderPG) ReadOrders() ([]order_repository.OrderItemMapped, errs.Error) {
	rows, err := orderPG.db.Query(getOrdersWithItemsQuery)

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	orderItems := []order_repository.OrderItem{}

	for rows.Next() {
		var orderItem order_repository.OrderItem

		err = rows.Scan(
			&orderItem.Order.OrderId, &orderItem.Order.CustomerName, &orderItem.Order.OrderedAt, &orderItem.Order.CreatedAt, &orderItem.Order.UpdatedAt,
			&orderItem.Item.ItemId, &orderItem.Item.ItemCode, &orderItem.Item.Quantity, &orderItem.Item.Description, &orderItem.Item.OrderId, &orderItem.Item.CreatedAt, &orderItem.Item.UpdatedAt,
		)

		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong")
		}

		orderItems = append(orderItems, orderItem)
	}

	return order_repository.OrderItems(orderItems).HandleMappingOrderWithItems(), nil

}

func (orderPG *orderPG) CreateOrderWithItems(orderPayload entity.Order, itemPayload []entity.Item) errs.Error {

	tx, err := orderPG.db.Begin()

	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	var orderId int

	orderRow := tx.QueryRow(createOrderQuery, orderPayload.OrderedAt, orderPayload.CustomerName)

	err = orderRow.Scan(&orderId)

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	for _, eachItem := range itemPayload {
		_, err := tx.Exec(createItemQuery, eachItem.ItemCode, eachItem.Description, eachItem.Quantity, orderId)

		if err != nil {
			tx.Rollback()
			return errs.NewInternalServerError("something went wrong")
		}
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (orderPG *orderPG) DeleteOrder(orderId int) errs.Error {
	_, err := orderPG.db.Exec(deleteOrderQuery, orderId)
	if err != nil {
		return errs.NewInternalServerError("failed to delete order")
	}
	return nil
}
