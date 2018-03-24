package responses

type GetInfo struct {
	Success uint8      `json:"success"`
	Return  InfoReturn `json:"return"`
	Error   string     `json:"error"`
}

type InfoReturn struct {
	Rights           InfoReturnRight    `json:"rights"`
	Funds            map[string]float64 `json:"funds"`
	FundsInclOrders  map[string]float64 `json:"funds_incl_orders"`
	TransactionCount int64              `json:"transaction_count"`
	OpenOrders       int64              `json:"open_orders"`
	ServerTime       uint64             `json:"server_time"`
}

type InfoReturnRight struct {
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
