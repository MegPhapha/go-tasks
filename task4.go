package main
import ("fmt")

func main() {
	scores := []int{100, 50, 60}
	for _, marks := range scores { 
marks += 25

scores = append(scores, marks, 85)

fmt.Println(scores)

} }





