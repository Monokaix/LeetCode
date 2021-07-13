package main

import (
	"sort"
	"strconv"
)

type TimeMap struct {
	key2TimeStamp   map[string][]int
	timeStamp2Value map[string]string
}

/** Initialize your data structure here. */
func Constructor4() TimeMap {
	return TimeMap{
		key2TimeStamp:   make(map[string][]int),
		timeStamp2Value: make(map[string]string),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.timeStamp2Value[key+strconv.Itoa(timestamp)] = value
	arr := this.key2TimeStamp[key]
	arr = append(arr, timestamp)
	this.key2TimeStamp[key] = arr
}

func (this *TimeMap) Get(key string, timestamp int) string {
	arr := this.key2TimeStamp[key]
	if len(arr) == 0 {
		return ""
	}
	index := sort.SearchInts(arr, timestamp)
	if index == len(arr) {
		return this.timeStamp2Value[key+strconv.Itoa(arr[index-1])]
	}
	if arr[index] == timestamp {
		return this.timeStamp2Value[key+strconv.Itoa(timestamp)]
	}
	if index > 0 {
		timestamp = arr[index-1]
		return this.timeStamp2Value[key+strconv.Itoa(timestamp)]
	}
	return ""
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
