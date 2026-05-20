/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	if root == nil{
		return root
	}
    if root.Right == nil && root.Left ==  nil{
		return root	
	}
	root.Right = 	invertTree(root.Right)
	root.Left = 	invertTree(root.Left)
	t:=root.Right
	root.Right = root.Left
	root.Left = t
	return root
}
