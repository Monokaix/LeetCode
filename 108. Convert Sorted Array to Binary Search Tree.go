package main

func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil {
		return nil
	}
	return create(nums, 0, len(nums)-1)
}
func create(nums []int, i int, j int) *TreeNode {
	if j < i{
		return nil
	}
	root := &TreeNode{}
	mid := (i + j) / 2
	root.Val = nums[mid]
	root.Left = create(nums, i, mid-1)
	root.Right = create(nums, mid+1, j)
	return root
}
