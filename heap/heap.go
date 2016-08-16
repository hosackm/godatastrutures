package main

import (
	"errors"
	"fmt"
)

const (
	MAXLENGTH = 65536
)

type Elem int

// Max Heap
type Heap struct {
	size int
	data []Elem
}

func leftchild(i int) int {
	return 2*n + 1
}

func rightchild(i int) int {
	return 2*n + 2
}

func parent(i int) int {
	return (n - 1) / 2
}

func (h *Heap) get(i int) Elem {

}

func (h *Heap) compareParent(i int) {
	if i < 1 {
		return
	}

	if h.get(i) > h.get(parent(i)) {
		h.data[i], h.data[parent(i)] = h.data[parent(i)], h.data[i]
		compareParent(i)
	}
}

func (h *Heap) percolateUp(index int) {

}

func (h *Heap) balance() {
	h.percolateDown(0)
}

func (h *Heap) percolateDown(index int) {
	value := h.get(index)
	l, r := leftchild(index), rightchild(index)

	if l <= h.size && value < h.get(l) {
		// swap with left
		h.data[index], h.data[l] = h.data[l], h.data[index]
		// run this func on leftchild
		percolateDown(l)
	} else if r <= size && value < h.get(r) {
		// swap with right
		h.data[index], h.data[r] = h.data[r], h.data[index]
		// run this func on rightchild
		percolateDown(r)
	}
}

func (h *Heap) Push(val Elem) {

}

func (h *Heap) Pop() (Elem, error) {
	if h.size == 0 {
		return -1, errors.New("Empty heap")
	}

	// return the head of heap
	ret := h.data[0]

	// move last index to the front and find it's place in the heap
	h.data[0] = h.data[h.size-1]
	h.size--
	h.balance()

	return ret, nil
}

func (h *Heap) String() string {
	s := "Heap["
	for i := 0; i < h.size; i++ {
		if i > 0 {
			s += ","
		}
		s += fmt.Sprintf(" %d", h.data[i])
	}
	return s + "]"
}

func NewHeap() *Heap {
	h := &Heap{0, make([]Elem, MAXLENGTH)}
	return h
}

func main() {
	h := NewHeap()
	fmt.Println(h)
}
