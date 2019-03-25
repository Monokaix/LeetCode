package main

import "fmt"

var d = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func exist(board [][]byte, word string) bool {
	var m = len(board)
	//var visited = make(map[int]bool)
	//var visited = [1000][1000]bool{}
	for i := 0; i < m; i++ {
		for j := 0; j < len(board[i]); j++ {
			if searchWord(board, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func searchWord(board [][]byte, word string, index int, startx int, starty int) bool {
	if index == len(word)-1 {
		return board[startx][starty] == word[index]
	}
	if board[startx][starty] == word[index] {
		temp := board[startx][starty]
		board[startx][starty] = 0
		for i := 0; i < 4; i++ { //从四个方向寻找
			newx := startx + d[i][0]
			newy := starty + d[i][1]
			if newx >= 0 && newx < len(board) && newy >= 0 && newy < len(board[0]) {
				fmt.Println(board)
				fmt.Println("######")
				if searchWord(board, word, index+1, newx, newy) {
					//fmt.Println("innner",board)
					return true
				}
			}
		}
		board[startx][starty] = temp
	}
	return false
}
func main(){
	board :=[][]byte{{'A','B','C','E'},{'S','F','E','S'},{'A','D','E','E'}}
	fmt.Println(exist(board,"ABCESEEEFS"))
}