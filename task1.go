/* You have the below list, write a go application to print the index 
of the cars together with their names. And the food print the name of index 3 in the same application.

cars:= []string{"Honda", "Toyota", "Ford", "Chevrolet", "BMW"}

food := []string{"Pizza", "Tacos", "Sushi", "Burgers", "Curry"} */



package main

import ("fmt")

func main() {
	cars := [5]string{"Honda", "Toyota", "Ford", "Chevrolet", "BMW"}

	food := [5]string{"Pizza", "Tacos", "Sushi", "Burgers", "Curry"}
	for i, a := range cars{
		fmt.Println("Index:", i, a)
	}

	fmt.Println("Index of 3 in the food list is:",food[3])
}
