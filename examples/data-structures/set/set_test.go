package set

import (
	"testing"
)

func TestAddAndContains(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(2)

	if !s.Contains(1) {
		t.Error("expected set to contain 1")
	}
	if !s.Contains(2) {
		t.Error("expected set to contain 2")
	}
	if s.Contains(3) {
		t.Error("expected set not to contain 3")
	}
}

func TestRemove(t *testing.T) {
	s := NewSet[string]()
	s.Add("a")
	s.Add("b")
	s.Remove("a")
	s.Remove("z") // несуществующий элемент — не должен паниковать

	if s.Contains("a") {
		t.Error("expected set not to contain 'a' after removal")
	}
	if !s.Contains("b") {
		t.Error("expected set to still contain 'b'")
	}
}

func TestUnion(t *testing.T) {
	a := NewSet[int]()
	a.Add(1)
	a.Add(2)

	b := NewSet[int]()
	b.Add(2)
	b.Add(3)

	u := a.Union(b)

	for _, v := range []int{1, 2, 3} {
		if !u.Contains(v) {
			t.Errorf("expected union to contain %d", v)
		}
	}

	// Исходные множества не изменились
	if a.Contains(3) {
		t.Error("Union must not modify receiver")
	}
	if b.Contains(1) {
		t.Error("Union must not modify argument")
	}
}

func TestIntersect(t *testing.T) {
	a := NewSet[int]()
	a.Add(1)
	a.Add(2)
	a.Add(3)

	b := NewSet[int]()
	b.Add(2)
	b.Add(3)
	b.Add(4)

	i := a.Intersect(b)

	if i.Contains(1) {
		t.Error("expected intersection not to contain 1")
	}
	if !i.Contains(2) {
		t.Error("expected intersection to contain 2")
	}
	if !i.Contains(3) {
		t.Error("expected intersection to contain 3")
	}
	if i.Contains(4) {
		t.Error("expected intersection not to contain 4")
	}

	// Исходные множества не изменились
	if !a.Contains(1) {
		t.Error("Intersect must not modify receiver")
	}
}

func TestEmptySets(t *testing.T) {
	a := NewSet[int]()
	b := NewSet[int]()

	u := a.Union(b)
	if u.Contains(0) {
		t.Error("union of empty sets should be empty")
	}

	i := a.Intersect(b)
	if i.Contains(0) {
		t.Error("intersection of empty sets should be empty")
	}
}
