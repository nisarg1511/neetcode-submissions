/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    queue:=make([]*ListNode,0)

	curr:=head.Next

	for curr!=nil{
		queue=append(queue,curr)
		curr = curr.Next
	}

	curr = head
	for len(queue) > 0 {
		curr.Next = queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		curr = curr.Next
		if len(queue) > 0{
			curr.Next = queue[0]
			queue = queue[1:]
			curr= curr.Next
		}else{
			break
		}
	}
	curr.Next = nil
}
