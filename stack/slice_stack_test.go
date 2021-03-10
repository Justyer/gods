package stack

import (
	"testing"
)

func TestSliceStack(t *testing.T) {
	s := &SliceStack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
}
