package main

import "fmt"

func moveZeroes(nums []int)  {
	index := 0
	for _,v := range nums{
		if v!= 0{
			nums[index] = v
			index++
		}
	}
	for index != len(nums){
		nums[index] = 0
		index++
	}
}
func main() {
	nums := []int{0,1,2,0,5,0,6}
	moveZeroes(nums)
	fmt.Println(nums)
}
