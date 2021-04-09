package main

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	row0 := false
	col0 := false
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			col0 = true
			break
		}
	}

	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			row0 = true
			break
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 检查一行和第一列，有0的话重新置0避免漏掉
	if row0 {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	if col0 {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
