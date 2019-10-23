package main

//从左下角开始遍历,每次缩减一行或一列
func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix) - 1
	col := 0
	for row >= 0 && col < len(matrix[0]) {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			col++
		} else {
			row--
		}
	}
	return false
}
