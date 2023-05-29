package main
import ("fmt")

func main() {
  previous_total := 30
  current_users := []int{22, 3, 23, 12, 10}
  all_records := make([]interface{}, 0)

  total := previous_total
  for _, user := range current_users {
	total += user
	all_records = append(all_records, total)
  }

  fmt.Println("Using range;", all_records)

  output := make ([]interface{}, 0)
  total = previous_total
  current_users = []int{22, 3, 23, 12, 10}
  new_current_users := current_users
  new_current_users[2] = 80


  //using i++
for i := 0
 i < len(new_current_users)
 i++ {
	total += new_current_users[i]
	output = append(output, total)
 }

 fmt.Println("Using i++;", output)

}




 