package responses

type Trades struct {
	Success  int                   `json:"success"`
	PairData map[string]TradeDatas `json:"return"`
	Error    string                `json:"error"`
}

type TradeDatas []TradeData

type TradeData struct {
	Type      string  `json:"type"`      // ask - продажа, bid - покупка
	Price     float64 `json:"price"`     // цена покупки/продажи
	Amount    float64 `json:"amount"`    // количество
	Tid       uint    `json:"tid"`       // идентификатор сделки
	Timestamp int64   `json:"timestamp"` // время сделки
}

// NewTrades returns new structure for Trade response
func NewTrades() Trades {
	trades := Trades{}
	trades.PairData = make(map[string]TradeDatas)
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
