package main
import "fmt"

func main() {
	a:= []int{2,3,5,7,11}
	fmt.Println("Slice a before:")
	fmt.Println(a)
	b:= a[:3]
	b[0] = 1
	fmt.Println("Slice a after:")
	fmt.Println(a)
	fmt.Println(b)
}