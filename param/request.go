package param

import "time"

type (
	OrderReq struct {
		OrderAt      time.Time `json:"order_at"`
		CustomerName string    `json:"customer_name"`
		Items        []ItemReq `json:"items"`
	}

	ItemReq struct {
		Id          int    `json:"id"`
		ItemCode    string `json:"item_code"`
		Description string `json:"description"`
		Quantity    int    `json:"quantity"`
	}
)
