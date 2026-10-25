package deque

import ()

type Deque[T any] struct {
	elements []T
}

func New[T any]() *Deque[T] {
	d := new(Deque[T])
	return d
}

func (d *Deque[T]) IsEmpty() bool {
	return len(d.elements) < 1
}

func (d *Deque[T]) Size() int {
	return len(d.elements)
}

func (d *Deque[T]) PushFront(data T) {
	d.elements = append([]T{data}, d.elements...)
}

func (d *Deque[T]) PeekFront() (T, bool) {
	if d.IsEmpty() {
		var null T
		return null, false
	}
	return d.elements[0], true
}

func (d *Deque[T]) PopFront() (T, bool) {
	if d.IsEmpty() {
		var null T
		return null, false
	}
	var null T
	e := d.elements[0]
	d.elements[0] = null
	d.elements = d.elements[1:]
	return e, true
}

func (d *Deque[T]) PushBack(data T) {
	d.elements = append(d.elements, data)
}

func (d *Deque[T]) PeekBack() (T, bool) {
	if d.IsEmpty() {
		var null T
		return null, false
	}
	idx := len(d.elements) - 1
	return d.elements[idx], true
}

func (d *Deque[T]) PopBack() (T, bool) {
	if d.IsEmpty() {
		var null T
		return null, false
	}
	var null T
	idx := len(d.elements) - 1
	e := d.elements[idx]
	d.elements[idx] = null
	d.elements = d.elements[:idx]
	return e, true
}
