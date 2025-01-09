package pqueue

import (
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	tests := []struct {
		name string
		fn   func(t *testing.T)
	}{
		{"empty queue operations", testEmptyQueue},
		{"basic operations", testBasicOperations},
		{"float priorities", testFloatPriorities},
		{"equal priorities", testEqualPriorities},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.fn)
	}
}

func testEmptyQueue(t *testing.T) {
	pq := New[string, int]()

	if !pq.IsEmpty() {
		t.Error("new queue should be empty")
	}
	if pq.Len() != 0 {
		t.Error("new queue should have length 0")
	}
	if _, _, ok := pq.Peek(); ok {
		t.Error("peek on empty queue should return false")
	}
	if _, _, ok := pq.Pop(); ok {
		t.Error("pop on empty queue should return false")
	}
}

func testBasicOperations(t *testing.T) {
	pq := New[string, int]()

	// Test insertions
	inputs := []struct {
		value    string
		priority int
	}{
		{"low", 1},
		{"high", 3},
		{"medium", 2},
	}

	for _, input := range inputs {
		pq.Push(input.value, input.priority)
	}

	if pq.Len() != 3 {
		t.Errorf("expected length 3, got %d", pq.Len())
	}

	// Test peek
	val, pri, ok := pq.Peek()
	if !ok || val != "high" || pri != 3 {
		t.Errorf("peek: got (%v, %v, %v), want (high, 3, true)", val, pri, ok)
	}

	// Test pop order
	expected := []struct {
		value    string
		priority int
	}{
		{"high", 3},
		{"medium", 2},
		{"low", 1},
	}

	for i, exp := range expected {
		val, pri, ok := pq.Pop()
		if !ok {
			t.Fatalf("pop %d: unexpected empty queue", i)
		}
		if val != exp.value || pri != exp.priority {
			t.Errorf("pop %d: got (%v, %v), want (%v, %v)",
				i, val, pri, exp.value, exp.priority)
		}
	}

	if !pq.IsEmpty() {
		t.Error("queue should be empty after popping all elements")
	}
}

func testFloatPriorities(t *testing.T) {
	pq := New[int, float64]()

	inputs := []struct {
		value    int
		priority float64
	}{
		{1, 1.5},
		{2, 2.5},
		{3, 0.5},
	}

	for _, input := range inputs {
		pq.Push(input.value, input.priority)
	}

	expected := []struct {
		value    int
		priority float64
	}{
		{2, 2.5},
		{1, 1.5},
		{3, 0.5},
	}

	for i, exp := range expected {
		val, pri, ok := pq.Pop()
		if !ok {
			t.Fatalf("pop %d: unexpected empty queue", i)
		}
		if val != exp.value || pri != exp.priority {
			t.Errorf("pop %d: got (%v, %v), want (%v, %v)",
				i, val, pri, exp.value, exp.priority)
		}
	}
}

func testEqualPriorities(t *testing.T) {
	pq := New[string, int]()

	// Push several items with the same priority
	inputs := []string{"first", "second", "third"}
	for _, v := range inputs {
		pq.Push(v, 1) // Same priority for all items
	}

	// Verify that we can pop all items
	seen := make(map[string]bool)
	for i := 0; i < len(inputs); i++ {
		val, pri, ok := pq.Pop()
		if !ok {
			t.Fatal("unexpected empty queue")
		}
		if pri != 1 {
			t.Errorf("expected priority 1, got %v", pri)
		}
		// Just verify that we get each item exactly once
		if seen[val] {
			t.Errorf("got duplicate value: %v", val)
		}
		seen[val] = true
	}

	// Verify we got all items
	for _, v := range inputs {
		if !seen[v] {
			t.Errorf("missing value: %v", v)
		}
	}
}
