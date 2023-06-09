// single case switch

package main
import ("fmt")

func main() {
	day := 1

	switch day {
	case 1: 
	fmt.Println("Monday")

	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	}


// multi-case switch

  day = 6

switch day {
case 1,3,5:
	fmt.Println("Odd weekday")
case 2,6:
	fmt.Println("Even weekday")
case 4,7:
	fmt.Println("Weekend")
default:
	fmt.Println("Invald day of day number")
} }