package recursion

import "fmt"

func main() {
	answer := factorial(5)

	fmt.Println(answer)

}

func factorial(number int) int {

	if number == 0 {
		return 1
	}

	return number * factorial(number-1)

	// result := 1

	// for i := 1; i <= number; i++ {
	// 	result = result * i
	// }

	// return result

	
	
}