package main

func minSubArrayLen(s int, nums []int) int {
	l := 0
	r := -1
	sum := 0
	res := len(nums)+1
	for l < len(nums){ //交替变换两个指针，找出最优解，即滑动窗口
		if r+1 < len(nums) && sum < s{
			r++
			sum += nums[r]
		}else{
			sum -= nums[l] //减掉最左边的一个元素
			l++
		}

		if(sum >= s){
			if res > r-l+1{
				res = r-l+1
			}
		}
	}

	if res == len(nums)+1{
		return 0 //没有找到，返回0
	}
	return res
}

func main() {
	
}
