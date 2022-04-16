package listLoops

import "testing"

type List struct {
	Prev *ListNode
	Next *ListNode
}

func TestHasCycle(t *testing.T) {
	head := newNode(3, nil)
	head.Next = newNode(2, head)
	head.Next.Next = newNode(0, head.Next)
	head.Next.Next.Next = newNode(-4, head.Next.Next)

	got := HasCycle(head)
	if got != true {
		t.Errorf("got %v, want %v", got, true)
	}
}

func newNode(value int, address *ListNode) *ListNode {
	var node ListNode
	node.Val = value
	node.Next = address
	return &node
}
