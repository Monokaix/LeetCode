package main

func climbStairs(n int) int {
	if n < 2 {
		return 1
	}
	var fibN int
	var fibOne int = 1
	var fibTwo int = 1
	for i := 2; i <= n; i++ {
		fibN = fibOne + fibTwo
		fibOne = fibTwo
		fibTwo = fibN
	}
	return fibN
}