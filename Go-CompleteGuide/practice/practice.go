package main

import "fmt"

// Time to practice what you learned!

type Product struct {
	title string
	id string
	price float64
}

func NewProduct(Title string, Id string, Price float64) (Product) {
	return Product{
		title: Title,
		id: Id,
		price: Price,
	}
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
	fmt.Println(hobbies[0:2])

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	slice_1 := hobbies[0:2]
	fmt.Println(slice_1)

	slice_2 := hobbies[:2]
	fmt.Println(slice_2)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	slice_3 := slice_1[1:3]
	fmt.Println(slice_3)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	course_goals:= []string{"Finish the course", "Finish some other courses"}

	fmt.Println(course_goals)

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	course_goals[1] = "A change of goal"
	course_goals = append(course_goals, "New goal")

	fmt.Println(course_goals)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	products:= []Product{NewProduct("Product 1", "01", 10.99), NewProduct("Product 2", "02", 8.99)}

	fmt.Println(products)

	products = append(products, NewProduct("Product 3", "03", 9.99))
	
	fmt.Println(products)

}	

