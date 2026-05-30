
func buildTree(preorder []int, inorder []int) *TreeNode {
	in := make(map[int]int)

	for i := range inorder {
		in[inorder[i]] = i
	}
	index:=0
	return build(0,len(inorder)-1,preorder,in,&index)
}

func build(start,end int, preorder []int,m map[int]int,i *int) *TreeNode {
	if start  > end {
		return nil
	}
	node:=&TreeNode{
		Val:preorder[*i],
	}
	*i+=1
	inorderIndex := m[node.Val]
	node.Left = build(start,inorderIndex-1,preorder,m,i)
	node.Right = build(inorderIndex+1,end,preorder,m,i)
	return node
}

