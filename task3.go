// Naming a baby with the if function statement //

/*package main
import ("fmt")

func main() {
	boy := 12
	girl := 8
	if boy < girl {
		fmt.Println("Baby will be named Davida.")
		} else if boy <= girl {
			fmt.Println("Baby will be named Davey.")
			} else {
				fmt.Println("Baby will be named David.")

}
}  */



package main
import ("fmt")

func main() {
	boy := []string{"Emmanuel", "David"}
	for i , boy := range boy {
		if i < 0 {
			fmt.Println("Baby will be named", boy)
			} else if i >= 1 {
				fmt.Println("Baby will be named ", boy)
	}  } }
	
