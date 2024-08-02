package main
import "fmt"

func myAppend(sl []int, val int) []int{
	sl = append(sl, val)
	printSlice(sl)
	return sl
}

func printSlice(sl []int){
	fmt.Printf("Slice: %v || ", sl)
	fmt.Printf("Len: %v, Cap: %v", len(sl), cap(sl))
	fmt.Println()
}

func main() {
	var mySlice = make([]int, 1)
	printSlice(mySlice)
	for i:=1; i<=5; i++ {
		mySlice = myAppend(mySlice,i)
	}
}