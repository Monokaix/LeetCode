package main

func permute(nums []int) [][]int {
	res := [][]int{}
	temp := []int{}
	used := make([]bool, len(nums))
	generatePermute(nums, 0, temp, &res, used)
	return res

}

func generatePermute(nums []int, index int, temp []int, res *[][]int, used []bool) {
	if index == len(nums) {
		tmp := make([]int,len(temp))
		copy(tmp,temp)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !used[i] {
			temp = append(temp, nums[i])
			used[i] = true
			generatePermute(nums, index+1, temp, res, used)
			temp = temp[:len(temp)-1]
			used[i] = false
		}
	}
}
func main() {

}
