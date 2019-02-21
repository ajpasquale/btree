package tree

import (
	"fmt"
)

//Node is a structure to hold a binary tree
type Node struct {
	value int
	left  *Node
	right *Node
}

// Insert into the binary tree
func Insert(root *Node, v int) *Node {
	if root == nil {
		root = &Node{
			value: v,
			left:  nil,
			right: nil,
		}
	} else if root.value > v {
		root.left = Insert(root.left, v)
	} else {
		root.right = Insert(root.right, v)
	}
	return root
}

//Traverse the tree
func Traverse(root *Node) {
	if root == nil {
		return
	}
	Traverse(root.left)
	Traverse(root.right)
}

//Size returns (number of nodes) in the tree
func Size(root *Node) int {
	if root == nil {
		return 0
	}
	return 1 + Size(root.left) + Size(root.right)
}

//Height returns the height of the tree
func Height(root *Node) int {
	if root == nil {
		return 0
	}
	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}
	return max(Height(root.left), Height(root.right)) + 1
}

//Find will return a value in the tree
func Find(root *Node, v int) int {
	for {
		if root == nil {
			break
		}
		if root.value > v {
			fmt.Printf("Value: %d is less than root value: %d. Going left.\n", v, root.value)
			root = root.left
		} else if root.value < v {
			fmt.Printf("Value: %d is greater than root value: %d. Going right.\n", v, root.value)
			root = root.right
		} else {
			return root.value
		}
	}
	return 0
}
