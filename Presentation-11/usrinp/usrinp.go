package main
import "fmt"

func main(){
	var usr string
	fmt.Print ("Please enter your name here : ")
	fmt.Scan (&usr)
	fmt.Println ("Name :" + usr)
}