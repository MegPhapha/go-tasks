// else if statement

package main
import ("fmt")

func main () {
	time := 20
	if time <= 10 {
		fmt.Println("Good morning")
	} else if time < 20 {
		fmt.Println("Good day")
	} else {
		fmt.Println("Good evening")
	}

// another example of the else if statement

a := 14
b := 14

if a < b {
	fmt.Println("a is less than b")
} else if a > b {
	fmt.Println("a is greater than b")
} else {
	fmt.Println("a is equal to b")
}
 

 // nested if statements

 
	num := 20
	if num >= 10 {
		fmt.Println("Num is more than 10.")
		if num > 15 {
			fmt.Println("Num is also more than 15.")
			 }
		} else {
			fmt.Println("Num is less than 10.")
		}
	}  
 