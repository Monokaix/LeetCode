package main

//思路：push:直接进stack1
//pop: 若s2为空,将s1中元素全部压入s2,取s2栈顶元素
//若s2不为空,直接取s2栈顶元素
type MyQueue struct {
	top1   int
	top2   int
	stack1 []int
	stack2 []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	queue := new(MyQueue)
	queue.top1 = 0
	queue.top2 = 0
	queue.stack1 = make([]int, 0)
	queue.stack2 = make([]int, 0)
	return *queue
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stack1 = append(this.stack1, x)
	this.top1++
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.top2 == 0 {
		for this.top1 > 0 {
			this.top1--
			this.stack2 = append(this.stack2, this.stack1[this.top1])
			this.top2++
		}
		this.stack1 = this.stack1[:0]
	}
	this.top2--
	val := this.stack2[this.top2]
	this.stack2 = this.stack2[:this.top2]
	return val
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.top2 == 0 {
		for this.top1 > 0 {
			this.top1--
			this.stack2 = append(this.stack2, this.stack1[this.top1])
			this.top2++
		}
		this.stack1 = this.stack1[:0]
	}
	return this.stack2[this.top2-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if this.top2 == this.top1 && this.top1 == 0 {
		return true
	}
	return false
}
