package entity

import (
	"time"
)

type Item struct {
	ItemId      uint
	ItemCode    string
	Quantity    int
	Description string
	OrderId     uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
