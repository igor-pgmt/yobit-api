package responses

type TradeHistory struct {
	Success uint8                       `json:"success"`
	Return  map[int]map[string]THReturn `json:"return"`
	Error   string                      `json:"error"`
}

type THReturn struct {
	Pair        string  `json:"pair"`          // pair
	Type        string  `json:"type"`          // transaction type
	Amount      float64 `json:"amount"`        // amount
	Rate        float64 `json:"rate"`          // price of buying or selling
	OrderID     int     `json:"order_id"`      // order ID
	IsYourOrder bool    `json:"is_your_order"` // is the order yours
	Timestamp   int     `json:"timestamp"`     // transaction time
}

func NewTradeHistory() TradeHistory {
	tradeHistory := TradeHistory{}
	tradeHistory.Return = make(map[int]map[string]THReturn)
	return tradeHistory
}
