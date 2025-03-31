package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Please enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Please enter the years: ")
	fmt.Scan(&years)

	fmt.Print("Please enter the return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future value (adjusted for inflation): %.2f\n", futureRealValue)

	outputText(formattedFV)
	outputText(formattedRFV)

}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64,float64) {
	fv :=investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	rfv:= fv / math.Pow(1+inflationRate/100, years)

	return fv, rfv
}