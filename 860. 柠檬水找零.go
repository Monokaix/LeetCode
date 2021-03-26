package main

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			five += 1
		} else if bills[i] == 10 {
			if five < 1 {
				return false
			}
			five--
			ten++
		} else {
			if five >= 1 && ten >= 1 {
				five--
				ten--
			} else if five >= 3 {
				five = five - 3
			} else {
				return false
			}
		}
	}
	return true
}
