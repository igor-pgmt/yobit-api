package responses

type WithdrawCoinsToAddress struct {
	Success uint8      `json:"success"`
	Return  WCTAReturn `json:"return"`
	Error   string     `json:"error"`
}

type WCTAReturn struct {
	ServerTime uint64 `json:"server_time"`
}

func NewWithdrawCoinsToAddress() WithdrawCoinsToAddress {
	withdrawCoinsToAddress := WithdrawCoinsToAddress{}
	withdrawCoinsToAddress.Return = WCTAReturn{}
	return withdrawCoinsToAddress
}
