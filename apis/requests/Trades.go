package requests

type TradesSettings struct {
	Pair  string `json:"pair"`  // pair (example: ltc_btc)
	Limit uint64 `json:"limit"` // limit stipulates size of withdrawal (on default 150 to 2000 maximum)
}
