package main

import (
	"container/list"
)

//拓扑排序判断是否有环,重点是建立图的邻接表,里面保存每个节点的入度
//首先将所有入度为0的节点放入队列,然后对该节点的邻接节点减1,如果减到0
//也放入队列,循环至队列为空,若所有节点都进入队列,则说明无环返回true
//否则即说明有些节点因为有环而无法放进队列,此时numCourses就不为0
func canFinish(numCourses int, prerequisites [][]int) bool {
	type GraphNode struct {
		val       int   //node的值
		count     int   //入度
		neighbors []int //邻接node
	}
	graph := make([]GraphNode, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		first := prerequisites[i][0]
		second := prerequisites[i][1]
		graph[first].count++
		graph[second].neighbors = append(graph[second].neighbors, first)
	}
	l := list.New()
	for i := 0; i < numCourses; i++ {
		if graph[i].count == 0 {
			l.PushBack(graph[i])
		}
	}
	for l.Len() != 0 {
		top := l.Front().Value.(GraphNode)
		l.Remove(l.Front())
		numCourses--
		node := top.neighbors
		for _, v := range node {
			graph[v].count--
			if graph[v].count == 0 {
				l.PushBack(graph[v])
			}
		}
	}
	if numCourses == 0 {
		return true
	}
	return false
}
