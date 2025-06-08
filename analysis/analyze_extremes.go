package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"sort"
)

type TestCase struct {
	Input struct {
		TripDurationDays    int     `json:"trip_duration_days"`
		MilesTraveled       float64 `json:"miles_traveled"`
		TotalReceiptsAmount float64 `json:"total_receipts_amount"`
	} `json:"input"`
	ExpectedOutput float64 `json:"expected_output"`
}

type ErrorCase struct {
	Case           TestCase
	Predicted      float64
	Error          float64
	ReceiptsPerDay float64
	Efficiency     float64
}

func calculateCurrentPrediction(days int, miles, receipts float64) float64 {
	// Current best coefficients
	base := 55*float64(days) + 0.42*miles + 0.52*receipts + 250
	
	efficiency := miles / float64(days)
	receiptsPerDay := receipts / float64(days)
	
	if efficiency >= 150 && efficiency <= 250 {
		base *= 1.06
	} else if efficiency < 80 {
		base *= 0.97
	} else if efficiency > 400 {
		base *= 0.98
	}
	
	if days >= 4 && days <= 6 {
		base *= 1.05
	} else if days >= 10 {
		base *= 0.95
	}
	
	if receiptsPerDay > 150 {
		base *= 0.94
	} else if receipts < 30 && days > 1 {
		base *= 0.97
	}
	
	if receipts > 1800 && receiptsPerDay > 200 {
		base *= 0.7
	}
	
	return math.Round(base*100) / 100
}

func main() {
	data, err := ioutil.ReadFile("public_cases.json")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	
	var testCases []TestCase
	err = json.Unmarshal(data, &testCases)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}
	
	var errorCases []ErrorCase
	
	// Calculate errors for all cases
	for _, testCase := range testCases {
		days := testCase.Input.TripDurationDays
		miles := testCase.Input.MilesTraveled
		receipts := testCase.Input.TotalReceiptsAmount
		expected := testCase.ExpectedOutput
		
		predicted := calculateCurrentPrediction(days, miles, receipts)
		error := math.Abs(predicted - expected)
		
		errorCases = append(errorCases, ErrorCase{
			Case:           testCase,
			Predicted:      predicted,
			Error:          error,
			ReceiptsPerDay: receipts / float64(days),
			Efficiency:     miles / float64(days),
		})
	}
	
	// Sort by error (highest first)
	sort.Slice(errorCases, func(i, j int) bool {
		return errorCases[i].Error > errorCases[j].Error
	})
	
	fmt.Println("=== TOP 20 HIGHEST ERROR CASES ===")
	for i := 0; i < 20 && i < len(errorCases); i++ {
		ec := errorCases[i]
		fmt.Printf("Case %d: %d days, %.0f miles, $%.2f receipts\n", 
			i+1, ec.Case.Input.TripDurationDays, ec.Case.Input.MilesTraveled, ec.Case.Input.TotalReceiptsAmount)
		fmt.Printf("  Expected: $%.2f, Predicted: $%.2f, Error: $%.2f\n", 
			ec.Case.ExpectedOutput, ec.Predicted, ec.Error)
		fmt.Printf("  Receipts/day: $%.2f, Efficiency: %.1f mi/day\n\n", 
			ec.ReceiptsPerDay, ec.Efficiency)
	}
	
	// Analyze patterns in high-error cases
	fmt.Println("=== HIGH-ERROR PATTERN ANALYSIS ===")
	
	// Cases with very high receipts per day
	highSpendingCases := 0
	for _, ec := range errorCases[:50] { // Top 50 error cases
		if ec.ReceiptsPerDay > 200 {
			highSpendingCases++
		}
	}
	fmt.Printf("High spending cases (>$200/day) in top 50 errors: %d\n", highSpendingCases)
	
	// Cases with low expected output despite high inputs
	luxuryPenaltyCases := 0
	for _, ec := range errorCases[:50] {
		if ec.Case.Input.TotalReceiptsAmount > 1000 && ec.Case.ExpectedOutput < 1000 {
			luxuryPenaltyCases++
		}
	}
	fmt.Printf("Luxury penalty cases (>$1000 receipts, <$1000 output) in top 50: %d\n", luxuryPenaltyCases)
	
	// Cases with very low expected vs predicted ratio
	extremePenaltyCases := 0
	for _, ec := range errorCases[:50] {
		ratio := ec.Case.ExpectedOutput / ec.Predicted
		if ratio < 0.5 {
			extremePenaltyCases++
		}
	}
	fmt.Printf("Extreme penalty cases (expected < 50%% of predicted) in top 50: %d\n", extremePenaltyCases)
	
	// Suggest penalty thresholds
	fmt.Println("\n=== SUGGESTED PENALTY IMPROVEMENTS ===")
	fmt.Println("1. Receipts > $1500 AND receipts/day > $180: Apply 0.5x multiplier")
	fmt.Println("2. Receipts > $2000: Apply 0.4x multiplier") 
	fmt.Println("3. Receipts/day > $300: Apply 0.3x multiplier")
	fmt.Println("4. Long trips (8+ days) with high spending: Additional 0.8x multiplier")
}