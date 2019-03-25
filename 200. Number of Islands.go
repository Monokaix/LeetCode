package main

import "fmt"

var d2 = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func numIslands(grid [][]byte) int {
	var m = len(grid)
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res++
				dfs3(grid, i, j)
			}
		}
	}
	return res
}

func dfs3(grid [][]byte, x int, y int) {
	grid[x][y] = '2' //标记为2代表已经被访问过
	for i := 0; i < 4; i++ {
		newX := x + d2[i][0]
		newY := y + d2[i][1]
		if newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid[0]) && grid[newX][newY] == '1' {
			dfs3(grid, newX, newY)
		}
	}
	return
}
func main(){
	grid := [][]byte{{'1','1','0','0','0'},{'1','1','0','0','0'},{'0','0','1','0','0'},{'0','0','0','1','1'}}
	fmt.Println(numIslands(grid))
}