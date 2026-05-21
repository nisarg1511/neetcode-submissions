/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
	
func diameterOfBinaryTree(root *TreeNode) int {
	max:=0
	calcDiameter(root,&max)
	return max
}

func calcDiameter(root *TreeNode,max *int) int {
	if root == nil {
		return 0
	}
	left:=calcDiameter(root.Left,max)
	right:=calcDiameter(root.Right,max)

	if left+right > *max{
		*max = left + right
	}

	if left > right {
		return left  + 1
	}
	return  right +1 
}