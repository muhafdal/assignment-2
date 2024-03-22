package entity

import "time"

type Order struct {
	OrderId      uint
	CustomerName string
	OrderedAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
