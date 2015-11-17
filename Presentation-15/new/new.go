package main
import "fmt"

func main(){
	var x *int = new (int)
	fmt.Println(x)
	fmt.Println(*x)

	var (y) + string = new(string)
	fmt.Println(y)
	fmt.Println(*y)

	var z *bool = new(bool)
	fmt.Println(z)
	fmt.Println(*z)

	slice := make([]int, 5, 5)
	fmt.Println(slice)

	myMap := make(map[int]string)
	fmt.Println(myMap)

}