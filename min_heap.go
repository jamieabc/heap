package main

import "fmt"

type MinHeap interface {
	Insert(int)
	Remove() int
	fmt.Stringer
}

type minHeapData struct {
	size int
	data []int
}

func (h minHeapData) parent(index int) int {
	return (index - 1) / 2
}

func (h minHeapData) leftChild(index int) int {
	return index*2 + 1
}

func (h minHeapData) rightChild(index int) int {
	return index*2 + 2
}

func (h *minHeapData) swap(dst, src int) {
	h.data[src], h.data[dst] = h.data[dst], h.data[src]
}

func (h *minHeapData) Insert(num int) {
	h.data = append(h.data, num)
	h.size++
	for index := h.size - 1; index > 0; {
		parent := h.parent(index)
		if h.data[parent] > h.data[index] {
			h.swap(parent, index)
			index = parent
		} else {
			break
		}
	}
}

func (h *minHeapData) Remove() int {
	removed := h.data[0]

	h.data[0] = h.data[h.size-1]
	h.data = h.data[:h.size-1]
	h.size--

	var left, right int
	for index := 0; index < h.size/2; {
		left = h.leftChild(index)
		right = h.rightChild(index)
		if h.data[index] > h.data[left] {
			h.swap(left, index)
			index = left
		} else if h.data[index] > h.data[right] {
			h.swap(right, index)
			index = right
		} else {
			break
		}
	}
	return removed
}

func (h minHeapData) String() string {
	return fmt.Sprintf("%#v", h.data)
}

func CreateHeap() MinHeap {
	return &minHeapData{
		size: 0,
		data: make([]int, 0),
	}
}
