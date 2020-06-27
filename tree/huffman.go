package tree

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

type HuffmanNode struct {
	L    *HuffmanNode `json:"L"`
	R    *HuffmanNode `json:"R"`
	V    uint         `json:"V"`
	B    byte         `json:"B"`
	HasB bool         `json:"-"`
}

func (n *HuffmanNode) String() string {
	return fmt.Sprintf("<B:%v - V:%v, L: %v R: %v>", n.B, n.V, n.L, n.R)
}

func MidOrder(node *HuffmanNode) {
	if node == nil {
		return
	}
	MidOrder(node.L)
	fmt.Println(node.String())
	MidOrder(node.R)
}

type HuffmanNodeList []*HuffmanNode

func (n HuffmanNodeList) Print() {
	var p []string
	for _, node := range n {
		t := fmt.Sprintf("<B:%v - V:%v>", node.B, node.V)
		p = append(p, t)
	}
	fmt.Println(strings.Join(p, ","))
}

func (n HuffmanNodeList) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func (n HuffmanNodeList) Len() int {
	return len(n)
}
func (n HuffmanNodeList) Less(i, j int) bool {
	return n[i].V > n[j].V
}

type Huffman struct{}

//
func (h *Huffman) Init(bs []byte) {
	mp := make(map[byte]uint)
	for _, b := range bs {
		mp[b]++
	}
	var nList HuffmanNodeList
	for k, v := range mp {
		fmt.Println(k, v)
		nList = append(nList, &HuffmanNode{B: k, V: v, HasB: true})
	}
	sort.Sort(nList)
	nList.Print()
	for nl := len(nList); nl > 1; nl = len(nList) {
		n1, n2 := nList[nl-1], nList[nl-2]
		node := &HuffmanNode{
			L: n1,
			R: n2,
			V: n1.V + n2.V,
		}
		nList = append(nList[:nl-2], node)
		sort.Sort(nList)
		nList.Print()
	}

	root := nList[0]

	fmt.Println("----")
	b, _ := json.Marshal(root)
	fmt.Println(string(b))

}

func FirstOrder(root *HuffmanNode) {
	var nodeList []*HuffmanNode

	for {
		if root.L != nil {
			nodeList = append(nodeList, root.L)
		}
		if root.R != nil {
			nodeList = append(nodeList, root.R)
		}

	}
}

func (h *Huffman) Encode() {

}

func (h *Huffman) Decode() {

}
