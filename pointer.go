package main

import "fmt"


func pointerPractice() {
	x := 12
	y := &x

	fmt.Println("the value of x is: ", x)
	fmt.Println("the adderess of x is: ", &x)
	fmt.Println("the value of y is: ", y)
	*y = 100//changing the adderess of x
	fmt.Println("the value of x is ", x)
	fmt.Println("the adderess of x is: ", &x)
}
