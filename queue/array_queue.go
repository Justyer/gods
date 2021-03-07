package queue

type ArrayQueue struct {
	head, tail int64
	list       []interface{}
}

func (q *ArrayQueue) Push(val interface{}) {

}

func (q *ArrayQueue) Pop() (val interface{}, err error) {
	return
}
