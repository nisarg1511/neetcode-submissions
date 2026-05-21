/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	maxDiff:=0
	calculateHeightDifference(root,&maxDiff)
	if maxDiff > 1 {
		return false
	}
	return true
}

func calculateHeightDifference(root *TreeNode, maxDiff *int) int{
	if root == nil {
		return 0
	}
	left,right:=calculateHeightDifference(root.Left,maxDiff),calculateHeightDifference(root.Right,maxDiff)
	if Abs(left - right) >  *maxDiff{
		*maxDiff = Abs(left - right)
	}
	
	if left > right {
		return left+ 1
	}
	return right + 1
}

func Abs(n int) int{
	if n < 0{
		return  n*-1
	}
	return n
}