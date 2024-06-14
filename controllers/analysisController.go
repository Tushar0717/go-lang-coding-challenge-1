package controllers

import (
	"net/http"
)

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

func AnalysisHandler(w http.ResponseWriter, r *http.Request) {	var contracts []OptionsContract

	// Decode the JSON request body into contracts slice
	if err := json.NewDecoder(r.Body).Decode(&contracts); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Perform the analysis
	xyValues := calculateXYValues(contracts)
	maxProfit := calculateMaxProfit(contracts)
	maxLoss := calculateMaxLoss(contracts)
	breakEvenPoints := calculateBreakEvenPoints(contracts)

	// Create the response
	response := AnalysisResponse{
		XYValues:        xyValues,
		MaxProfit:       maxProfit,
		MaxLoss:         maxLoss,
		BreakEvenPoints: breakEvenPoints,
	}

	// Encode and send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)// Your code here
	// Your code here
}

func calculateXYValues(contracts []OptionsContract) []XYValue {
	for _, contract := range contracts {
		// Simplified calculation for illustration
		x := contract.StrikePrice
		y := contract.Bid - contract.Ask // Simplified profit/loss calculation
		xyValues = append(xyValues, XYValue{X: x, Y: y})
	}
	return xyValues// Your code here
	return nil
}

func calculateMaxProfit(contracts []OptionsContract) float64 {
	var maxProfit float64
	for _, contract := range contracts {
		if contract.LongShort == "long" {
			maxProfit += contract.Bid
		} else if contract.LongShort == "short" {
			maxProfit -= contract.Ask
		}
	}
	return maxProfit// Your code here
	return 0
}

func calculateMaxLoss(contracts []OptionsContract) float64 {
	var maxLoss float64
	for _, contract := range contracts {
		if contract.LongShort == "long" {
			maxLoss -= contract.Ask
		} else if contract.LongShort == "short" {
			maxLoss += contract.Bid
		}
	}
	return maxLoss// Your code here
	return 0
}

func calculateBreakEvenPoints(contracts []OptionsContract) []float64 {
	var breakEvenPoints []float64
	for _, contract := range contracts {
		breakEvenPoint := contract.StrikePrice
		if contract.LongShort == "long" {
			breakEvenPoint += contract.Bid
		} else if contract.LongShort == "short" {
			breakEvenPoint -= contract.Ask
		}
		breakEvenPoints = append(breakEvenPoints, breakEvenPoint)
	}
	return breakEvenPoints// Your code here
	return nil
}
