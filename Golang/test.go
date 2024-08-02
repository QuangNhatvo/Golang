package main
import "fmt"

func main(){
	defer fmt.Printf("Word")
	fmt.Printf("Hello ")
}