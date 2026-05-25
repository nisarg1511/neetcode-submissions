/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	left:=count(root.Left,root)
	right:=count(root.Right,root)
	return left+right+1
}

func count (root *TreeNode,largest *TreeNode) int {
	if root == nil{
		return 0
	}
	total:=0
	if root.Val >= largest.Val{
		total++
		largest = root
	}
	left:=count(root.Left,largest)
	right:=count(root.Right,largest)
	return left+right+total
}