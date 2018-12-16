package main

func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers)-1
	res := make([]int,2,2)
	for l < r{
		if numbers[l]+numbers[r] == target{
			res[0] = l+1
			res[1] = r+1
			return res
		}else if numbers[l]+numbers[r] < target{
			l++
		}else{
			r--
		}
	}
	return res
}

func main() {

}
