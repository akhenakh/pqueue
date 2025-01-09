# pqueue

pqueue implements a generic priority queue data structure using a max-heap.

the priority queue maintains elements in descending order of priority, where the
element with the highest priority value is always at the front of the queue.
Each element consists of two components:
- A value of any type T
- A priority of any ordered type P (numbers, strings, etc.)

The implementation uses Go's container/heap package internally, providing O(log n)
time complexity for Push and Pop operations, and O(1) for Peek and Len operations.

Example usage:

```go
pq := pqueue.New[string, int]()
pq.Push("task1", 3)
pq.Push("task2", 1)
pq.Push("task3", 4)

value, priority, ok := pq.Pop()  // Returns "task3", 4, true
```
  

All operations are safe to use with empty queues, returning appropriate zero
values and a boolean false when attempting to access empty queues.

