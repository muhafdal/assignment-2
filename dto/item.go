package dto

import "time"

type NewItemRequestDto struct {
	ItemCode    string `json:"itemCode" example:"889"`
	Description string `json:"description" example:"BMW"`
	Quantity    int    `json:"quantity" example:"13"`
}

type GetItemResponseDto struct {
	ItemId      int       `json:"itemId"`
	ItemCode    string    `json:"itemCode"`
	Quantity    int       `json:"quantity"`
	Description string    `json:"description"`
	OrderId     int       `json:"orderId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
