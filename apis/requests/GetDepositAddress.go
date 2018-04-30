package requests

type GetDepositAddressSettings struct {
	CoinName string `json:"coin_name"` // ticker (example: BTC)
	NeedNew  uint64 `json:"need_new"`  // value: 0 or 1, on default: 0
}
