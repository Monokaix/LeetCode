package main

import "container/list"
//利用宽度搜索可以保证是最短的
func ladderLength(beginWord string, endWord string, wordList []string) int {
	l := list.New()
	type info struct {
		word string
		step int
	}
	l.PushBack(info{beginWord,0})
	visited := make(map[string]bool)
	visited[beginWord] = true
	for l.Len() > 0{
		word := l.Front().Value.(info).word
		step := l.Front().Value.(info).step
		l.Remove(l.Front())
		// Find a word that differs by one letter in wordList
		for i:= 0;i<len(wordList);i++{
			different := 0
			for j := 0;j<len(wordList[i]);j++{
				if word[j] != wordList[i][j]{
					different++
				}
			}
			if different == 1{
				if _,ok := visited[wordList[i]];!ok{
					l.PushBack(info{wordList[i],step+1})
					visited[wordList[i]] = true
				}
			}
		}
		if word == endWord{ //find the shortest path
			return step+1
		}
	}
	return 0
}