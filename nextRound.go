package main
import "fmt"
func nextround() int {
	var numOfPlayers int
	var comparablePlayer int
	var counter int
	counter = 0
	fmt.Scan(numOfPlayers)
	fmt.Scan(comparablePlayer)
	var a [numOfPlayers]int
	for i := 0; i < numOfPlayers; i++ {
		fmt.Scan(a[i])
	}
	for i := 0; i < numOfPlayers; i++ {
		if (a[i] > a[comparablePlayer]) {
			counter++
		}
	}
	return counter
}