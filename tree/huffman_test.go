package tree

import "testing"

func TestHuffman(t *testing.T) {
	huff := Huffman{}
	s := "qwerqqqwwe"
	huff.Init([]byte(s))
}
