package main

func merge(nums1 []int, m int, nums2 []int, n int)  {
	i := m-1
	j := n-1
	t := m+n-1
	for i>= 0 && j >= 0{
		if nums1[i] >= nums2[j]{ //谁大谁就放在nums1数组最后面，直到两个数组其中一个遍历完，此时跳出循环
			nums1[t] = nums1[i]
			i--
			t--
		}else{
			nums1[t] = nums2[j]
			j--
			t--
		}
	}
	//将nums1序列元素放入nums1
	for i >= 0{
		nums1[t] = nums1[i]
		t--
		i--
	}
	//将nums2序列元素放入nums1
	for j >= 0{
		nums1[t] = nums2[j]
		t--
		j--
	}
}

func main() {
	
}
