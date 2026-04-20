package queue

type Queue[T any] struct {
	data  []T
	head  int
	tail  int
	count int
}

func (q *Queue[T]) Enqueue(v T) {
	if q.count == len(q.data) {
		q.grow()
	}
	q.data[q.tail] = v
	q.tail = (q.tail + 1) % len(q.data)
	q.count++
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	v := q.data[q.head]
	var zero T
	q.data[q.head] = zero
	q.head = (q.head + 1) % len(q.data)
	q.count--
	return v, true
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	return q.data[q.head], true
}

func (q *Queue[T]) IsEmpty() bool {
	return q.count == 0
}

func (q *Queue[T]) grow() {
	newCap := 8
	if len(q.data) > 0 {
		newCap = len(q.data) * 2
	}
	newData := make([]T, newCap)
	for i := 0; i < q.count; i++ {
		newData[i] = q.data[(q.head+i)%len(q.data)]
	}
	q.data = newData
	q.head = 0
	q.tail = q.count
}
