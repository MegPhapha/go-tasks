// the continue statement
package main
import ("fmt")

func main() {
	for i := 0; i < 5; i++ {
		if i == 3 { 
		continue
	}
fmt.Println("contine statement example;",i)
}

// the break statement
for a := 0; a < 5; a++ {
	if a == 4 { 
	break
}
fmt.Println("break statement example;",a)
}

//nested loops

adj := []string{"big", "tasty"}
fruits := []string{"orange","banana","apple"}
for i := 0; i < len(adj); i++ {
	for j := 0; j < len(fruits); j++ {
		fmt.Println(adj[i], fruits[j])
	}
}
}

