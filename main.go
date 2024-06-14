package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/analyze", analyzeHandler)

	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func analyzeHandler(w http.ResponseWriter, r *http.Request) {
	var contracts []OptionsContract
	err := json.NewDecoder(r.Body).Decode(&contracts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Perform the analysis logic
	// For simplicity, let's assume we calculate the total bid and ask prices
	totalBid := 0.0
	totalAsk := 0.0
	for _, contract := range contracts {
		totalBid += contract.Bid
		totalAsk += contract.Ask
	}

	// Create an AnalysisResponse with the total bid and ask prices
	response := AnalysisResponse{
		TotalBid: totalBid,
		TotalAsk: totalAsk,
	}

	// Encode the response as JSON and write it to the response writer
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// OptionsContract represents the data structure of an options contract
type OptionsContract struct {
	Type           string  `json:"type"`           // "Call" or "Put"
	StrikePrice    float64 `json:"strike_price"`
	Bid            float64 `json:"bid"`
	Ask            float64 `json:"ask"`
	ExpirationDate string  `json:"expiration_date"`
	LongShort      string  `json:"long_short"` // "long" or "short"
}

// AnalysisResponse represents the data structure of the analysis result
type AnalysisResponse struct {
	TotalBid float64 `json:"total_bid"`
	TotalAsk float64 `json:"total_ask"`// Your code here
}