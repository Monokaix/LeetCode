package main

var d3 = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
//思路
//1、搜索四周，dfs遍历将四周及四周相连的O改为临时遍历Y
//2、还原，将所有O换为X，将所有Y换为O
func solve(board [][]byte) {
	if board == nil || len(board)<=0{
		return
	}
	for i := 0;i<len(board);i++{ //最左最右两列
		dfs4(board,i,0)
		dfs4(board,i,len(board[0])-1)
	}
	for i := 0;i<len(board[0]);i++{ //最上最下两行
		dfs4(board,0,i)
		dfs4(board,len(board)-1,i)
	}
	for i:=0;i<len(board);i++{ //还原
		for j := 0;j<len(board[0]);j++{ //两个还原顺序不能颠倒
			if board[i][j] == 'O'{
				board[i][j] = 'X'
			}
			if board[i][j] == 'Y'{
				board[i][j] = 'O'
			}
		}
	}
}

func dfs4(board [][]byte, x int, y int) {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || board[x][y] != 'O' {
		return
	}
	board[x][y] = 'Y'
	for i := 0; i < 4; i++ {
		dfs4(board, x+d3[i][0],y+d3[i][1])
	}
}
