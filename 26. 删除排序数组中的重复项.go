package main

//i指针一直往前移动,j指针指向当前应该被覆盖的位置
//i指向的值和前一个值一样,则i继续滑动
//遇到和前一个值不一样的时候,i覆盖掉j
func removeDuplicates(nums []int) int {
	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[j] = nums[i] //直接替换
			j++               //同时j向前移动
		}
	}
	return j
}
