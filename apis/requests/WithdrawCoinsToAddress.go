package requests

type WithdrawCoinsToAddressSettings struct {
	CoinName string  `json:"coin_name"` // ticker (example: BTC)
	Amount   float64 `json:"amount"`    // amount to withdraw
	Address  string  `json:"address"`   // destination address
}
