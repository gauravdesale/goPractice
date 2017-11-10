package main
import "fmt"
import "os"

func calculateWatermelon() string {
	y = os.Args // just getting the command line arguements
	x := y
	result := "NO"
	fmt.Scan(&x)
	//here is the real function now
	//first check if the value of x in in between 1 and 100 
	//then check if the number is even
	//then loop through the first half of the number and check if there exists a number if the second half
	if x > 0 && x < 100 {
		if x % 2 == 0 {
			for i := 0; i < x / 2; i++ {
				for j := x / 2; j < x; j++ {
					if (i == j && (i+j == x)) {
						result := "YES"
					}
				}
			}
		}
	}
	return result
}