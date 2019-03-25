package main

var inPacific bool
var inAtlantic bool
var d4 = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func pacificAtlantic(matrix [][]int) [][]int {
	res := [][]int{}
	m := len(matrix)
	if m == 0 {
		return res
	}
	n := len(matrix[0])
	visited := [][]bool{}
	for i := 0;i<m;i++{
		tmp := make([]bool,n)
		visited = append(visited, tmp)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			inPacific = false
			inAtlantic = false
			dfs5(matrix, i, j, visited)
			if inPacific && inAtlantic {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

//深度优先判断是否能同时流入太平洋和大西洋
func dfs5(matrix [][]int, x int, y int, visited [][]bool) {
	if x == 0 || y == 0 {
		inPacific = true //能流进太平洋
	}
	if x == len(matrix)-1 || y == len(matrix[0])-1 {
		inAtlantic = true // 能流进大西洋
	}
	if inPacific && inAtlantic {
		return
	}
	visited[x][y] = true
	for i := 0; i < 4; i++ {
		newX := x + d4[i][0]
		newY := y + d4[i][1]
		if newX >= 0 && newX < len(matrix) && newY >= 0 && newY < len(matrix[0]) && !visited[newX][newY] && matrix[newX][newY] <= matrix[x][y] {
			dfs5(matrix, newX, newY, visited)
		}
	}
	visited[x][y] = false
}
