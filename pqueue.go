package pqueue

import (
	"cmp"
	"container/heap"
)

// Package pqueue implements a generic priority queue data structure using a max-heap.
//
// The priority queue maintains elements in descending order of priority, where the
// element with the highest priority value is always at the front of the queue.
// Each element consists of two components:
//   - A value of any type T
//   - A priority of any ordered type P (numbers, strings, etc.)
//
// The implementation uses Go's container/heap package internally, providing O(log n)
// time complexity for Push and Pop operations, and O(1) for Peek and Len operations.
//
// Example usage:
//
//	pq := pqueue.New[string, int]()
//	pq.Push("task1", 3)
//	pq.Push("task2", 1)
//	pq.Push("task3", 4)
//
//	value, priority, ok := pq.Pop() // Returns "task3", 4, true
//
// All operations are safe to use with empty queues, returning appropriate zero
// values and a boolean false when attempting to access empty queues.

// PriorityQueue represents a max-heap implementation of a priority queue.
type PriorityQueue[T any, P cmp.Ordered] struct {
	items *itemHeap[T, P]
}

// item represents a single element in the priority queue.
type item[T any, P cmp.Ordered] struct {
	value    T
	priority P
}

// itemHeap implements heap.Interface
type itemHeap[T any, P cmp.Ordered] []item[T, P]

func (h itemHeap[T, P]) Len() int           { return len(h) }
func (h itemHeap[T, P]) Less(i, j int) bool { return h[i].priority > h[j].priority } // Max heap
func (h itemHeap[T, P]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *itemHeap[T, P]) Push(x any) {
	*h = append(*h, x.(item[T, P]))
}

func (h *itemHeap[T, P]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// New creates a new empty priority queue.
func New[T any, P cmp.Ordered]() *PriorityQueue[T, P] {
	items := &itemHeap[T, P]{}
	heap.Init(items)
	return &PriorityQueue[T, P]{
		items: items,
	}
}

// Push adds a new element to the priority queue.
func (pq *PriorityQueue[T, P]) Push(value T, priority P) {
	heap.Push(pq.items, item[T, P]{value: value, priority: priority})
}

// Pop removes and returns the highest priority element.
// Returns zero values and false if the queue is empty.
func (pq *PriorityQueue[T, P]) Pop() (T, P, bool) {
	if pq.IsEmpty() {
		var zero T
		var zeroP P
		return zero, zeroP, false
	}
	item := heap.Pop(pq.items).(item[T, P])
	return item.value, item.priority, true
}

// Peek returns the highest priority element without removing it.
// Returns zero values and false if the queue is empty.
func (pq *PriorityQueue[T, P]) Peek() (T, P, bool) {
	if pq.IsEmpty() {
		var zero T
		var zeroP P
		return zero, zeroP, false
	}
	item := (*pq.items)[0]
	return item.value, item.priority, true
}

// Len returns the number of elements in the queue.
func (pq *PriorityQueue[T, P]) Len() int {
	return pq.items.Len()
}

// IsEmpty returns true if the queue has no elements.
func (pq *PriorityQueue[T, P]) IsEmpty() bool {
	return pq.Len() == 0
}
