package das

import (
	"errors"
	"fmt"
)

type SingleLinkedListNode[T any] struct {
	value T
	next  *SingleLinkedListNode[T]
}

type SingleLinkedList[T any] struct {
	head *SingleLinkedListNode[T]
}

type QueueNode[T any] struct {
	value T
	next  *QueueNode[T]
}

type Queue[T any] struct {
	head *QueueNode[T]
	tail *QueueNode[T]
}

var ErrEmptyQueue = errors.New("empty queue")

func NewQueue[T any](elements []T) (new_queue Queue[T]) {
	for _, value := range elements {
		new_queue.enqueue(value)
	}
	return
}

func (q *Queue[T]) deque() (result T, err error) {
	if q.head == nil {
		return result, ErrEmptyQueue
	}
	result = q.head.value
	q.head = q.head.next
	return
}

func (q *Queue[T]) enqueue(value T) {
	next_element := &QueueNode[T]{value: value, next: nil}
	if q.tail == nil {
		q.head = next_element
		q.tail = q.head
	} else {
		q.tail.next = next_element
		q.tail = next_element
	}

}

func (q Queue[T]) String() string {
	var result string
	node := q.head
	if node == nil {
		return result
	}
	for ; node.next != nil; node = node.next {
		result += fmt.Sprintf("%v->", node.value)
	}
	result += fmt.Sprintf("%v", node.value)
	return result
}

type StackNode[T any] struct {
	value T
	prev  *StackNode[T]
}

type Stack[T any] struct {
	top *StackNode[T]
}

var ErrEmptyStack = errors.New("empty stack")

func (s *Stack[T]) push(value T) {
	node := &StackNode[T]{value: value, prev: nil}
	if s.top != nil {
		node.prev = s.top
	}
	s.top = node
}

func (s *Stack[T]) pop() (result T, err error) {
	if s.top == nil {
		return result, ErrEmptyStack
	}
	result = s.top.value
	s.top = s.top.prev
	return
}

func (s Stack[T]) String() string {
	var result string
	node := s.top
	if node == nil {
		return result
	}
	for ; node.prev != nil; node = node.prev {
		result += fmt.Sprintf("%v->", node.value)
	}
	result += fmt.Sprintf("%v", node.value)
	return result
}
