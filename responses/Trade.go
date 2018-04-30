package responses

type Trade struct {
	Success uint8       `json:"success"`
	Return  TradeReturn `json:"return"`
	Error   string      `json:"error"`
}

type TradeReturn struct {
	Received float64            `json:"received"` // amount of currency bought / sold
	Remains  float64            `json:"remains"`  // amount of currency to buy / to sell
	OrderID  int                `json:"order_id"` // created order ID
	Funds    map[string]float64 `json:"funds"`    // funds active after request
}

func NewTrade() Trade {
	trade := Trade{}
	trade.Return = TradeReturn{}
	trade.Return.Funds = make(map[string]float64)
	return trade
}
