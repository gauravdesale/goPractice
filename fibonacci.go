package main
import (
	"fmt"
)
func fib(x int) int {
	if x<=3{
		return 1
	}else{
		return fib(x-1) + fib(x-2)
	}
}

func main() {
	x := fib(5)
	fmt.Printf("The fifth fibonacci number is: %d\n", x)
}
