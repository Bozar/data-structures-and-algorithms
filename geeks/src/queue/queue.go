package queue

import ()

type Queue[T any] struct {
	elements []T
}

func New[T any]() *Queue[T] {
	return new(Queue[T])
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) < 1
}

func (q *Queue[T]) Size() int {
	return len(q.elements)
}

func (q *Queue[T]) Enqueue(data T) {
	q.elements = append(q.elements, data)
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.IsEmpty() {
		var null T
		return null, false
	}
	return q.elements[0], true
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var null T
		return null, false
	}

	var null T
	data := q.elements[0]
	q.elements[0] = null
	q.elements = q.elements[1:]
	return data, true
}
