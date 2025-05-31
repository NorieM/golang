package main

import "fmt"

func main() {
	// numbers := []int{1, 2, 3, 4, 5}

	result := sumup(20, 1,10, 15)

	fmt.Println(result)
}

func sumup(startingValue int, numbers ...int) int {
	sum := startingValue

	for _, val := range numbers {
		sum += val
	}
	return sum
}