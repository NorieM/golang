package main

import (
	"fmt"
)

func main() {

	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate:")

	ebt, profit, ratio := calculateReturn(revenue, expenses, taxRate)

	prettyPrint(ebt, "%.1f\n")
	prettyPrint(profit, "%.1f\n")
	prettyPrint(ratio, "%.3f\n")
}

func getUserInput(prompt string) float64 {
	var inputValue float64
	fmt.Print(prompt)
	fmt.Scan(&inputValue)
	return inputValue
}

func calculateReturn(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func prettyPrint(value float64, format string) {
	fmt.Printf(format, value)
}