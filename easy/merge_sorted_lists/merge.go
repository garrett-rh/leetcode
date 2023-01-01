package merge

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var (
		splicedList *ListNode
		previous    *ListNode
	)
	if list1.Val <= list2.Val {
		splicedList = list1
		list1 = list1.Next
	} else {
		splicedList = list2
		list2 = list2.Next
	}
	previous = splicedList

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			previous.Next = list1
			previous = list1
			list1 = list1.Next
		} else {
			previous.Next = list2
			previous = list2
			list2 = list2.Next
		}
	}

	if list1 == nil {
		previous.Next = list2
	} else {
		previous.Next = list1
	}

	return splicedList
}
