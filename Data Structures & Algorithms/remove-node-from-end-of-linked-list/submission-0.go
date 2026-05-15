/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    list:=make([]*ListNode,0)
	curr :=head

	for curr !=nil{
		list = append(list,curr)
		curr = curr.Next
	}
	if len(list) == 1{
		head = nil
	}else{
		i:=len(list)-n
		if i == len(list)-1{
			list[i-1].Next = nil
		}else if i == 0 {
			head = list[i+1]
		}else{
			list[i-1].Next = list[i+1]
		}
	}
	return head
}
