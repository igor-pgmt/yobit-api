package requests

type TradeSettings struct {
	Pair   string  // пара (пример: ltc_btc)
	Type   string  // тип операции (пример: buy или sell)
	Rate   float64 // курс, по которому необходимо купить или продать (значение: числовое)
	Amount float64 // количество, которое необходимо купить или продать (значение: числовое)
}
