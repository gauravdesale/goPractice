package main
import (
	"fmt"
	"strconv"
)


func countLetters(lol ...string) string {
	var letters = "abcdefghijklmnopqrstuvwxyz"
	var counter = 0
	var firstLetter = lol[0]
	var lastLetter = lol[len(lol)-1]
	for _, c := range lol {
		counter++
	}
	var resultCounter = strconv.Itoa(counter)
	var result = firstLetter + resultCounter + lastLetter
	fmt.Printf(result)
	return result	
}

func main() {
	countLetters("pneumonoultramicroscopicsilicovolcanoconiosis")
}