package main

func convert(s string, numRows int) string {
	str := []byte(s)
	res := ""

	if numRows == 1 {
		return s
	}
	cycle := 2*numRows - 2
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < len(str); j += cycle {
			res += string(str[j+i])
			if i != 0 && i != numRows-1 && j+cycle-i < len(str) {
				res += string(str[j+cycle-i])
			}
		}
	}

	return res
}
