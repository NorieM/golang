package main

import (
	"errors"
	"fmt"
)

func main() {

	revenue, err := getUserInput("Revenue: ")

	if err != nil {
		panic("Invalid value entered for revenue!")
	}

	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		panic("Invalid value entered for expenses!")
	}

	taxRate, err := getUserInput("Tax Rate:")

	if err != nil {
		panic("Invalid value entered for tax rate!")
	}

	ebt, profit, ratio := calculateReturn(revenue, expenses, taxRate)

	prettyPrint(ebt, "%.1f\n")
	prettyPrint(profit, "%.1f\n")
	prettyPrint(ratio, "%.3f\n")
}

func getUserInput(prompt string) (float64, error) {
	var inputValue float64
	fmt.Print(prompt)
	fmt.Scan(&inputValue)

	if inputValue <=0 {
		return 0, errors.New("invalid value entered")
	}

	return inputValue, nil
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