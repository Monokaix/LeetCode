package main

func rotate(matrix [][]int) {
	length := len(matrix)
	// 矩阵是对称的，i，j范围只需到一半就行，把矩阵分成了四部分
	// matrix new[j][n-i-1] = matrix old[i][j]
	for i := 0; i < length/2; i++ {
		for j := 0; j < (length+1)/2; j++ {
			tmp := matrix[i][j]
			m := length - i - 1
			n := length - j - 1
			// 这里依次轮转替换
			matrix[i][j] = matrix[n][i]
			matrix[n][i] = matrix[m][n]
			matrix[m][n] = matrix[j][m]
			matrix[j][m] = tmp
		}
	}
}
