package tree

import (
	"encoding/json"
	"errors"

	"github.com/Justyer/gods/internal/node"
)

var (
	errElemRepeat = errors.New("elem is already exist.")
)

// BST is BinarySearchTree.
type BST struct {
	root *node.BinaryTree
}

func (b *BST) Add(val int64) (err error) {
	if b.root == nil {
		b.root = &node.BinaryTree{Val: val}
		return
	}
	p := b.root
	for {
		switch {
		case val < p.Val:
			if p.Left != nil {
				p = p.Left
				continue
			} else {
				p.Left = &node.BinaryTree{Val: val}
				return
			}
		case p.Val < val:
			if p.Right != nil {
				p = p.Right
				continue
			} else {
				p.Right = &node.BinaryTree{Val: val}
				return
			}
		default:
			return errElemRepeat
		}
	}
}

func (b *BST) String() string {
	bb, _ := json.Marshal(b.root)
	return string(bb)
}
