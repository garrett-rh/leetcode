package stoneSmash

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i int, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	n := len(*h)
	ret := (*h)[h.Len()-1]
	*h = (*h)[0 : n-1]

	return ret
}

func LastStoneWeight(stones []int) int {
	stoneHeap := IntHeap(stones)
	heap.Init(&stoneHeap)
	fmt.Println(stoneHeap)

	for stoneHeap.Len() > 1 {
		x, y := heap.Pop(&stoneHeap).(int), heap.Pop(&stoneHeap).(int)
		if x != y {
			heap.Push(&stoneHeap, x-y)
		}
	}

	if stoneHeap.Len() == 0 {
		return 0
	}

	return heap.Pop(&stoneHeap).(int)
}
