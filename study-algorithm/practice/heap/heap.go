package heap

import (
	"fmt"
)

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

type Heap struct {
	num    int
	buffer []int
}

func NewHeap(size int) *Heap {
	return &Heap{num: 0, buffer: make([]int, 0, size)}
}

func (h *Heap) IsEmpty() bool {
	return h.num == 0
}

func (h *Heap) Push(v int) {
	h.buffer = append(h.buffer, v)
	h.num++
	for i := h.num; i/2 > 0; {
		if h.buffer[i-1] <= h.buffer[i/2-1] {
			break
		}
		h.buffer[i/2-1], h.buffer[i-1] = h.buffer[i-1], h.buffer[i/2-1]
		i /= 2
	}
}
func (h *Heap) Pop() int {
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
	h.buffer = h.buffer[0:h.num]
	return v
}
