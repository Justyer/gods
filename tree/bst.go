package tree

import (
	"encoding/json"
	"fmt"
)

type BinarySearchTree struct {
	Root *TreeNode
}

// 添加节点
func (n *BinarySearchTree) AddElem(value int) {
	if n.Root == nil {
		n.Root = &TreeNode{Value: value}
		return
	}
	node := n.Root
	for {
		if value < node.Value {
			if node.Left == nil {
				node.Left = &TreeNode{Value: value}
				return
			} else {
				node = node.Left
			}
		} else {
			if node.Right == nil {
				node.Right = &TreeNode{Value: value}
				return
			} else {
				node = node.Left
			}
		}
	}
}

// Json形式显示树状结构
func (n *BinarySearchTree) ShowJson() string {
	b, _ := json.Marshal(&n.Root)
	return string(b)
}

// 树中最大元素节点
func (n *BinarySearchTree) MaxNode() *TreeNode {
	node := n.Root
	for node.Right != nil {
		node = node.Right
	}
	return node
}

// 树中最小元素节点
func (n *BinarySearchTree) MinNode() *TreeNode {
	node := n.Root
	for node.Left != nil {
		node = node.Left
	}
	return node
}

// 查找目标值是否存在
// TreeNode if vlaue_exist else nil
func (n *BinarySearchTree) Find(value int) *TreeNode {
	node := n.Root
	for node != nil {
		switch {
		case value == node.Value:
			return node
		case value < node.Value:
			node = node.Left
		case node.Value < value:
			node = node.Right
		}
	}
	return node
}

// 前序遍历
func (n *BinarySearchTree) PreOrder() {
	n.preOrderByRecursion(n.Root)
}

// 中序遍历
func (n *BinarySearchTree) InOrder() {
	n.inOrderByRecursion(n.Root)
}

// 后序遍历
func (n *BinarySearchTree) PostOrder() {
	n.postOrderByRecursion(n.Root)
}

// 层序遍历
func (n *BinarySearchTree) LevelOrder() {
	if n.Root == nil {
		return
	}
	node := n.Root
	var nodeQueue []*TreeNode
	nodeQueue = append(nodeQueue, node)
	for ql := len(nodeQueue); ql > 0; ql = len(nodeQueue) {
		nq := nodeQueue[0]
		nodeQueue = append(nodeQueue[1:])
		fmt.Println(nq.Value)
		if nq.Left != nil {
			nodeQueue = append(nodeQueue, nq.Left)
		}
		if nq.Right != nil {
			nodeQueue = append(nodeQueue, nq.Right)
		}
	}
}

// 前序遍历递归实现
func (n *BinarySearchTree) preOrderByRecursion(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Value)
	n.inOrderByRecursion(node.Left)
	n.inOrderByRecursion(node.Right)
}

// 中序遍历递归实现
func (n *BinarySearchTree) inOrderByRecursion(node *TreeNode) {
	if node == nil {
		return
	}
	n.inOrderByRecursion(node.Left)
	fmt.Println(node.Value)
	n.inOrderByRecursion(node.Right)
}

// 后序遍历递归实现
func (n *BinarySearchTree) postOrderByRecursion(node *TreeNode) {
	if node == nil {
		return
	}
	n.inOrderByRecursion(node.Left)
	n.inOrderByRecursion(node.Right)
	fmt.Println(node.Value)
}
