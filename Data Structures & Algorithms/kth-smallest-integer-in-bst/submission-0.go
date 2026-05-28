/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    arr:=make([]int,0,1000)
	inorder(&arr,root)
	return arr[k-1]
}
func inorder(arr *[]int, root *TreeNode){
	if root == nil{
		return
	}
	inorder(arr, root.Left)
	*arr = append(*arr,root.Val)
	inorder(arr,root.Right)
}