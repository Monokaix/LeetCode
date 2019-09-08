package main

import "container/heap"

type Pair struct {
	cost   int
	profit int
}
type Heap []*Pair

func (h Heap) Len() int { return len(h) } // 绑定len方法,返回长度
func (h Heap) Less(i, j int) bool { // 绑定less方法
	return h[i].cost < h[j].cost // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}
func (h Heap) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *Heap) Push(x interface{}) { // 绑定push方法，插入新元素
	*h = append(*h, x.(*Pair))
}
func (h *big) Less(i, j int) bool { // 绑定less方法
	return h.Heap[i].profit > h.Heap[j].profit // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}

type small struct {
	// 小根堆
	Heap
}
type big struct {
	// 大根堆
	Heap
}

//以上建立了大根堆和小根堆,然后先将数据按成本放入小根堆,再依次取出满足条件的放入大根堆,每次关注大根堆即可
func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
	s := new(small)
	b := new(big)
	heap.Init(s)
	heap.Init(b)
	for i := 0; i < len(Profits); i++ {
		p := &Pair{
			Capital[i],
			Profits[i],
		}
		heap.Push(s, p)
	}
	for i := 0; i < k; i++ {
		for len(s.Heap) > 0 && s.Heap[0].cost <= W { //小根堆里的成本比当前w小的就一直出堆放入大根堆
			heap.Push(b, s.Heap[0])
			heap.Pop(s)
		}
		//如果大根堆为空,当前资金不够,直接返回
		if len(b.Heap) == 0 {
			return W
		}
		W += b.Heap[0].profit
		heap.Pop(b)
	}
	return W
}
