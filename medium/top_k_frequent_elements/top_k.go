package topKFrequentElements

import "container/heap"

type pair struct {
	key   int
	value int
}

type MaxHeap []pair

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m MaxHeap) Less(i, j int) bool {
	return m[i].value > m[j].value
}

func (m *MaxHeap) Push(x interface{}) {
	*m = append(*m, x.(pair))
}

func (m *MaxHeap) Pop() interface{} {
	old := *m
	n := len(old)
	*m = old[0 : n-1]
	return old[n-1]
}

func TopKFrequent(nums []int, k int) []int {

	if k == len(nums) {
		return nums
	}

	hashMap := make(map[int]int)
	for _, num := range nums {
		hashMap[num]++
	}

	m := MaxHeap{}

	for k, v := range hashMap {
		heap.Push(&m, pair{key: k, value: v})
	}

	topKElements := []int{}

	for i := 0; i < k; i++ {
		top := heap.Pop(&m).(pair).key
		topKElements = append(topKElements, top)
	}

	return topKElements

}

/* First solution
func topKFrequent(nums []int, k int) []int {
	if k == len(nums) {
		return nums
	}
	hashMap := make(map[int]int)
	freq := []pair{}

	for _, num := range nums {
		hashMap[num]++
	}

	for k, v := range hashMap {
		freq = append(freq, pair{k, v})
	}

	sort.Slice(freq, func(i, j int) bool {
		return freq[i].value > freq[j].value
	})

	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[i] = freq[i].key
	}

	return ret

    }

type pair struct {
    key int
    value int
}
*/
