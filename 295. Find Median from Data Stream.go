package main

import (
	"fmt"
	"container/heap"
)

// 使用最大最小堆实现,维护两个堆,一个大根堆,一个小根堆,并且两个堆元素个数差不超过1
// 这样大根堆存储前半部分有序数据,小根堆存储后半部分有序数据
// 然后根据两个堆大小判断把来的数据放在哪个堆里 AddNum(num int)
// 若两个堆元素个数相等,取两个堆顶元素平均数即为中位数
// 否则谁的元素多取谁的堆顶元素为中位数 FindMedian()
type intHeap []int

func (h intHeap) Len() int { return len(h) } // 绑定len方法,返回长度
func (h intHeap) Less(i, j int) bool { // 绑定less方法
	return h[i] < h[j] // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}
func (h intHeap) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *intHeap) Push(x interface{}) { // 绑定push方法，插入新元素
	*h = append(*h, x.(int))
}

type smallHeap struct {
	// 小根堆
	intHeap
}
type bigHeap struct {
	// 大根堆
	intHeap
}

func (h bigHeap) Less(i, j int) bool { // 绑定less方法
	return h.intHeap[i] > h.intHeap[j] // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}

type MedianFinder struct {
	left  *bigHeap
	right *smallHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	right := new(smallHeap)
	heap.Init(right)
	left := new(bigHeap)
	heap.Init(left)
	this := MedianFinder{
		left,
		right,
	}
	return this
}

func (this *MedianFinder) AddNum(num int) {
	if this.left.Len() == this.right.Len() {//每次插入大顶堆，保证大顶堆元素始终大于等于小顶堆元素
		heap.Push(this.left, num)
	} else {
		heap.Push(this.right, num)
	}

	if this.right.Len() > 0 && this.left.intHeap[0] > this.right.intHeap[0] { //这里用来保证大顶堆元素始终小于小顶堆
		this.left.intHeap[0], this.right.intHeap[0] = this.right.intHeap[0], this.left.intHeap[0]
		heap.Fix(this.left, 0)
		heap.Fix(this.right, 0)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.left.Len() == this.right.Len() {
		return float64(this.left.intHeap[0]+this.right.intHeap[0]) / 2
	}
	return float64(this.left.intHeap[0])
}
func main() {
	this := Constructor()
	this.AddNum(-1)
	this.AddNum(-2)
	this.AddNum(-3)
	this.AddNum(-4)
	this.AddNum(-5)
	fmt.Println(this.left)
	fmt.Println(this.right)
	fmt.Println(this.FindMedian())
}
