package responses

type CreateYobicode struct {
	Success uint8    `json:"success"`
	Return  CYReturn `json:"return"`
	Error   string   `json:"error"`
}

type CYReturn struct {
	Coupon  string             `json:"coupon"`  // Yobicode
	TransID uint8              `json:"transID"` // always 1 for compatibility with api of other exchanges
	Funds   map[string]float64 `json:"funds"`   // balances active after request
}

func NewCreateYobicode() CreateYobicode {
	createYobicode := CreateYobicode{}
	createYobicode.Return = CYReturn{}
	return createYobicode
}
