package param

import "time"

type (
	OrderResp struct {
		OrderId      int        `json:"order_id"`
		OrderAt      time.Time  `json:"order_at"`
		CustomerName string     `json:"customer_name"`
		Items        []ItemResp `json:"items"`
	}

	ItemResp struct {
		Id          int    `json:"id"`
		ItemCode    string `json:"item_code"`
		Description string `json:"description"`
		Quantity    int    `json:"quantity"`
	}
)
