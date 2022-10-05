package models

import "time"

type (
	Order struct {
		Id           int
		CustomerName string
		Item         []OrderItem
		OrderAt      time.Time
	}
)
