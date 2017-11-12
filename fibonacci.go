package main

func fib(x int) int {
	if x<=3{
		return 1
	}else{
		return fib(x-1) + fib(x-2)
	}
}


