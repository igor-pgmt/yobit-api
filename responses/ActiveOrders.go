package responses

type ActiveOrders struct {
	Success uint8                          `json:"success"`
	Return  map[int]map[string]interface{} `json:"return"`
	Error   string                         `json:"error"`
}

func NewActiveOrders() ActiveOrders {
	activeOrders := ActiveOrders{}
	activeOrders.Return = make(map[int]map[string]interface{})
	return activeOrders
}
