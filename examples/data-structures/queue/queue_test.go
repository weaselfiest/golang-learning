package queue

import "testing"

func TestIsEmptyOnNew(t *testing.T) {
	q := Queue[int]{}
	if !q.IsEmpty() {
		t.Fatal("new queue must be empty")
	}
}

func TestEnqueueMakesNonEmpty(t *testing.T) {
	q := Queue[string]{}
	q.Enqueue("hello")
	if q.IsEmpty() {
		t.Fatal("queue must not be empty after Enqueue")
	}
}

func TestDequeueOrder(t *testing.T) {
	q := Queue[int]{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	for _, want := range []int{1, 2, 3} {
		v, ok := q.Dequeue()
		if !ok {
			t.Fatalf("expected ok=true, got false")
		}
		if v != want {
			t.Fatalf("want %d, got %d", want, v)
		}
	}
}

func TestDequeueEmptyQueue(t *testing.T) {
	q := Queue[float64]{}
	v, ok := q.Dequeue()
	if ok {
		t.Fatal("expected ok=false on empty queue")
	}
	var zero float64
	if v != zero {
		t.Fatalf("expected zero value, got %v", v)
	}
}

func TestPeekDoesNotRemove(t *testing.T) {
	q := Queue[int]{}
	q.Enqueue(42)

	for i := 0; i < 3; i++ {
		v, ok := q.Peek()
		if !ok {
			t.Fatal("expected ok=true")
		}
		if v != 42 {
			t.Fatalf("want 42, got %d", v)
		}
	}

	if q.IsEmpty() {
		t.Fatal("Peek must not remove elements")
	}
}

func TestPeekEmptyQueue(t *testing.T) {
	q := Queue[int]{}
	v, ok := q.Peek()
	if ok {
		t.Fatal("expected ok=false on empty queue")
	}
	var zero int
	if v != zero {
		t.Fatalf("expected zero value, got %v", v)
	}
}

func TestIsEmptyAfterEnqueueDequeue(t *testing.T) {
	q := Queue[int]{}
	q.Enqueue(1)
	q.Dequeue()
	if !q.IsEmpty() {
		t.Fatal("queue must be empty after dequeuing last element")
	}
}
