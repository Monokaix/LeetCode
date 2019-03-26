package main

import "math"
//dp from bottom to top
func minimumTotal(triangle [][]int) int {
	for i := len(triangle)-2;i>=0;i--{
		for j :=0;j< len(triangle[i]);j++{
			triangle[i][j] += int(math.Min(float64(triangle[i+1][j]),float64(triangle[i+1][j+1])))
		}
	}
	return triangle[0][0]
}
