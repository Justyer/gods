package tree

import (
	"fmt"
	"testing"
)

func TestBST(t *testing.T) {
	node := &BinarySearchTree{}
	node.AddElem(3)
	node.AddElem(1)
	node.AddElem(2)
	node.AddElem(5)
	js := node.ShowJson()
	fmt.Println(js)
	fmt.Println(node.MinNode().Value)
	fmt.Println(node.MaxNode().Value)
	fmt.Println(node.Find(1))
	fmt.Println(node.Find(9))
}
