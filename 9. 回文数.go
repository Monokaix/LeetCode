package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	sum := 0
	tmp := x
	for x != 0 {
		sum = sum*10 + x%10
		x = x / 10
	}
	return tmp == sum
}

