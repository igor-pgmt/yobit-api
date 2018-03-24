package responses

type OrderInfo struct {
	Success uint8                             `json:"success"`
	Return  map[string]map[string]interface{} `json:"return"`
}

func NewOrderInfo() OrderInfo {
	orderInfo := OrderInfo{}
	orderInfo.Return = make(map[string]map[string]interface{})
	return orderInfo
}
