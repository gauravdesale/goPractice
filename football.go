package main 
import "fmt"
func safeorNot(someThing string) bool {
	var resultFinal bool
	resultFinal = false
	var finalLength int
	finalLength := len(someThing)
	var somethingElse [finalLength]int
	somethingElse := [finalLength]int(someThing)

	for i, elem := range somethingElse {
		var temp int
		temp = elem
		for j := 0; j < 7; j++ {
			if temp == somethingElse[i+j] {
				resultFinal = true
			}
		}
	}
	return resultFinal
}

func main() {
	var final bool
	final = safeorNot("00100110111111101")
	fmt.Print(final)
}