package main

import (
	"container/list"
)

type AnimalShelf struct {
	cat   list.List
	dog   list.List
	count int
}

func Constructor() AnimalShelf {
	return AnimalShelf{
		cat: list.List{},
		dog: list.List{},
	}
}

func (this *AnimalShelf) Enqueue(animal []int) {
	val := make([]int, 2)
	if animal[1] == 0 {
		val = []int{this.count, 0}
		this.cat.PushBack(val)
	} else {
		val = []int{this.count, 1}
		this.dog.PushBack(val)
	}
	this.count++
}

func (this *AnimalShelf) DequeueAny() []int {
	if this.cat.Len() == 0 && this.dog.Len() == 0 {
		return []int{-1, -1}
	}

	if this.cat.Len() == 0 {
		res := this.dog.Front().Value.([]int)
		this.dog.Remove(this.dog.Front())
		return res
	}

	if this.dog.Len() == 0 {
		res := this.cat.Front().Value.([]int)
		this.cat.Remove(this.cat.Front())
		return res
	}

	c := this.cat.Front().Value.([]int)
	d := this.dog.Front().Value.([]int)
	if c[0] < d[0] {
		this.cat.Remove(this.cat.Front())
		return c
	}

	this.dog.Remove(this.dog.Front())
	return d
}

func (this *AnimalShelf) DequeueDog() []int {
	if this.dog.Len() == 0 {
		return []int{-1, -1}
	}
	res := this.dog.Front().Value.([]int)
	this.dog.Remove(this.dog.Front())
	return res
}

func (this *AnimalShelf) DequeueCat() []int {
	if this.cat.Len() == 0 {
		return []int{-1, -1}
	}
	res := this.cat.Front().Value.([]int)
	this.cat.Remove(this.cat.Front())
	return res
}

/**
 * Your AnimalShelf object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Enqueue(animal);
 * param_2 := obj.DequeueAny();
 * param_3 := obj.DequeueDog();
 * param_4 := obj.DequeueCat();
 */
