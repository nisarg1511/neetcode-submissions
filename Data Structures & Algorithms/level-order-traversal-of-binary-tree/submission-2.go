/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
   	arr:=make([][]int,0)
	if root == nil {
		return arr
	}
	left :=levelOrder(root.Left)
	right:=levelOrder(root.Right)
	i:=0
	j:=0
	arr = append(arr,[]int{root.Val})

	for i < len(left) && j < len(right){
		arr = append(arr,merge(left[i],right[i]))
		i++
		j++
	}
	for i < len(left){
		arr = append(arr,left[i])
		i++
	}
	for  j< len(right){
		arr = append(arr,right[j])
		j++
	}
	return arr
}

func merge(left []int,right[]int)[]int{
	merged:=make([]int,0)
	for _,v:=range left{
		merged = append(merged,v)
	}
	for _,v:=range right{
		merged = append(merged,v)
	}
	return merged
}
