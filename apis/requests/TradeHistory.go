package requests

type TradeHistorySettings struct {
	From   uint64 `json:"from"`    // No. of transaction from which withdrawal starts (value: numeral, on default: 0)
	Count  uint64 `json:"count"`   // quantity of withrawal transactions (value: numeral, on default: 1000)
	FromID uint64 `json:"from_id"` // ID of transaction from which withdrawal starts (value: numeral, on default: 0)
	EndID  uint64 `json:"end_id"`  // ID of transaction at which withdrawal finishes (value: numeral, on default: ∞)
	Order  string `json:"order"`   // sorting at withdrawal (value: ASC or DESC, on default: DESC)
	Since  uint64 `json:"since"`   // the time to start the display (value: unix time, on default: 0)
	End    uint64 `json:"end"`     // the time to end the display (value: unix time, on default: ∞)
	Pair   string `json:"pair"`    // pair (example: ltc_btc)
}
