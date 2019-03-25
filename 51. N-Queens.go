package main

func solveNQueens(n int) [][]string {
	res := [][]string{}
	//每一列是否被占用
	col := make([]bool, n)
	//左对角线是否有元素
	dia1 := make([]bool, 2*n-1) //因为有2*n-1条对角线
	//右对角线是否有元素
	dia2 := make([]bool, 2*n-1) //因为有2*n-1条对角线
	row := []int{}
	var putQueen func(int, int, []int)
	putQueen = func(index int, n int, row []int) { //index代表当前处理的行，n是问题规模，row代表第
		if index == n { //找到一个解
			res = append(res, generateBoard(n, row))
		}
		for i := 0; i < n; i++ { //遍历每一列寻找不冲突的位置
			if !col[i] && !dia1[index+i] && !dia2[index-i+n-1] {
				row = append(row, i)
				col[i] = true
				dia1[index+i] = true
				dia2[index-i+n-1] = true
				putQueen(index+1, n, row)
				col[i] = false
				dia1[index+i] = false
				dia2[index-i+n-1] = false
				row = row[:len(row)-1]
			}
		}
	}
	putQueen(0, n, row)
	return res
}

func generateBoard(n int, row []int) []string {
	res := []string{}
	for _, v := range row {
		tmp := []byte{}
		for i := 0; i < n; i++ {
			tmp = append(tmp, '.')
		}
		tmp[v] = 'Q'
		res = append(res, string(tmp))
	}
	return res
}
