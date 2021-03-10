package heap

type Heap struct {
	nelem int64
	elems []interface{}
}

func (h *Heap) Add(elem interface{})
