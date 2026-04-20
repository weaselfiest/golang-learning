package stack

import "testing"

func TestIsEmptyOnNew(t *testing.T) {
	s := Stack[int]{}
	if !s.IsEmpty() {
		t.Fatal("new stack must be empty")
	}
}

func TestPushMakesNonEmpty(t *testing.T) {
	s := Stack[string]{}
	s.Push("hello")
	if s.IsEmpty() {
		t.Fatal("stack must not be empty after Push")
	}
}

func TestPopOrder(t *testing.T) {
	s := Stack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)

	for _, want := range []int{3, 2, 1} {
		v, ok := s.Pop()
		if !ok {
			t.Fatalf("expected ok=true, got false")
		}
		if v != want {
			t.Fatalf("want %d, got %d", want, v)
		}
	}
}

func TestPopEmptyStack(t *testing.T) {
	s := Stack[float64]{}
	v, ok := s.Pop()
	if ok {
		t.Fatal("expected ok=false on empty stack")
	}
	var zero float64
	if v != zero {
		t.Fatalf("expected zero value, got %v", v)
	}
}

func TestPeekDoesNotRemove(t *testing.T) {
	s := Stack[int]{}
	s.Push(42)

	for i := 0; i < 3; i++ {
		v, ok := s.Peek()
		if !ok {
			t.Fatal("expected ok=true")
		}
		if v != 42 {
			t.Fatalf("want 42, got %d", v)
		}
	}

	if s.IsEmpty() {
		t.Fatal("Peek must not remove elements")
	}
}

func TestPeekEmptyStack(t *testing.T) {
	s := Stack[int]{}
	v, ok := s.Peek()
	if ok {
		t.Fatal("expected ok=false on empty stack")
	}
	var zero int
	if v != zero {
		t.Fatalf("expected zero value, got %v", v)
	}
}

func TestIsEmptyAfterPushPop(t *testing.T) {
	s := Stack[int]{}
	s.Push(1)
	s.Pop()
	if !s.IsEmpty() {
		t.Fatal("stack must be empty after popping last element")
	}
}
