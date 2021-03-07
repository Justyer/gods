package queue

type Queue interface {
	Push(elem interface{})
	Pop() (elem interface{}, err error)
}

const (
	TypeLinkedQueue int32 = 1
	TypeArrayQueue  int32 = 2
)

func NewQueue(tp int32) Queue {
	switch tp {
	case TypeLinkedQueue:
		return &LinkedQueue{}
	case TypeArrayQueue:
		return &ArrayQueue{}
	default:
		return &LinkedQueue{}
	}
}
