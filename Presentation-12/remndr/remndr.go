package main
import "fmt"

func main(){
	fmt.Println("Please enter two numbers : ")
	var dividend, divisor int
	fmt.Scan(&dividend)
	fmt.Scan(&divisor)
	fmt.Println("The remainder is : ", dividend & divisor)
}