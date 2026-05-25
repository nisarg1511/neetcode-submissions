/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right ==  nil{
		return [] int{root.Val}
	}
	arr:=make([]int,0)
	arr  = append(arr,root.Val)
	right:=rightSideView(root.Right)
	left:=rightSideView(root.Left)
	
	if len(right)>len(left){
		arr = append(arr,right...)
		return arr
	}
	for i :=range len(right){
		left[i] = right[i]
	}
	
	arr = append(arr,left...)
	return arr
}
