package main
import "fmt"

type student struct{
	name string
	age int
}

func main(){
	sai := student{name: "Sai", age:21}
	vaish := student{name:"Vaishnavi"}

	fmt.Println(sai.name, sai.age)
	fmt.Println(vaish.name)

	vaish.age=19
	fmt.Println(vaish.age)
}