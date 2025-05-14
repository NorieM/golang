package lists

import "fmt"


func main() {
	prices:= []float64{10.99, 8.99}

	fmt.Println(prices[1])

	prices = append(prices, 5.99)

	fmt.Println(prices)
}

// func main() {
// 	var productNames [4]string = [4]string{"A book"}
// 	prices := []float64{1.34, 334.3, 123.0, 20.0}
// 	fmt.Println(prices)
// 	fmt.Println(productNames)

// 	featuredPrices := prices[1:3]

// 	fmt.Println(featuredPrices)
// }

