package bst

import (
	"errors"
)

func (bst *BST) Insert(value int) error {
	if bst.Root == nil {
		bst.Root = &Node{Value: value}
		bst.Size++

		return nil
	}

	node := bst.Root
	for {
		if node.Value == value {
			return errors.New("value already exists")
		}

		if node.Value < value {
			if node.Right == nil {
				break
			}
			node = node.Right
		} else {
			if node.Left == nil {
				break
			}
			node = node.Left
		}
	}

	if node.Value < value {
		node.Right = &Node{Value: value}
	} else {
		node.Left = &Node{Value: value}
	}

	bst.Size++

	return nil
}

func (bst *BST) Exists(value int) bool {
	if bst.Root == nil {
		return false
	}

	node := bst.Root
	for node != nil {
		if node.Value == value {
			return true
		}

		if node.Value < value {
			node = node.Right
		} else {
			node = node.Left
		}
	}

	return false

}

func (bst *BST) Minimum() (int, error) {
	if bst.Root == nil {
		return 0, errors.New("there are no elements")
	}

	node := bst.Root
	for node.Left != nil {
		node = node.Left
	}

	return node.Value, nil

}

func (bst *BST) Maximum() (int, error) {
	if bst.Root == nil {
		return 0, errors.New("there are no elements")
	}

	node := bst.Root
	for node.Right != nil {
		node = node.Right
	}

	return node.Value, nil

}

func (bst *BST) Inorder() []int {
	return inorderAux(bst.Root)
}

func inorderAux(node *Node) []int {
	if node == nil {
		return []int{}
	}

	var result []int
	result = append(result, inorderAux(node.Left)...)
	result = append(result, node.Value)
	result = append(result, inorderAux(node.Right)...)

	return result
}

func (bst *BST) Next(value int) (int, error) {
	if bst.Root == nil {
		return 0, errors.New("element not found")
	}

	orderedValues := bst.Inorder()

	//binary search
	start := 0
	end := len(orderedValues)
	for start <= end {
		middle := start + (end-start)/2
		if orderedValues[middle] == value {
			if len(orderedValues) <= middle+1 {
				return 0, errors.New("there is no next element")
			}
			return orderedValues[middle+1], nil
		}

		if orderedValues[middle] < value {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}

	return 0, errors.New("element not found")
}

func (bst *BST) Delete(value int) error {

	if bst.Root == nil {
		return errors.New("value not found")
	}

	node := bst.Root
	if bst.Root.Value == value {
		bst.Size--
		replaceNode := deleteNode(bst.Root)
		bst.Root = replaceNode
		return nil
	}

	for node != nil {
		if node.Right != nil && node.Right.Value == value {
			replaceNode := deleteNode(node.Right)
			node.Right = replaceNode
			bst.Size--
			break
		}

		if node.Left != nil && node.Left.Value == value {
			replaceNode := deleteNode(node.Left)
			node.Left = replaceNode
			bst.Size--
			break
		}

		if node.Value < value {

			if node.Right == nil {
				return errors.New("value not found")
			}
			node = node.Right

		} else {

			if node.Left == nil {
				return errors.New("value not found")
			}
			node = node.Left

		}
	}

	return nil

}

func deleteNode(nodeToDelete *Node) *Node {
	if nodeToDelete.Left == nil && nodeToDelete.Right == nil {
		return nil
	}

	if nodeToDelete.Left != nil && nodeToDelete.Right == nil {
		return nodeToDelete.Left
	}

	if nodeToDelete.Left == nil && nodeToDelete.Right != nil {
		return nodeToDelete.Right
	}

	return deleteWith2Childs(nodeToDelete)
}

func deleteWith2Childs(nodeToDelete *Node) *Node {
	var parent *Node
	node := nodeToDelete.Left
	for node.Right != nil {
		parent = node
		node = node.Right
	}

	if parent != nil {
		if node.Left != nil {
			parent.Right = node.Left
		} else {
			parent.Right = nil
		}
		node.Left = nodeToDelete.Left
	}
	node.Right = nodeToDelete.Right

	return node
}
