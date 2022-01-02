// Package pqueue provides a very simple priority queue
// It's basically copy and paste from the heap package example, only to override the Push method.
package pqueue

import "container/heap"

// A PriorityQueue (simple wrap to override Push)
type PriorityQueue struct {
	items priorityQueueItems
}

// implements heap.Interface and holds Items.
type priorityQueueItems []*priorityQueueItem

// An priorityQueueItem is something we manage in a priority queue.
type priorityQueueItem struct {
	value    interface{} // The value of the item; arbitrary.
	priority int         // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// New returns a new PriorityQueue ready to be used
func New() *PriorityQueue {
	pq := &PriorityQueue{
		items: make(priorityQueueItems, 0),
	}
	heap.Init(&pq.items)

	return pq
}

// Push pushes a value into the queue with a priority
func (pq *PriorityQueue) Push(priority int, value interface{}) {
	item := &priorityQueueItem{
		value:    value,
		priority: priority,
	}

	heap.Push(&pq.items, item)
}

// Pop returns values sorted by priority, nil on empty
func (pq *PriorityQueue) Pop() interface{} {
	item := heap.Pop(&pq.items)
	if item == nil {
		return nil
	}
	pqi := item.(*priorityQueueItem)

	return pqi.value
}

func (items priorityQueueItems) Len() int { return len(items) }

func (items priorityQueueItems) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return items[i].priority > items[j].priority
}

func (items priorityQueueItems) Swap(i, j int) {
	// since we don't want to panic check boundaries first
	// priorityQueueItems is not meant to be reused
	if j <0 || i < 0 {
		return
	}
	items[i], items[j] = items[j], items[i]
	items[i].index = i
	items[j].index = j
}

func (items *priorityQueueItems) Push(x interface{}) {
	n := len(*items)
	item := x.(*priorityQueueItem)
	item.index = n
	*items = append(*items, item)
}

func (items *priorityQueueItems) Pop() interface{} {
	old := *items
	n := len(old)
	if n == 0 {
		return nil
	}
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*items = old[0 : n-1]

	return item
}
