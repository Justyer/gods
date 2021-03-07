package queue

import (
	"testing"
)

func TestLinkedQueue(t *testing.T) {
	q := &LinkedQueue{}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	t.Log(q.Pop())
	t.Log(q.Pop())
	t.Log(q.Pop())
}
