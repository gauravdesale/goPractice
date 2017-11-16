package main
import "fmt"
func piling() int {
	var x int
	var y int
	var counter1 int
	var counter2 int
	var result int
	fmt.Scan(x)
	fmt.Scan(y)
	if x >= 0 &&n x <= 16 {
		if x % 2  {
			counter1++
		}
		if y >= 0 && y <= 16 {
			counter2++
		}
	}
	result := counter1 * counter2
	return result
}