package queue

type SliceQueue struct {
	head, tail int64
	list       []interface{}
}

func (q *SliceQueue) Push(val interface{}) {

}

func (q *SliceQueue) Pop() (val interface{}, err error) {
	return
}
