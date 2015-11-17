package main
import "fmt"

func main(){
	mySlice := []int{2,3,4}
	fmt.Println(mySlice)

	mySecondSlice := []string{"Good", "Morning"}
	fmt.Println(mySecondSlice)

	var myThirdSlice []int = make([]int, 5, 10)
	myThirdSlice[0] = 1
	myThirdSlice[1] = 2
	myThirdSlice[2] = 3
	myThirdSlice[3] = 4
	myThirdSlice[4] = 5
	fmt.Println(myThirdSlice)
	fmt.Println(len(myThirdSlice))
	fmt.Println(cap(myThirdSlice))

	myNewSlice := []int{3,4,5}
	myOtherSlice := []int{4,5,6}
	fmt.Println(append(myNewSlice, myOtherSlice...))

	myNewSlice = append(myNewSlice[:3], myNewSlice[4:]...)
	fmt.Println(myNewSlice)

	myLastSlice := make([]int, 2, 3)
	myLastSlice[2]=7
}
