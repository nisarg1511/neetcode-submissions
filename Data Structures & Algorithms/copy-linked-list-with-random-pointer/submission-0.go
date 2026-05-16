/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    curr:=head
	nodeMap:=make(map[*Node]*Node)

	for curr!=nil{
		nodeMap[curr] = &Node{
			Val :curr.Val,
			Next: nil,
			Random: nil,
		}
		curr =  curr.Next
	}

	for k,v:=range nodeMap{
		v.Next = nodeMap[k.Next]
		v.Random = nodeMap[k.Random]
	}

	return nodeMap[head]
}
