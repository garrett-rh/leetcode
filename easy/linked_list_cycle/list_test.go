package listLoops

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head   *ListNode
	Tail   *ListNode
	Length int
}

func TestHasCycle(t *testing.T) {
	testCases := []struct {
		data []int
		pos  int
		res  bool
	}{
		{data: []int{3, 2, 0, -4}, pos: 1, res: true},
		{data: []int{1, 2}, pos: 0, res: true},
		{data: []int{1}, pos: -1, res: false},
	}
	for _, test := range testCases {
		t.Run(fmt.Sprintf("Testing if a the list %+v with position %d loops", test.data, test.pos), func(t *testing.T) {
			currentList := LinkedList{}
			for _, data := range test.data {
				currentList.Append(data)
			}
			got := HasCycle(currentList.Head)
			fmt.Println(currentList)

			if got != test.res {
				t.Errorf("HasCycle returned %v instead of %v", got, test.res)
			}
		})
	}

}

func (list *LinkedList) Append(data int) {
	newNode := &ListNode{
		Val: data,
	}

	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
	} else {
		list.Tail.Next = newNode
		list.Tail = newNode
	}

	list.Length++
}
