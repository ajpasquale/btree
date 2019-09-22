package tree

import "fmt"

// BTree represents the entire binary tree structure
type BTree struct {
	root *node
}

// node is a structure to hold a binary tree
type node struct {
	value int
	left  *node
	right *node
}

// Insert adds a new node to the tree with the given value.
func (b *BTree) Insert(v int) {
	if b.root == nil {
		b.root = &node{
			value: v,
			left:  nil,
			right: nil,
		}
	} else {
		b.root.insert(v)
	}
}

func (n *node) insert(v int) {
	if n.value < v {
		if n.left == nil {
			n.left = &node{
				value: v,
				left:  nil,
				right: nil,
			}
		} else {
			n.left.insert(v)
		}
	} else {
		if n.right == nil {
			n.right = &node{
				value: v,
				left:  nil,
				right: nil,
			}
		} else {
			n.right.insert(v)
		}
	}
}

// Size returns the total number of nodes in the tree.
func (b *BTree) Size() int {
	if b.root == nil {
		return 0
	}
	return b.root.size()
}

func (n *node) size() int {
	if n == nil {
		return 0
	}
	return 1 + n.left.size() + n.right.size()
}

// Height returns the height of the tree.
func (b BTree) Height() int {
	return b.root.height()
}

func (n *node) height() int {
	if n == nil {
		return 0
	}
	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}
	return max(n.left.height(), n.right.height()) + 1
}

// Find will return a value in the tree
func (b *BTree) Find(v int) (int, error) {
	n := b.root
	for {
		if n == nil {
			break
		}
		if n.value < v {
			n = n.left
		} else if n.value > v {
			n = n.right
		} else {
			return n.value, nil
		}
	}
	return 0, fmt.Errorf("cannot find %d", v)
}
