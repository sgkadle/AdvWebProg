package main
import "fmt"

func main(){
	mySentence := "2 4"
	var two, four int
	fmt.Sscan(mySentence, &two, &four)
	fmt.Println(two, four)
	mySentence = "This is a 2nd attempt"
	var myWords []string
	num, err := fmt.Sscan(mySentence, myWords...)
	fmt.Println(num, err)
	fmt.Println(myWords)
}