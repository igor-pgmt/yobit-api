package responses

type OrderInfo struct {
	Success uint8                          `json:"success"`
	Return  map[int]map[string]interface{} `json:"return"`
	Error   string                         `json:"error"`
}

func NewOrderInfo() OrderInfo {
	orderInfo := OrderInfo{}
	orderInfo.Return = make(map[int]map[string]interface{})
	return orderInfo
}
