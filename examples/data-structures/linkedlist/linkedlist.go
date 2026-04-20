package linkedlist

import "fmt"

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
}

func (l *LinkedList[T]) Insert(val T) {
	node := &Node[T]{value: val}
	if l.head == nil {
		l.head = node
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = node
}

func (l *LinkedList[T]) Delete(val T) bool {
	if l.head == nil {
		return false
	}
	if l.head.value == val {
		l.head = l.head.next
		return true
	}
	prev := l.head
	current := l.head.next
	for current != nil {
		if current.value == val {
			prev.next = current.next
			return true
		}
		prev = current
		current = current.next
	}
	return false
}

func (l *LinkedList[T]) Search(val T) bool {
	current := l.head
	for current != nil {
		if current.value == val {
			return true
		}
		current = current.next
	}
	return false
}

func (l *LinkedList[T]) Print() {
	current := l.head
	for current != nil {
		fmt.Printf("%v -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}
