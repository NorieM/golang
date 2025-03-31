package main

import "fmt"

func makeSliceOfInt(min, max int) []int {
	a := make([]int, max-min+1)
	for i:= range a {
		a[i] = min + i
	}

	return a
}

func main() {
	vals := makeSliceOfInt(0,200)

	for _, val := range vals {
		if val%2  == 0 {
			fmt.Printf("%d is even", val)
		} else {
			fmt.Printf("%d is odd", val)
		}
		fmt.Println()
	}
}