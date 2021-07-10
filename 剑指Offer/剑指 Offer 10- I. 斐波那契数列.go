package main

import "fmt"

func fib(n int) int {
	if n < 2 {
		return n
	}
	fib1 := 0
	fib2 := 1
	fibN := 0
	for i := 2; i <= n; i++ {
		fibN = (fib1 + fib2) % 1000000007
		fib1 = fib2
		fib2 = fibN
	}
	return fibN
}
func main() {
	fmt.Println(fib(3))
}
