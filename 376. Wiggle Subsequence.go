package main

//动态规划，但是可以优化空间复杂度，只使用一个变量即可
//一升一降，所以是up时，down+1，down时，up+1
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	up := 1
	down := 1
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			up = down + 1
		} else if nums[i] > nums[i-1] {
			down = up + 1
		}
	}
	if up > down {
		return up
	}
	return down
}

//贪心
func wiggleMaxLength2(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	//定义三种状态
	BEGIN := 0
	UP := 1
	DOWN := 2
	STATE := BEGIN
	cnt := 1
	for i := 1; i < n; i++ {
		switch  STATE{
		case BEGIN:
			if nums[i]>nums[i-1]{
				cnt++
				STATE = UP
			}
			if nums[i] < nums[i-1]{
				cnt++
				STATE = DOWN
			}
		case UP:
			if nums[i] < nums[i-1]{ // 之前是up状态，现在遇到一个down序列，说明需要++
				STATE = DOWN        // 若还是递增序列，则不用管他，即不做任何操作
				cnt++
			}
		case DOWN :
			if nums[i] > nums[i-1]{ // 之前是down状态，现在遇到一个up序列，说明需要++
				STATE = UP
				cnt++
			}
		}
	}
	return cnt
}
