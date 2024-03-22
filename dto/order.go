package dto

import "time"

type NewOrderRequestDto struct {
	OrderedAt    time.Time           `json:"orderedAt" example:"2023-07-10T21:21:46+00:00"`
	CustomerName string              `json:"customerName" example:"John Doe"`
	Items        []NewItemRequestDto `json:"items"`
}

type NewOrderResponseDto struct {
	BaseResponse
}

type GetOrdersResponseDto struct {
	BaseResponse
	Data []OrderWithItems `json:"data"`
} // @name GetOrdersResponse

type OrderWithItems struct {
	OrderId      int                  `json:"orderId"`
	CustomerName string               `json:"customerName"`
	OrderedAt    time.Time            `json:"orderedAt"`
	CreatedAt    time.Time            `json:"createdAt"`
	UpdatedAt    time.Time            `json:"updatedAt"`
	Items        []GetItemResponseDto `json:"items"`
}
