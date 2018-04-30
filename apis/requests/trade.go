package requests

type TradeSettings struct {
	Pair   string  `json:"pair"`   // pair (example: ltc_btc)
	Type   string  `json:"type"`   // transaction type (example: buy or sell)
	Rate   float64 `json:"rate"`   // exchange rate for buying or selling (value: numeral)
	Amount float64 `json:"amount"` // amount needed for buying or selling (value: numeral)
}
