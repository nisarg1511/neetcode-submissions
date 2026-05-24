/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    res:=make([][]int,0,1000)
	if root == nil {
		return res
	}
	
	queue:=make([]*TreeNode,0,1000)
	queue = append(queue,root)
	for len(queue) > 0{
		level:=make([]int,len(queue))
		for i:=range level{
			node:=queue[0]
			level[i] = node.Val
			queue = queue[1:]
			if node.Left!=nil{
				queue = append(queue,node.Left)
			}
			if node.Right !=nil{
				queue = append(queue,node.Right)
			}
		}
		res = append(res,level)
	}

	return res
}
