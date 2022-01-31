package main

import "fmt"

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
func main() {
	fmt.Println(factorial(5))
	fmt.Println(fib(5))
}

/*
fib(0) = 0 , fib(1) = 1 , fib(n) = fib(n-1) + fib(n-2) .
Write a recursive function which can find
fib(n) .
*/
func fib(x uint) uint {
	if x == 1 {
		return 1
	}
	return x + fib(x-1)
}
