package main

////方法一:额外数组标记
//func spiralOrder(matrix [][]int) []int {
//	x := 0
//	y := 0
//	m := len(matrix)
//	if m == 0 {
//		return nil
//	}
//	n := len(matrix[0])
//	mark := make([][]int, 0) //标记元素是否被访问过
//	res := make([]int, 0)
//	for i := 0; i < m; i++ {
//		tmp := make([]int, n)
//		mark = append(mark, tmp)
//	}
//	i := 1
//	mark[0][0] = 1
//	res = append(res, matrix[0][0])
//	for i < m*n {
//		for y+1 < n && mark[x][y+1] == 0 { //访问最上边一行
//			y++
//			i++
//			res = append(res, matrix[x][y])
//			mark[x][y] = 1
//		}
//		for x+1 < m && mark[x+1][y] == 0 { //访问最右边一行
//			x++
//			i++
//			res = append(res, matrix[x][y])
//			mark[x][y] = 1
//		}
//
//		for y-1 >= 0 && mark[x][y-1] == 0 { //访问最下边一行
//			y--
//			i++
//			res = append(res, matrix[x][y])
//			mark[x][y] = 1
//		}
//		for x-1 >= 0 && mark[x-1][y] == 0 { //访问最左边一行
//			x--
//			i++
//			res = append(res, matrix[x][y])
//			mark[x][y] = 1
//		}
//	}
//	return res
//}

//方法二:不标记,直接循环,用四个值分别标记每次遍历的范围
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}
	n := len(matrix[0])
	up := 0        //最上面一行范围
	right := n - 1 //最右边一行
	down := m - 1  //最下边一行
	left := 0      //最左边一行
	res := make([]int, 0)
	for {
		for i := left; i <= right; i++ { //从左到右遍历第一行
			res = append(res, matrix[up][i])
		}
		up++ //每做完一轮,up递增
		if up > down {
			break
		}
		for i := up; i <= down; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		if right < left {
			break
		}
		for i := right; i >= left; i-- {
			res = append(res, matrix[down][i])
		}
		down--
		if down < up {
			break
		}
		for i := down; i >= up; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return res
}
