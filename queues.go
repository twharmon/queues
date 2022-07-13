package queues

import (
	"fmt"
	"sync"
)

type Queue[T any] struct {
	mu    sync.Mutex
	items []T
	pos   int
}

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(item ...T) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.pos > 0 && cap(q.items) == len(q.items) {
		q.items = q.items[q.pos:]
		q.pos = 0
	}
	q.items = append(q.items, item...)
}

func (q *Queue[T]) Dequeue() T {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.pos >= len(q.items) {
		var t T
		return t
	}
	item := q.items[q.pos]
	q.pos++
	return item
}

func (q *Queue[T]) Peek() T {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.pos >= len(q.items) {
		var t T
		return t
	}
	t := q.items[q.pos]
	return t
}

func (q *Queue[T]) Clear() {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.items = nil
	q.pos = 0
}

func (q *Queue[T]) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return len(q.items) - q.pos
}

func (q *Queue[T]) Slice() []T {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.items[q.pos:]
}

func (q *Queue[T]) String() string {
	q.mu.Lock()
	defer q.mu.Unlock()
	return fmt.Sprintf("%v", q.items[q.pos:])
}
