package main

//记忆化搜索
//memo[i][c] 使用索引为[0...i]的元素是否可以填充容量为c的背包
func canPartition(nums []int) bool {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	memo := make([][]int, 0)
	for i := 0; i < n; i++ {
		tmp := make([]int, sum/2+1)
		for i := 0; i < sum/2+1; i++ {
			tmp[i] = -1
		}
		memo = append(memo, tmp)
	}
	return tryPartition(nums, sum/2, n-1, memo)
}

//memo[i] -1表示未访问过，0表示false，1表示true
func tryPartition(nums []int, sum int, index int, memo [][]int) bool {
	if sum == 0 { //刚好可以填满
		return true
	}
	if index < 0 || sum < 0 {
		return false
	}
	if memo[index][sum] != -1 {
		return memo[index][sum] == 1
	}

	if tryPartition(nums, sum, index-1, memo) || tryPartition(nums, sum-nums[index], index-1, memo) {
		memo[index][sum] = 1
	} else {
		memo[index][sum] = 0
	}
	return memo[index][sum] == 1
}

//方法二：dp，还是和背包问题类似
func canPartition2(nums []int) bool {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	C := sum/2 + 1
	memo := make([]bool, C)
	for j := 0; j <= C; j++ {
		memo[j] = nums[0] == j
	}
	for i := 1; i < n; i++ {
		for j := 0; j < C; j++ {
			memo[j] = memo[j-1] || memo[j-nums[i]]
		}
	}
	return memo[C]
}
