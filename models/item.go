package models

type (
	OrderItem struct {
		Id          int
		ItemCode    string
		Description string
		Quantity    int
		OrderId     int
	}
)
