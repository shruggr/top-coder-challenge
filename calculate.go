package main

import (
	"fmt"
	"os"
	"strconv"
)

func calculateReimbursement(days int, miles, receipts float64) float64 {
	// Handle zero trips
	if days == 0 && miles == 0 && receipts == 0 {
		return 0
	}
	
	// Optimized 3-tier formula from fine-tuning
	// Score: 11,413 (improvement from 11,419)
	if receipts <= 1629 {
		return 53.6*float64(days) + 0.455*miles + 0.706*receipts + 43.9
	} else if receipts <= 2141 {
		return 38.0*float64(days) + 0.376*miles + 0.564*receipts + 68.0
	} else {
		return 32.3*float64(days) + 0.249*miles + 0.592*receipts - 147.4
	}
}

func main() {
	if len(os.Args) != 4 {
		fmt.Printf("Usage: go run calculate.go <days> <miles> <receipts>")
		os.Exit(1)
	}
	
	days, _ := strconv.Atoi(os.Args[1])
	miles, _ := strconv.ParseFloat(os.Args[2], 64)
	receipts, _ := strconv.ParseFloat(os.Args[3], 64)
	
	result := calculateReimbursement(days, miles, receipts)
	fmt.Printf("%.2f", result)
}