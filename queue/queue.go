package main

import "errors"

var ErrEmptyQueue = errors.New("Empty")

type QueueNode struct {
	value int
	next  *QueueNode
}

type Queue struct {
	length int
	head   *QueueNode
	tail   *QueueNode
}

func (q *Queue) enqueue(item int) {
    node := &QueueNode{value: item, next: nil}
    q.length++
    if q.head != nil {
        q.tail.next = node
        q.tail = q.tail.next
        return
    }

    q.head = node
    q.tail = node
}

func (q *Queue) deque() (int, error) {
    if q.head != nil {
        q.length--
        value := q.head
        q.head = q.head.next
        return value.value, nil
    }

    return 0, ErrEmptyQueue 
}

func (q *Queue) peek() (int, error) {
    if q.head != nil {
        return q.head.value, nil
    }

    return 0, ErrEmptyQueue
}
