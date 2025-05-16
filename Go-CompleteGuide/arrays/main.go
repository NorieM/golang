package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output(){

	fmt.Println(m)
}

func main() {

	userNames := make([]string, 2, 5)

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Norie")

	fmt.Println(userNames)

	coursesRatings :=make(floatMap, 3)

	coursesRatings["go"] =  4.7
	coursesRatings["react"] =  4.8
	coursesRatings["angular"] =  4.8
	
	coursesRatings.output()

	for index,value := range(userNames){
		fmt.Println("Index: ", index)
		fmt.Println("Value: ", value)
	}

		for key,value := range(coursesRatings){
		fmt.Println("Key: ", key)
		fmt.Println("Value: ", value)
	}

}