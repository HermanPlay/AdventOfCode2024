package day1

import (
	"fmt"
	"math"
	"strings"
)

type Heap struct {
	heap   []int
	length int
}

func (h *Heap) Pop() int {
	if h.length == 0 {
		panic("poping from empty heap")
	}
	value := h.heap[0]
	h.heap[0] = math.MaxInt
	h.pushDown(0)
	return value
}

func (h *Heap) Insert(value int) {
	h.heap[h.length] = value
	h.length += 1

	h.pushUp(h.length - 1)
}

func (h *Heap) pushUp(index int) {
	if index == 0 {
		return
	}

	var parent int = (index - 1) / 2
	if h.heap[index] < h.heap[parent] {
		temp := h.heap[index]
		h.heap[index] = h.heap[parent]
		h.heap[parent] = temp
		h.pushUp(parent)
	}
}

func (h *Heap) pushDown(index int) {
	if index == h.length-1 {
		return
	}
	leftChild := 2*index + 1
	rightChild := 2*index + 2
	smallestChild := -1
	if leftChild > h.length {
		// No left or right child, finish
		return
	}

	if h.heap[leftChild] < h.heap[index] {
		smallestChild = leftChild
	}
	if rightChild < h.length {
		if smallestChild == -1 && h.heap[rightChild] < h.heap[index] || smallestChild != -1 && h.heap[rightChild] < h.heap[smallestChild] {
			smallestChild = rightChild
		}
	}
	if smallestChild == -1 {
		// Pushed the lowest
		return
	}

	temp := h.heap[index]
	h.heap[index] = h.heap[smallestChild]
	h.heap[smallestChild] = temp
	if h.heap[index] == math.MaxInt {
		fmt.Println(h.String())
		panic("swapped with max")
	}
	// fmt.Printf("After pushed down: parent: %d child: %d\n", h.heap[index], h.heap[smallestChild])
	h.pushDown(smallestChild)
}

func (h *Heap) String() string {
	if h.length == 0 {
		return ""
	}

	var result strings.Builder

	// Calculate tree depth
	var depth int = int(math.Log2(float64(h.length))) + 1

	levelStart := 0
	for i := 0; i < depth; i++ {
		levelSize := 1 << i // 2^i elements at level i
		levelEnd := levelStart + levelSize
		if levelEnd > h.length {
			levelEnd = h.length
		}
		// Calculate spacing for alignment
		space := (1 << (depth - i)) - 1 // 2^(depth - i) - 1
		line := strings.Repeat(" ", space)

		for j := levelStart; j < levelEnd; j++ {
			line += fmt.Sprintf("%d", h.heap[j])
			if j != levelEnd-1 {
				line += strings.Repeat(" ", 2*space+1)
			}
		}
		line += "\n"

		result.WriteString(line)
		levelStart = levelEnd
	}

	return result.String()
}
func NewHeap(size int) *Heap {
	if size < 0 {
		panic("heap cannot have negative size")
	}
	return &Heap{
		heap: make([]int, size+1),
		// Start with 1, for easier calculation, we treat elememnt at index 0 as a dummy value
		length: 0,
	}
}
