package responses

type CancelOrder struct {
	Success uint8  `json:"success"`
	Return  COData `json:"return"`
	Error   string `json:"error"`
}

type COData struct {
	OrderID int                    `json:"order_id"` // order ID
	Funds   map[string]interface{} `json:"funds"`    // balances active after request
}

func NewCancelOrder() CancelOrder {
	cancelOrder := CancelOrder{}
	cancelOrder.Return.Funds = make(map[string]interface{})
	return cancelOrder
}
