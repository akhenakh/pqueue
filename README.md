# Pqueue

Package pqueue provides a very simple priority queue
It's basically copy and paste from the heap package example, only to override the Push method.

```go
func (pq *PriorityQueue) Push(priority int, value interface{})
func (pq *PriorityQueue) Pop() interface{}
```