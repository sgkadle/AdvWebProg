package main
import "fmt"

func main(){
	myMap := map[string]string{
		"Sai":"Kadle",
		"Steve":"Jobs",
		"Tony":"Stark",
	}

	myMap["HSHK"] = "BlahBlah"
	myMap["KHSH"] = "halBhalB"
	delete(myMap, "KHSH")
	for key, val := range myMap{
		fmt.Println(key, " - ", val)
	}

	fmt.Println(len(myMap))
	if val, ok := myMap["Tony"]; ok{
		fmt.Println(val)
	}
}