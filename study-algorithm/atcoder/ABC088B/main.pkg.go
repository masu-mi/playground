package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	h := &intHeap{}
	var tmp int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &tmp)
		h.Push(tmp)
	}
	var diff int
	move := 1
	for h.Len() != 0 {
		diff += move * *(heap.Pop(h)).(*int)
		move *= -1
	}
	fmt.Printf("%d\n", diff)
}

type intHeap struct {
	sort.IntSlice
}

func (h *intHeap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}
func (h *intHeap) Pop() interface{} {
	item := ([]int)(h.IntSlice)[len(h.IntSlice)-1]
	h.IntSlice = ([]int)(h.IntSlice)[0 : len(h.IntSlice)-1]
	return &item
}
