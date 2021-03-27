package main

// 空间复杂度为O(n)的算法
//type MinStack struct {
//	data []int
//	// 最小值栈存的是最小值在栈中的索引值，避免重复存最小值
//	min []int
//}
//
//func Constructor() MinStack {
//	return MinStack{
//		data: make([]int, 0),
//		min:  make([]int, 0),
//	}
//}
//
//func (this *MinStack) Push(val int) {
//	this.data = append(this.data, val)
//
//	if len(this.min) == 0 {
//		// 存的是数据栈元素的索引，即数据栈当前长度-1
//		this.min = append(this.min, 0)
//		return
//	}
//
//	minVal := this.GetMin()
//	if val < minVal {
//		this.min = append(this.min, len(this.data)-1)
//	}
//}
//
//func (this *MinStack) Pop() {
//	if len(this.data) == 0 {
//		return
//	}
//
//	// 比较索引
//	dataIndex := len(this.data) - 1
//	minIndex := this.min[len(this.min)-1]
//
//	if dataIndex == minIndex {
//		this.min = this.min[:len(this.min)-1]
//	}
//
//	this.data = this.data[:len(this.data)-1]
//}
//
//func (this *MinStack) Top() int {
//	return this.data[len(this.data)-1]
//}
//
//func (this *MinStack) GetMin() int {
//	minIndex := this.min[len(this.min)-1]
//	return this.data[minIndex]
//}

// 方法二：空间复杂度O(1)
// 栈存的是栈顶元素和最小值的差值
type MinStack struct {
	data []int
	// 最小值栈存的是最小值在栈中的索引值，避免重复存最小值
	min int
}

func Constructor2() MinStack {
	return MinStack{
		data: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	if len(this.data) == 0 {
		// 初始时放入0，代表当前元素与即为最小值，差值为0
		this.data = append(this.data, 0)
		this.min = val
	} else {
		diff := val - this.min
		// 当前元素比最小值小，需要更新最小值
		if diff < 0 {
			this.min = val
		}
		this.data = append(this.data, diff)
	}
}

func (this *MinStack) Pop() {
	if len(this.data) == 0 {
		return
	}

	top := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	// 说明此时最小值被弹出，需要更新最小值
	if top < 0 {
		this.min -= top
	}
}

func (this *MinStack) Top() int {
	top := this.data[len(this.data)-1]

	// 说明此时真实元素即为最小值，不然不会为0
	if top < 0 {
		return this.min
	}
	return top + this.min
}

func (this *MinStack) GetMin() int {
	return this.min
}
