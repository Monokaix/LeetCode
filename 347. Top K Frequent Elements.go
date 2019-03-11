package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    int // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	//index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	//pq[i].index = i
	//pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	//n := len(*pq)
	item := x.(*Item)
	//item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	//item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	res := make([]int,k)
	m := make(map[int]int)
	for _,v := range nums{
		m[v]++
	}
	pq := make(PriorityQueue,0)
	heap.Init(&pq)

	for index,v := range m{
		if pq.Len() == k{
			if v > pq[0].priority {
				heap.Pop(&pq)
				heap.Push(&pq, &Item{index, v})
			}
		}else{
			heap.Push(&pq,&Item{index,v})
		}
	}
	for pq.Len() > 0{
		item := heap.Pop(&pq).(*Item).value
		res[k-1] = item
		k--

	}
	return res
}

func main(){
	nums := []int{6,6,6,6,5,5,5,3,4,4}
	res := topKFrequent(nums,3)
	fmt.Println(res)
}