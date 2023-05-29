/*var previous_total int = 30
current_users := []int{22, 3, 23, 12, 10}
all_records := make([]interface{}, 0)

You have the above declaration. You’re to add the previous total to the first index value in this case becomes 52 and then add the first index value to the next until you get the desire output => [52 55 78 90 100]

Use range and another implementation use i++ for the two scenarios. Best of luck */

package main

import (
	"fmt"
)

func main() {
	// range
	previous_total := 30
	current_users := []int{22, 3, 23, 12, 10}
	all_records := make([]interface{}, 0)

	total := previous_total
	for _, user := range current_users {
		total += user
		all_records = append(all_records, total)
	}
	fmt.Println(all_records, "range")

	// i++

	output := make([]interface{}, 0)
	total = previous_total
	current_users = []int{22, 3, 23, 12, 10}
	for i := 0; i < len(current_users); i++ {
		total += current_users[i]
		output = append(output, total)
	}

	fmt.Println(output, "i++")
}
