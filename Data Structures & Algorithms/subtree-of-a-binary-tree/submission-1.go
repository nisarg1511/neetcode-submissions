/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {

	if root==nil && subRoot == nil {
		return true
	}
	if root == nil || subRoot == nil{
		return false
	}
	if root.Val == subRoot.Val {
		sol:=compare(root,subRoot)
		if sol{
			return true
		}
	}

	left:=isSubtree(root.Left,subRoot)
	right:=isSubtree(root.Right,subRoot)
	return left || right
}

func compare(t1 *TreeNode, t2 *TreeNode) bool{
	if t1 == nil  && t2 == nil{
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val == t2.Val{
		left:=compare(t1.Left,t2.Left)
		right:=compare(t1.Right,t2.Right)
		return left && right	
	}
	return false
}