package responses

type Depth struct {
	Success  int `json:"success"`
	PairData map[string]PData
	Error    string `json:"error"`
}

type PData struct {
	Asks [][2]float64 `json:"asks"` // ордера на продажу
	Bids [][2]float64 `json:"bids"` // ордера на покупку
}

func NewDepth() Depth {
	depth := Depth{}
	depth.PairData = make(map[string]PData)
	return depth
}
