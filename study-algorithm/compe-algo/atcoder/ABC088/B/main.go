package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	h := NewHeap(n)
	var tmp int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &tmp)
		h.Push(tmp)
	}
	var diff int
	move := 1
	for !h.IsEmpty() {
		diff += move * h.Pop()
		move *= -1
	}
	fmt.Printf("%d\n", diff)
}

type heap struct {
	num    int
	buffer []int
}

func NewHeap(size int) *heap {
	return &heap{num: 0, buffer: make([]int, size)}
}

func (h *heap) IsEmpty() bool {
	return h.num == 0
}

func (h *heap) Push(v int) {
	h.num++
	h.buffer[h.num-1] = v
	for i := h.num; i/2 > 0; {
		if h.buffer[i-1] <= h.buffer[i/2-1] {
			break
		}
		h.buffer[i/2-1], h.buffer[i-1] = h.buffer[i-1], h.buffer[i/2-1]
		i /= 2
	}
}
func (h *heap) Pop() int {
	v := h.buffer[0]
	h.buffer[0] = h.buffer[h.num-1]
	for i := 1; 2*i <= h.num; {
		if 2*i+1 > h.num || h.buffer[2*i] <= h.buffer[2*i-1] {
			if h.buffer[i-1] >= h.buffer[i*2-1] {
				break
			}
			h.buffer[i-1], h.buffer[i*2-1] = h.buffer[i*2-1], h.buffer[i-1]
			i = i * 2
		} else {
			if h.buffer[i-1] >= h.buffer[i*2] {
				break
			}
			h.buffer[i-1], h.buffer[i*2] = h.buffer[i*2], h.buffer[i-1]
			i = i*2 + 1
		}
	}
	h.num--
	return v
}
