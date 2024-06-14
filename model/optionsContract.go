package model
// OptionsContract represents the data structure of an options contract
type OptionsContract struct {
	Type           string  `json:"type"`          // "call" or "put"
	StrikePrice    float64 `json:"strike_price"`  // Strike price of the option
	Bid            float64 `json:"bid"`           // Bid price of the option
	Ask            float64 `json:"ask"`           // Ask price of the option
	ExpirationDate string  `json:"expiration_date"` // Expiration date of the option
	LongShort      string  `json:"long_short"`    // "long" or "short" position
}

// AnalysisResponse represents the data structure of the analysis result
type AnalysisResponse struct {
	XYValues        []XYValue `json:"xy_values"`
	MaxProfit       float64   `json:"max_profit"`
	MaxLoss         float64   `json:"max_loss"`
	BreakEvenPoints []float64 `json:"break_even_points"`
}

// XYValue represents a pair of X and Y values
type XYValue struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
// Your model here