package main
import(
	"fmt"
	"math"
)

func main(){
	var x float64
	fmt.Println ("Enter a number of float type1 : ")
	fmt.Scanln (&x)
	fmt.Println(math.Ceil(x))
}