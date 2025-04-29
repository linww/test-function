package main

import (
	"fmt"
)

// CalculateInvestment calculates the future value of a monthly investment
// over a given number of years with a fixed annual interest rate.
func CalculateInvestment(monthlyInvestment float64, annualInterestRate float64, years int) (float64, float64) {
	months := years * 12
	monthlyInterestRate := annualInterestRate / 12 / 100
	totalValue := 0.0
	totalInvestment := monthlyInvestment * float64(months)

	for i := 0; i < months; i++ {
		totalValue = (totalValue + monthlyInvestment) * (1 + monthlyInterestRate)
	}

	return totalInvestment, totalValue
}

func main() {
	monthlyInvestment := 1500.0 // 每月定投金额
	annualInterestRate := 7.0   // 年化收益率
	years := 30                 // 投资年限

	totalInvestment, futureValue := CalculateInvestment(monthlyInvestment, annualInterestRate, years)
	fmt.Printf("%d年后总成本: %.2f, %d年后总价值: %.2f\n", years, totalInvestment, years, futureValue)
}
