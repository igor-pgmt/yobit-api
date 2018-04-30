package responses

type GetInfo struct {
	Success uint8      `json:"success"`
	Return  InfoReturn `json:"return"`
	Error   string     `json:"error"`
}

type InfoReturn struct {
	Funds            map[string]float64 `json:"funds"`             // available account balance (does not include money on open orders)
	FundsInclOrders  map[string]float64 `json:"funds_incl_orders"` // available account balance (include money on open orders)
	Rights           InfoReturnRights   `json:"rights"`            // priviledges of key. withdraw is not used (reserved)
	TransactionCount int64              `json:"transaction_count"` // always 0 (outdated)
	OpenOrders       int64              `json:"open_orders"`       // always 0 (outdated)
	ServerTime       uint64             `json:"server_time"`       // server time
}

type InfoReturnRights struct {
	Info     int64 `json:"info"`
	Trade    int64 `json:"trade"`
	Deposit  int64 `json:"deposit"`
	Withdraw int64 `json:"withdraw"`
}

func NewBalance() GetInfo {
	info := GetInfo{}
	info.Return = InfoReturn{}
	info.Return.Funds = make(map[string]float64)
	info.Return.FundsInclOrders = make(map[string]float64)
	return info
}
