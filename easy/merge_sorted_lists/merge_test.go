package merge

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	testCases := []struct {
		list1 *ListNode
		list2 *ListNode
		want  *ListNode
	}{
		{list1: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}},
			list2: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}},
			want:  &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: nil}}}}}}},

		{list1: &ListNode{Next: nil},
			list2: &ListNode{Next: nil},
			want:  &ListNode{Next: nil}},

		{list1: &ListNode{Val: 0, Next: nil},
			list2: &ListNode{Next: nil},
			want:  &ListNode{Val: 0, Next: nil}},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("splice arrays %v and %v", test.list1, test.list2), func(t *testing.T) {
			got := MergeTwoLists(test.list1, test.list2)

			if got == test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})
	}
}
