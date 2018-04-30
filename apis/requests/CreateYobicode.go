package requests

type CreateYobicodeSettings struct {
	Currency string  `json:"coin_name"` // ticker (example: BTC)
	Amount   float64 `json:"amount"`    // amount to withdraw
}
