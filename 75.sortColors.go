package main

import "fmt"

func sortColors(nums []int)  {
	zero := -1
	two :=len(nums)
	for i:=0;i<two;{
		if nums[i] == 1{
			i++
		}else if nums[i] == 2{
			two--
			nums[i],nums[two] = nums[two],nums[i]
		}else{
			zero++
			nums[zero],nums[i] = nums[i],nums[zero] //直接交换即可，因为nums[zero]一定是1
			i++
		}
	}
}
//使用快排思路，最左边全是0，中间是1，最右边是2，
//遇到0或者2就和当前位置交换，遇到1则继续往前走
func main() {
	nums := []int{2,1,0,1,2,0}
	sortColors(nums)
	fmt.Println(nums)
}
