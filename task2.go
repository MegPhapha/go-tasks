package main
import ("fmt")

func main() {
	age := []int{10,20,30}
	names := []string{"Kwame", "Fafs", "Adwoa"}
for i, name:= range names{
	age:=age[i]
	fmt.Println(name, "", "is", age, "","years", "", "old")

}
}

	//for i:=0; i < len(age); i++ {
		//for j:=0; j < len(names); j++ {
		//	fmt.Println(names[j],age[i])
			
//	if i > 0;{
	//	fmt.Println(names[i],age[])