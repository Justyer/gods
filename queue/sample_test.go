package queue

import (
	"testing"
)

func TestQueueInterface(t *testing.T) {
	q := NewQueue(TypeLinkedQueue)
	q.Push(3)
	q.Push("ddd")
	t.Log(q.Pop())
	t.Log(q.Pop())
}
