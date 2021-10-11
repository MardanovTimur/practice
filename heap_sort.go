package main

import "fmt"

type IHeap interface {
	Head(i int) int
	Left(i int) int
	Right(i int) int
	Length() int
	Heapify(i int, n int)
}

type Heap struct {
	arr []int
}

func (h *Heap) Left(i int) int {
	return h.arr[2*i+1]
}

func (h *Heap) Length() int {
	return len(h.arr)
}

func (h *Heap) Right(i int) int {
	return h.arr[2*i+2]
}

func (h *Heap) Head(i int) int {
	return h.arr[i]
}

func (h *Heap) Heapify(i int, n int) {
	var largest int = i
	if 2*i+1 < n && h.Head(i) < h.Left(i) {
		largest = i*2 + 1
	}
	if 2*i+2 < n && h.arr[largest] < h.Right(i) {
		largest = i*2 + 2
	}
	if largest != i {
		h.arr[i], h.arr[largest] = h.arr[largest], h.Head(i)
		h.Heapify(largest, n)
	}
}

func HeapSort(heap *Heap) {
	for i := (heap.Length() / 2) - 1; i >= 0; i-- {
		heap.Heapify(i, heap.Length())
	}

	for i := heap.Length() - 1; i >= 0; i-- {
		heap.arr[0], heap.arr[i] = heap.arr[i], heap.arr[0]
		heap.Heapify(0, i)
	}
}

func main() {
	arr := []int{4, -1, 123, 120, 10, 0, 3, 5, 1}
	heap := &Heap{arr}
	HeapSort(heap)
	fmt.Printf("Sorted %v \n ", heap.arr)
}
