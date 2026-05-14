/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 ==nil{
		return list2
	}else if list2==nil{
		return list1
	}
	
    var merged *ListNode
	var head *ListNode
	if list1.Val < list2.Val{
        merged = list1
		list1 = list1.Next
		head=merged
	}else if list2.Val < list1.Val{
		merged = list2
		list2 = list2.Next
		head=merged
	}else{
		merged = list1
		list1 = list1.Next
		merged.Next = list2
		list2 = list2.Next
		head=merged
		merged = merged.Next
	}
	
	for list1!=nil &&  list2!=nil{
		
		if list1.Val == list2.Val{
			merged.Next = list1
			list1 = list1.Next
			merged =  merged.Next
			merged.Next = list2
			list2 = list2.Next
		}else if list1.Val < list2.Val{
			merged.Next = list1
			list1 = list1.Next
		}else {
			merged.Next = list2
			list2 = list2.Next
		}
		merged = merged.Next
	}

	for list1!=nil{
		merged.Next = list1
		list1= list1.Next
		merged = merged.Next
	}

	for list2!=nil{
		merged.Next = list2
		list2= list2.Next
		merged = merged.Next
	}
	return head
}
