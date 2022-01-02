package pqueue

import (
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	q := New()

	q.Push(0, "c")
	q.Push(5, "a")
	q.Push(1, "b")

	if v := q.Pop();v != "a" {
		t.Fatalf("we are expecting sorted by priority got: %s", v)
	}

	if v := q.Pop();v != "b" {
		t.Fatalf("we are expecting sorted by priority got: %s", v)
	}

	if v := q.Pop();v != "c" {
		t.Fatalf("we are expecting sorted by priority got: %s", v)
	}

	if v := q.Pop();v != nil {
		t.Fatalf("we are expecting nil value got: %s", v)
	}
}
