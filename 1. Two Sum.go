package main

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	res := make([]int,2)
	for i:=0;i<len(nums);i++{
		complement := target-nums[i]
		if _,ok :=m[complement];ok{
			res[0] = m[complement]
			res[1] = i
			return res
		}
		m[nums[i]] = i
	}
	return res
}
