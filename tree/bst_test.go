package tree

import "testing"

func TestBST(t *testing.T) {
	bst := &BST{}

	elems := []int64{1, 2, 3, 4, 1}
	for _, elem := range elems {
		err := bst.Add(elem)
		if err != nil {
			t.Fatal(err)
		}
	}

	t.Log(bst.String())
}
