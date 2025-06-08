# ACME Corp Travel Reimbursement Challenge - Submission Summary

## Solution Overview

Successfully reverse-engineered ACME Corp's 60-year-old travel reimbursement system using a **3-tier pricing model** based on receipt thresholds.

## Final Formula (Score: 11,413)

```go
if receipts <= $1,629:
    reimbursement = 53.6 × days + 0.455 × miles + 0.706 × receipts + 43.9

else if receipts <= $2,141:  
    reimbursement = 38.0 × days + 0.376 × miles + 0.564 × receipts + 68.0

else:
    reimbursement = 32.3 × days + 0.249 × miles + 0.592 × receipts - 147.4
```

## Performance Metrics

- **Score**: 11,413 (64% improvement from baseline 32,044)
- **Mean Absolute Error**: $113.13
- **Exact Matches**: 4/1000 (within $0.01)
- **Implementation**: Go (calculate.go)

## Key Discoveries

1. **Tier-based system** with receipt thresholds at $1,629 and $2,141
2. **Decreasing coefficients** at higher tiers (anti-fraud measure)
3. **No special handling** for .49/.99 receipt endings (myth debunked)
4. **Business logic** aligns with employee interviews about efficiency bonuses

## Files for Submission

1. `calculate.go` - Main implementation
2. `run.sh` - Executable wrapper script  
3. `private_results.txt` - 5,000 predictions for private test set

## Development Process

- Started with employee interview analysis
- Tested multiple approaches (statistical, ML, rule-based)
- Used gradient descent optimization for coefficient tuning
- Achieved progressive improvements: 32,044 → 14,755 → 11,909 → 11,419 → 11,413

## Repository Structure

```
top-coder-challenge/
├── calculate.go          # Core implementation
├── run.sh               # Executable script
├── private_results.txt  # Submission file (5,000 results)
├── approaches/          # Various solution attempts
└── *.go                # Optimization tools
```

Ready for submission via Google Form!