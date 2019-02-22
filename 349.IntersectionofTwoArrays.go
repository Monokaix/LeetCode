package main

func intersection(nums1 []int, nums2 []int) []int {
	map1 := make(map[int]int)
	map2 := make(map[int]int)
	res := []int{}
	for _,v := range nums1{
		map1[v] = 0
	}

	for _,v := range nums2{
		if _,ok := map1[v];ok{
			map2[v] = 0
		}
	}
	for key,_:=range map2{
		res = append(res,key)
	}
	return res
}
