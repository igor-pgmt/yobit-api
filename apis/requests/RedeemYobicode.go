package requests

type RedeemYobicodeSettings struct {
	Coupon string `json:"coupon"` // yobicode to redeem (example: YOBITUZ0HHSTB...OQX3H01BTC)
}
