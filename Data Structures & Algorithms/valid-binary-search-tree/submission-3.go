/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	arr:=traverse(root)
	// smallest :=arr[0]
	for i:=1;i<len(arr);i++{
		if arr[i]<=arr[i-1]{
			return false
		} 
	}
	return true
}

func traverse(root *TreeNode) []int{
	if root ==  nil {
		return []int{}
	}
	left:=traverse(root.Left)
	left = append(left,root.Val)
	right:=traverse(root.Right)
	left = append(left,right...)
	return  left
}