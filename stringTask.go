package main 
import (
	"strings"
	"fmt"
	"string"
)
func stringTask() string {
	var inputString string
	var counter int
	counter := 0
	fmt.Scanln(inputString)
	var vowels [5]char 
	vowels := ["a","e","i","o","u"]
	for i := range inputString {
        for j := range vowels {
			if i == j {
				counter++
				inputString := strings.Replace(inputString, j, "", counter)
			}
		} 
	}
	for i := range inputString {
		if i % 2 == 0 {
			s = s[:i] + "." + s[i:]
		}
	}
	return inputString
}