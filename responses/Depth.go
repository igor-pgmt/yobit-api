package responses

type Depth struct {
	PairData map[string]PData
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
