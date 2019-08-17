package main

import "fmt"

//
//func countSmaller(nums []int) []int {
//	n := len(nums)
//	count := make([]int, n)
//	arr := make([][]int, 0)
//	for i := 0; i < n; i++ {
//		tmp := make([]int, 2)
//		tmp[0] = nums[i]
//		tmp[1] = i
//		arr = append(arr, tmp)
//	}
//	temp := make([][]int, 0)
//	for i := 0; i < n; i++ {
//		tmp := make([]int, 2)
//		temp = append(temp, tmp)
//	}
//	mergeSort(arr, 0, n-1, temp, &count, nums)
//	return count
//}
//
//func mergeSort(arr [][]int, left, right int, temp [][]int, count *[]int, nums []int) {
//	if left < right {
//		mid := (right + left) / 2
//		mergeSort(arr, left, mid, temp, count, nums)
//		mergeSort(arr, mid+1, right, temp, count, nums)
//		//merge2(arr, left, mid, right, temp, count)
//		//合并两个有序数组
//		i := left
//		j := mid + 1
//		t := 0 //临时数组索引
//		for i <= mid && j <= right {
//			if arr[i][0] <= arr[j][0] { //谁小谁就放在temp数组最前面，直到两个数组其中一个遍历完，此时跳出循环
//				temp[t][0] = arr[i][0]
//				temp[t][1] = arr[i][1]
//				(*count)[arr[i][1]] += j - mid - 1 //这个j-mid-1就代表当前比arr[i][0]元素小的个数，因为下面的else就一直让j++，
//				t++                                //因此j-mid-1就是统计的个数
//				i++
//			} else {
//				temp[t][0] = arr[j][0]
//				temp[t][1] = arr[j][1]
//				t++
//				j++
//			}
//		}
//		//将左边序列元素放入temp
//		for i <= mid {
//			temp[t][0] = arr[i][0]
//			temp[t][1] = arr[i][1]
//			(*count)[arr[i][1]] += j - mid - 1
//			t++
//			i++
//		}
//		//将右边序列元素放入temp
//		for j <= right {
//			temp[t][0] = arr[j][0]
//			temp[t][1] = arr[j][1]
//			t++
//			j++
//		}
//		//temp数组整体拷贝至原数组
//		t = 0
//		for left <= right {
//			arr[left][0] = temp[t][0]
//			arr[left][1] = temp[t][1]
//			left++
//			t++
//		}
//	}
//}
//方法二：二叉排序树,每次插入节点的时候多设置一个变量count,记录左边有几个比他小的
//当插入节点小于当前节点node,当前节点count++
//若插入节点大于等于当前节点,插入节点的count+=node.count+1
type BSTNode struct {
	count int
	val   int
	left  *BSTNode
	right *BSTNode
}

func countSmaller(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}
	res := make([]int, n)
	root := new(BSTNode)
	root.val = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		count := 0
		node := new(BSTNode)
		node.val = nums[i]
		insert(root, node, &count)
		res[i] = count
	}
	return res
}
func insert(root *BSTNode, node *BSTNode, count *int) {
	if node.val <= root.val { //注意这里是<=,因为是求小于当前元素的个数,这个count++不影响跟他相同元素的个数,只有++,才能
		root.count++ //保证后面元素计算是正确的
		if root.left != nil {
			insert(root.left, node, count)
		} else {
			root.left = node
		}
	} else {
		*count += root.count + 1
		if root.right != nil {
			insert(root.right, node, count)
		} else {
			root.right = node
		}
	}
}
func main() {
	nums := []int{2, 0, 1}
	fmt.Println(countSmaller(nums))
}
