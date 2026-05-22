/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil  && q == nil {
		return true
	}
	if p==nil || q == nil{
		return false
	}
	if p.Val == q.Val {
		left:= isSameTree(p.Left,q.Left)
		right:=isSameTree(p.Right,q.Right)
		if left && right {
		if p.Left == nil && q.Left == nil || p.Right == nil && q.Right==nil{
			return true
		}
			if p.Left.Val == q.Left.Val && p.Right.Val == q.Right.Val{
				return true
			}
		}
	}
	

	
	return false
}
