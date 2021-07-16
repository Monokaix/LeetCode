package main

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	visited := make([][]bool, 0)
	for i := 0; i < m; i++ {
		tmp := make([]bool, n)
		visited = append(visited, tmp)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				if dfs(board, word, i, j, m, n, 0, visited) {
					return true
				}
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, x, y, m, n, index int, visited [][]bool) bool {
	if index == len(word) {
		return true
	}
	if x < 0 || x == m || y < 0 || y == n || board[x][y] != word[index] || visited[x][y] {
		return false
	}

	visited[x][y] = true

	res := dfs(board, word, x, y-1, m, n, index+1, visited) ||
		dfs(board, word, x, y+1, m, n, index+1, visited) ||
		dfs(board, word, x-1, y, m, n, index+1, visited) ||
		dfs(board, word, x+1, y, m, n, index+1, visited)

	visited[x][y] = false
	return res
}
