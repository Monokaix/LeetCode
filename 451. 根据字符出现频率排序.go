package main

import (
	"fmt"
	"sort"
)

type pair1 struct {
	char byte
	cnt  int
}

type pairList []pair1

func (p pairList) Len() int { return len(p) }
func (p pairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p pairList) Less(i, j int) bool {
	return p[i].cnt > p[j].cnt
}
func frequencySort(s string) string {
	arr := [256]int{}
	p := make(pairList, 0)
	for i := range s {
		arr[s[i]]++
	}
	for k, v := range arr {
		if v != 0 {
			p = append(p, pair1{byte(k), v})
		}

	}
	sort.Sort(p)
	res := ""
	for i := range p {
		for j := 0; j < p[i].cnt; j++ {
			res += string(p[i].char)
		}
	}
	return res
}

func main() {
	s := "tree"
	fmt.Println(frequencySort(s))
}
