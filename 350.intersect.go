package main

func intersect(nums1 []int, nums2 []int) []int {
	mymap := make(map[int]int)
	res := []int{}
	for _,v := range nums1{
		if _,ok := mymap[v]; ok{
			mymap[v]++
		}else{
			mymap[v] = 1
		}
	}

	for _,v := range nums2{
		if value,ok := mymap[v]; ok && value>0{
			res = append(res,v)
			mymap[v]--
		}
	}

	return res
}

func main() {
	
}
