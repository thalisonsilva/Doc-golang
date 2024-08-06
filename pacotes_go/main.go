package main

import (
	"fmt"
	"example.com/myapp/myapp/mathutil"
)
func main (){
	sum := mathutil.Add (3, 4)
	difference := mathutil.Subtract(10, 6)

	fmt.Println("Sum: ", sum)
	fmt.Println("Difference: ", difference)
}