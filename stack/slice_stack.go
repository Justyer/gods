package stack

import (
	"errors"
	"sync/atomic"
)

var errStackEmpty = errors.New("stack is empty.")

type SliceStack struct {
	nelem int64
	elems []interface{}
}

func (s *SliceStack) Push(elem interface{}) {
	s.elems = append(s.elems, elem)
	atomic.AddInt64(&s.nelem, 1)
}

func (s *SliceStack) Pop() (elem interface{}, err error) {
	if s.nelem == 0 {
		return nil, errStackEmpty
	}
	elem = s.elems[s.nelem-1]
	atomic.AddInt64(&s.nelem, -1)
	return elem, nil
}
