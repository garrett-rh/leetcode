package listLoops

type ListNode struct {
	Val  int
	Next *ListNode
}

func HasCycle(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head.Next, head.Next.Next

	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next

	}
	return true
}
