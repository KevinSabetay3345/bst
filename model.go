package bst

type BST struct {
	Root *Node
	Size int
}

type Node struct {
	Left  *Node
	Right *Node
	Value int
}
