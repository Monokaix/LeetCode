package main

//func binaryTreePaths(root *TreeNode) []string {
//	res := make([]string,0)
//	if root == nil{
//		return res
//	}
//	if root.Left == nil && root.Right == nil{
//		res = append(res, strconv.Itoa(root.Val))
//		return res
//	}
//	leftRes := binaryTreePaths(root.Left)
//	for i:= 0;i<len(leftRes);i++{
//		res = append(res, strconv.Itoa(root.Val)+"->"+leftRes[i])
//	}
//	rightRes := binaryTreePaths(root.Right)
//	for i:= 0;i<len(rightRes);i++{
//		res = append(res, strconv.Itoa(root.Val)+"->"+rightRes[i])
//	}
//
//	return res
//}
func binaryTreePaths(root *TreeNode) []string {

}