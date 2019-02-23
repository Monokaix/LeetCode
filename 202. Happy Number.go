package main

import (
	"fmt"
)

//using map

func isHappy(n int) bool {
	m := make(map[int]bool)
	for {
		if n == 1 {
			return true
		}

		if _, ok := m[n]; ok {
			return false
		} else {
			m[n] = true
		}
		n = split(n)
	}

}

func split(n int) int {
	num := 0
	for n != 0 {
		num += (n % 10) * (n % 10)
		n = n / 10
	}
	return num
}

func main() {
	fmt.Println(isHappy(10))
}
