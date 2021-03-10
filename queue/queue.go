package queue

type Queue interface {
	Push(elem interface{})
	Pop() (elem interface{}, err error)
}

const (
	TypeLinkedQueue int32 = 1
	TypeSliceQueue  int32 = 2
)

func NewQueue(tp int32) Queue {
	switch tp {
	case TypeLinkedQueue:
		return &LinkedQueue{}
	case TypeSliceQueue:
		return &SliceQueue{}
	default:
		return &LinkedQueue{}
	}
}
