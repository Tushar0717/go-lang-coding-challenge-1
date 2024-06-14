package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// OptionsContract structure for the request body
type OptionsContract struct {
	Type          string  `json:"type" binding:"required,oneof=call put"`           // "call" or "put"
	StrikePrice   float64 `json:"strike_price" binding:"required"`   // Strike price of the option
	Bid           float64 `json:"bid" binding:"required"`            // Bid price of the option
	Ask           float64 `json:"ask" binding:"required"`            // Ask price of the option
	ExpirationDate string  `json:"expiration_date" binding:"required"` // Expiration date of the option
	LongShort     string  `json:"long_short" binding:"required,oneof=long short"`     // "long" or "short"// Your code here
}

// AnalysisResult structure for the response body
type AnalysisResult struct {
	GraphData       []GraphPoint `json:"graph_data"`
	MaxProfit       float64      `json:"max_profit"`
	MaxLoss         float64      `json:"max_loss"`
	BreakEvenPoints []float64    `json:"break_even_points"`
}

// GraphPoint structure for X & Y values of the risk & reward graph
type GraphPoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/analyze", func(c *gin.Context) {
		var contracts []OptionsContract

		if err := c.ShouldBindJSON(&contracts); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := performAnalysis(contracts)// Your code here

		c.JSON(http.StatusOK, result)
	})

	return router
}
// performAnalysis performs the risk and reward analysis on the given options contracts
func performAnalysis(contracts []OptionsContract) AnalysisResult {
	var maxProfit, maxLoss float64
	var breakEvenPoints []float64
	graphData := make([]GraphPoint, 0)

	for price := 0.0; price <= 200.0; price += 1.0 {
		totalProfitLoss := 0.0

		for _, contract := range contracts {
			profitLoss := calculateProfitLoss(contract, price)
			totalProfitLoss += profitLoss
		}

		if totalProfitLoss > maxProfit {
			maxProfit = totalProfitLoss
		}
		if totalProfitLoss < maxLoss {
			maxLoss = totalProfitLoss
		}

		if math.Abs(totalProfitLoss) < 1e-6 {
			breakEvenPoints = append(breakEvenPoints, price)
		}

		graphData = append(graphData, GraphPoint{X: price, Y: totalProfitLoss})
	}

	breakEvenPoints = removeDuplicates(breakEvenPoints)

	return AnalysisResult{
		GraphData:       graphData,
		MaxProfit:       maxProfit,
		MaxLoss:         maxLoss,
		BreakEvenPoints: breakEvenPoints,
	}
}

// calculateProfitLoss calculates the profit or loss for a given options contract and underlying price
func calculateProfitLoss(contract OptionsContract, price float64) float64 {
	var profitLoss float64

	switch contract.Type {
	case "call":
		if contract.LongShort == "long" {
			profitLoss = math.Max(0, price-contract.StrikePrice) - contract.Ask
		} else {
			profitLoss = contract.Bid - math.Max(0, price-contract.StrikePrice)
		}
	case "put":
		if contract.LongShort == "long" {
			profitLoss = math.Max(0, contract.StrikePrice-price) - contract.Ask
		} else {
			profitLoss = contract.Bid - math.Max(0, contract.StrikePrice-price)
		}
	}

	return profitLoss
}

// removeDuplicates removes duplicate float64 values from a slice and returns a sorted slice of unique values
func removeDuplicates(values []float64) []float64 {
	uniqueValues := make(map[float64]bool)
	var result []float64
	for _, value := range values {
		if _, exists := uniqueValues[value]; !exists {
			uniqueValues[value] = true
			result = append(result, value)
		}
	}
	sort.Float64s(result)
	return result
}