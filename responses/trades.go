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
	Tid       int     `json:"tid"`       // идентификатор сделки
	Timestamp int     `json:"timestamp"` // время сделки
}

func NewTrades() Trades {
	trades := Trades{}
	trades.PairData = make(map[string]TradeDatas)
	return trades
}
