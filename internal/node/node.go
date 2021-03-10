package node

type OneWay struct {
	Val  interface{}
	Next *OneWay
}

type TwoWay struct {
	Val  interface{}
	Prev *TwoWay
	Next *TwoWay
}

type BinaryTree struct {
	Val   int64
	Left  *BinaryTree
	Right *BinaryTree
}

type Graph struct {
	Val      interface{}
	InDegree []*Graph
	OutGraph []*Graph
}
