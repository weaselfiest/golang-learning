package linkedlist

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func captureOutput(f func()) string {
	r, w, _ := os.Pipe()
	old := os.Stdout
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestPrintEmpty(t *testing.T) {
	ll := LinkedList[int]{}
	out := captureOutput(ll.Print)
	if out != "nil\n" {
		t.Errorf("got %q, want %q", out, "nil\n")
	}
}

func TestInsertAndPrint(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)

	out := captureOutput(ll.Print)
	want := "1 -> 2 -> 3 -> nil\n"
	if out != want {
		t.Errorf("got %q, want %q", out, want)
	}
}

func TestInsertStrings(t *testing.T) {
	ll := LinkedList[string]{}
	ll.Insert("a")
	ll.Insert("b")

	out := captureOutput(ll.Print)
	want := "a -> b -> nil\n"
	if out != want {
		t.Errorf("got %q, want %q", out, want)
	}
}

func TestSearchFound(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Insert(10)
	ll.Insert(20)
	ll.Insert(30)

	if !ll.Search(20) {
		t.Error("expected to find 20")
	}
}

func TestSearchNotFound(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Insert(10)

	if ll.Search(99) {
		t.Error("expected not to find 99")
	}
}

func TestSearchEmpty(t *testing.T) {
	ll := LinkedList[int]{}
	if ll.Search(1) {
		t.Error("search in empty list should return false")
	}
}

func TestDeleteMiddle(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Delete(2)

	out := captureOutput(ll.Print)
	want := "1 -> 3 -> nil\n"
	if out != want {
		t.Errorf("got %q, want %q", out, want)
	}
}

func TestDeleteHead(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Insert(1)
	ll.Insert(2)
	ll.Delete(1)

	out := captureOutput(ll.Print)
	want := "2 -> nil\n"
	if out != want {
		t.Errorf("got %q, want %q", out, want)
	}
}

func TestDeleteTail(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Delete(3)

	out := captureOutput(ll.Print)
	want := "1 -> 2 -> nil\n"
	if out != want {
		t.Errorf("got %q, want %q", out, want)
	}
}

func TestDeleteOnlyElement(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Insert(42)
	deleted := ll.Delete(42)

	if !deleted {
		t.Error("expected Delete to return true")
	}
	out := captureOutput(ll.Print)
	if out != "nil\n" {
		t.Errorf("got %q, want %q", out, "nil\n")
	}
}

func TestDeleteNotFound(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Insert(1)
	ll.Insert(2)

	deleted := ll.Delete(99)
	if deleted {
		t.Error("expected Delete to return false for missing element")
	}
}

func TestDeleteDuplicates(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(2)
	ll.Insert(3)
	ll.Delete(2)

	out := captureOutput(ll.Print)
	want := "1 -> 2 -> 3 -> nil\n"
	if out != want {
		t.Errorf("got %q, want %q", out, want)
	}
}

func TestInsertAfterDelete(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Insert(1)
	ll.Insert(2)
	ll.Delete(1)
	ll.Insert(3)

	out := captureOutput(ll.Print)
	want := "2 -> 3 -> nil\n"
	if out != want {
		t.Errorf("got %q, want %q", out, want)
	}
}

func TestSearchAfterDelete(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Insert(5)
	ll.Insert(10)
	ll.Delete(5)

	if ll.Search(5) {
		t.Error("deleted element should not be found")
	}
	if !ll.Search(10) {
		t.Error("remaining element should be found")
	}
}

func TestPrintSingleElement(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Insert(7)

	out := captureOutput(ll.Print)
	want := "7 -> nil\n"
	if out != want {
		t.Errorf("got %q, want %q", out, want)
	}
}
