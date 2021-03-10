package queue

import (
	"errors"
	"sync"

	"github.com/Justyer/gods/internal/node"
)

var errQueueEmpty = errors.New("queue is empty.")

type LinkedQueue struct {
	nelem int64
	head  *node.OneWay
	lock  sync.Mutex
}

func (q *LinkedQueue) Push(val interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	nd := &node.OneWay{Val: val}
	if q.head == nil {
		q.head = nd
		return
	}
	var p *node.OneWay
	for p = q.head; p.Next != nil; p = q.head.Next {
	}
	p.Next = nd
}

func (q *LinkedQueue) Pop() (val interface{}, err error) {
	q.lock.Lock()
	defer q.lock.Unlock()
	if q.head == nil {
		return nil, errQueueEmpty
	}
	val = q.head.Val
	q.head = q.head.Next
	return
}
