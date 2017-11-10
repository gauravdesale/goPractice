package main
import (
	"fmt"
)

func theatreCalculate() int {
	//input those three variables into functions
	//run two for loops and go through those
	//check if the product of the two is equal to or divisible by 4 and if so then add to a counter variale
	//return this counter variable at the end
	fmt.Print("Enter all dimensions: ")
	var xBig int
	var yBig int
	var squareSwag int
	fmt.Scan(&xBig)
	fmt.Scan(&yBig)
	fmt.Scan(&squareSwag)
	var result int
	result = 0
	for i:=0; i < xBig; i++ {
		for j:=0; j < yBig; j++ {
			if (i+j) % squareSwag == 0 {
				result++
			}
		}
	}
	return result
}
