package misc

import (
	"errors"
)

type QueueUtil[T any] struct {
	head   *ListNode[T]
	tail   *ListNode[T]
	length int
}

func (q *QueueUtil[T]) Length() int {
	return q.length
}

func (q *QueueUtil[T]) Enqueue(item T) {
	newNode := &ListNode[T]{Val: item}

	if q.head == nil {
		q.head = newNode
		q.tail = newNode
		q.length = 1
	} else {
		q.tail.Next = newNode
		q.tail = newNode
		q.length++
	}
}

func (q *QueueUtil[T]) Dequeue() (*T, error) {
	if q.length <= 0 {
		return nil, errors.New("nothing to dequeue")
	}

	ret := &q.head.Val
	q.head = q.head.Next

	q.length--

	return ret, nil
}

func (q *QueueUtil[T]) Peek() (*T, error) {
	if q.length <= 0 {
		return nil, errors.New("nothing to peek at")
	}

	return &q.head.Val, nil
}
