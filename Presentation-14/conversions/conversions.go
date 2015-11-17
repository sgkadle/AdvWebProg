package main
import "fmt"

func main(){
	//int to float64
	myInt := 5
	fmt.Println(int64(myInt))

	//float64 to int
	myFloat := 9.99
	fmt.Println(int(myFloat))

	//[]byte to string
	myBytes := []byte{'m','o','r','n','i','n','g'}
	fmt.Println(string(myBytes))

	//string to [byte]
	myString := "good night"
	fmt.Println([]byte(myString))
}