package tree

type TreeNode struct {
	Value int       `json:"v"`
	Left  *TreeNode `json:"l,omitempty"`
	Right *TreeNode `json:"r,omitempty"`
}
