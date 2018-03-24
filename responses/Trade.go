package responses

type Trade struct {
	Success uint8       `json:"success"`
	Return  TradeReturn `json:"return"`
	Error   string      `json:"error"`
}

type TradeReturn struct {
	Received float64            `json:"received"` // сколько валюты куплено/продано
	Remains  float64            `json:"remains"`  // сколько валюты осталось купить/продать
	OrderID  int                `json:"order_id"` // ID созданного ордера
	Funds    map[string]float64 `json:"funds"`    // балансы, актуальные после запроса
}

func NewTrade() Trade {
	trade := Trade{}
	trade.Return = TradeReturn{}
	trade.Return.Funds = make(map[string]float64)
	return trade
}
