package main

func totalNQueens(n int) int {
	res := 0
	//每一列是否被占用
	col := make([]bool, n)
	//左对角线是否有元素
	dia1 := make([]bool, 2*n-1) //因为有2*n-1条对角线
	//右对角线是否有元素
	dia2 := make([]bool, 2*n-1) //因为有2*n-1条对角线
	var putQueen func(int, int)
	putQueen = func(index int, n int) { //index代表当前处理的行，n是问题规模，row代表第
		if index == n { //找到一个解
			res++
		}
		for i := 0; i < n; i++ { //遍历每一列寻找不冲突的位置
			if !col[i] && !dia1[index+i] && !dia2[index-i+n-1] {
				col[i] = true
				dia1[index+i] = true
				dia2[index-i+n-1] = true
				putQueen(index+1, n)
				col[i] = false
				dia1[index+i] = false
				dia2[index-i+n-1] = false //不用关心输出的每一个解是什么，仅计数即可
			}
		}
	}
	putQueen(0, n)
	return res
}
