package responses

type Trades struct {
	Success  int                    `json:"success"`
	PairData map[string][]TradeData `json:"return"`
	Error    string                 `json:"error"`
}

type TradeData struct {
	Type      string  `json:"type"`      // ask - sell, bid - buy
	Price     float64 `json:"price"`     // buying / selling price
	Amount    float64 `json:"amount"`    // amount
	Tid       uint    `json:"tid"`       // transaction id
	Timestamp int64   `json:"timestamp"` // transaction timestamp
}

// NewTrades returns new structure for Trade response
func NewTrades() Trades {
	trades := Trades{}
	trades.PairData = make(map[string][]TradeData)
	return trades
}

// Separate separates asks and bids from the main result
func Separate(t Trades, pair string) ([]TradeData, []TradeData) {
	var asks, bids []TradeData
	for i := 0; i < len(t.PairData[pair]); i++ {
		switch t.PairData[pair][i].Type {
		case "ask":
			asks = append(asks, t.PairData[pair][i])
		case "bid":
			bids = append(bids, t.PairData[pair][i])
		}
	}
	return asks, bids
}

// GetPriceBefore returns first action price before specified timestamp
func GetPriceBefore(tds []TradeData, before int64) (price float64) {
	var currentTimeVal int64
	for _, val := range tds {
		if val.Timestamp < before && val.Timestamp > currentTimeVal {
			currentTimeVal = val.Timestamp
			price = val.Price
		}
	}
	return
}
