package queue

import (
	"errors"
	"sync"
)

var errQueueEmpty = errors.New("queue is empty.")

type node struct {
	val  interface{}
	next *node
}

type LinkedQueue struct {
	nelem int64
	head  *node
	lock  sync.Mutex
}

func (q *LinkedQueue) Push(val interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	nd := &node{val: val}
	if q.head == nil {
		q.head = nd
		return
	}
	var p *node
	for p = q.head; p.next != nil; p = q.head.next {
	}
	p.next = nd
}

func (q *LinkedQueue) Pop() (val interface{}, err error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.head == nil {
		return nil, errQueueEmpty
	}
	val = q.head.val
	q.head = q.head.next
	return
}
