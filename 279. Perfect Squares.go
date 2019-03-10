package main

import "container/list"

func numSquares(n int) int {
	l := list.New()
	type info struct {
		num int
		step int
	}
	l.PushBack(info{n,0})
	visited := make([]bool,n+1)
	visited[n] = true
	for l.Len()>0{
		num := l.Front().Value.(info).num
		step := l.Front().Value.(info).step
		l.Remove(l.Front())

		for i := 0;;i++{
			a := num - i*i
			if a <0 {
				break
			}
			if a == 0{
				return step+1
			}
			if !visited[a] {
				l.PushBack(info{a,step+1})
			}
			visited[a] = true
		}
	}
	return 0
}
