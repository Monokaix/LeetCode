package main

//思路：push时,先将x加入q2,然后将q1元素依次加入q2,再把q2元素一次加入q1
//pop,top直接取q1即可
type MyStack struct {
	q1 []int
	q2 []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	mystack := new(MyStack)
	mystack.q1 = make([]int, 0)
	mystack.q2 = make([]int, 0)
	return *mystack
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.q2 = append(this.q2, x)
	for i := 0; i < len(this.q1); i++ {
		this.q2 = append(this.q2, this.q1[i])
	}
	this.q1 = this.q1[:0]
	for i := 0; i < len(this.q2); i++ {
		this.q1 = append(this.q1, this.q2[i])
	}
	this.q2 = this.q2[:0]
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	val := this.q1[0]
	this.q1 = this.q1[1:]
	return val
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.q1[0]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.q1) == 0
}
