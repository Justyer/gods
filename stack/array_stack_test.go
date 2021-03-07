package stack

import (
	"testing"
)

func TestArrayStack(t *testing.T) {
	s := &ArrayStack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
}
