package main

import "fmt"

// Time to practice what you learned!

type Product struct {
	title string
	id    string
	price float64
}

func main() {

	// 1) Create a new array (!) that contains three hobbies you have
	hobbies := [3]string{"Swimming", "Cycling", "Reading"}

	// 		Output (print) that array in the command line.
	fmt.Println(hobbies)

	// 2) Also output more data about that array:

	//		- The first element (standalone)
	fmt.Println(hobbies[0])

	//		- The second and third element combined as a new list
	fmt.Println(hobbies[1:3])

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	mainHobbies := hobbies[0:2]
	fmt.Println(mainHobbies)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	slice_3 := mainHobbies[1:3]
	fmt.Println(slice_3)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	courseGoals := []string{"Finish the course", "Finish some other courses"}

	fmt.Println(courseGoals)

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	courseGoals[1] = "A change of goal"
	courseGoals = append(courseGoals, "New goal")

	fmt.Println(courseGoals)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	products := []Product{{"Product 1", "01", 10.99}, {"Product 2", "02", 8.99}}
	
	fmt.Println(products)
	
	//		Then add a third product to the existing list of products.
	newProduct := Product{"Product 3", "03", 9.99}

	products = append(products, newProduct)

	fmt.Println(products)

}
