package responses

type GetDepositAddress struct {
	Success uint8     `json:"success"`
	Return  GDAReturn `json:"return"`
	Error   string    `json:"error"`
}

type GDAReturn struct {
	Address         string  `json:"address"`
	ProcessedAmount float64 `json:"processed_amount"`
	ServerTime      uint64  `json:"server_time"`
}

func NewGetDepositAddress() GetDepositAddress {
	getDepositAddress := GetDepositAddress{}
	getDepositAddress.Return = GDAReturn{}
	return getDepositAddress
}
