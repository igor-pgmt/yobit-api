package responses

type Ticker struct {
	Success  int `json:"success"`
	PairData map[string]TData
	Error    string `json:"error"`
}

type TData struct {
	High    float64 `json:"high"`    // maximal price
	Low     float64 `json:"low"`     // minimal price
	Avg     float64 `json:"avg"`     // average price
	Vol     float64 `json:"vol"`     // traded volume
	VolCur  float64 `json:"vol_cur"` // traded volume in currency
	Last    float64 `json:"last"`    // last transaction price
	Buy     float64 `json:"buy"`     // buying price
	Sell    float64 `json:"sell"`    // selling price
	Updated float64 `json:"int"`     //  last cache upgrade
}

func NewTicker() Ticker {
	ticker := Ticker{}
	ticker.PairData = make(map[string]TData)
	return ticker
}
