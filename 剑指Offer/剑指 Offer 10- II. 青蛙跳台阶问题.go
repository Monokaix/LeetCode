package main

import "fmt"

func numWays(n int) int {
	if n < 2 {
		return 1
	}
	fib1 := 1
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
	fmt.Println(numWays(1))
}
