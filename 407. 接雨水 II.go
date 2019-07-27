package main

import (
	"container/heap"
)

type Info struct {
	x int // The value of the item; arbitrary.
	y int // The priority of the item in the queue.
	h int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue2 []*Info

func (pq PriorityQueue2) Len() int { return len(pq) }

func (pq PriorityQueue2) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].h < pq[j].h //小根堆
}

func (pq PriorityQueue2) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue2) Push(x interface{}) {
	item := x.(*Info)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue2) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func trapRainWater(heightMap [][]int) int {
	pq := make(PriorityQueue2, 0)
	heap.Init(&pq)
	row := len(heightMap)
	if row == 0 {
		return 0
	}
	col := len(heightMap[0])
	if row < 3 || col < 3 { //保证大于等于3
		return 0
	}
	mark := make([][]int, 0) //标记是否访问过
	for i := 0; i < row; i++ {
		tmp := make([]int, col)
		mark = append(mark, tmp)
	}
	//四周进行标记
	for i := 0; i < row; i++ {
		heap.Push(&pq, &Info{i, 0, heightMap[i][0]})
		mark[i][0] = 1
		heap.Push(&pq, &Info{i, col - 1, heightMap[i][col-1]})
		mark[i][col-1] = 1
	}
	for j := 0; j < col; j++ {
		heap.Push(&pq, &Info{0, j, heightMap[0][j]})
		mark[0][j] = 1
		heap.Push(&pq, &Info{row - 1, j, heightMap[row-1][j]})
		mark[row-1][j] = 1
	}
	//扩展的方向数组
	dx := []int{-1, 1, 0, 0};
	dy := []int{0, 0, 1, -1};
	res := 0
	for pq.Len() > 0 {
		x := pq[0].x
		y := pq[0].y
		h := pq[0].h
		heap.Pop(&pq)
		//四个方向扩展
		for i := 0; i < 4; i++ {
			newX := x + dx[i]
			newY := y + dy[i]
			if newX < 0 || newX >= row || newY < 0 || newY >= col || mark[newX][newY] == 1 {
				continue
			}
			//当前节点高于拓展节点
			if h > heightMap[newX][newY] {
				res += h - heightMap[newX][newY]
				heightMap[newX][newY] = h
			}
			heap.Push(&pq, &Info{newX, newY, heightMap[newX][newY]})
			mark[newX][newY] = 1
		}
	}
	return res
}
