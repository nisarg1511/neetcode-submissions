/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	hMap:=make(map[*ListNode]*ListNode)
	curr:=head
	
	// if head == nil || head.Next == nil{
	// 	return false
	// }

	for curr!=nil{
		if hMap[curr] != nil{
			return true
		}
		hMap[curr] = curr.Next
		curr = curr.Next
	}

	return false
}
