package tree

import (
	"testing"
)

func TestNewTree(t *testing.T) {
	var tree BTree
	tree.Insert(10)
	if tree.root == nil {
		t.Fatal("should not be nil")
	}
	if tree.root.value != 10 {
		t.Fatal("should be value of ten")
	}
}

func TestTreeSize(t *testing.T) {
	var tree BTree
	if tree.Size() != 0 {
		t.Fatal("should be size of zero")
	}
	tree.Insert(1)
	if tree.Size() != 1 {
		t.Fatal("should be size of one")
	}
	tree.Insert(1)
	tree.Insert(1)
	if tree.Size() != 3 {
		t.Fatal("should be size of three")
	}
}

func TestTreeHeight(t *testing.T) {
	var tree BTree
	if tree.Height() > 0 {
		t.Fatal("should have a height of zero")
	}
	tree.Insert(1000)
	if tree.Height() != 1 {
		t.Fatal("should have a height of one")
	}
	tree.Insert(1)
	tree.Insert(10000)
	if tree.Height() != 2 {
		t.Fatal("should have a height of two")
	}

	tree.Insert(2)
	if tree.Height() != 3 {
		t.Fatal("should hav a height of three")
	}

}

func TestTreeFind(t *testing.T) {
	var tree BTree

	tree.Insert(7)
	tree.Insert(3)
	tree.Insert(11)
	tree.Insert(5)
	tree.Insert(9)
	tree.Insert(1)
	tree.Insert(13)
	tree.Insert(8)
	tree.Insert(6)
	tree.Insert(12)
	tree.Insert(4)
	tree.Insert(14)

	v, err := tree.Find(11)

	if err != nil || v != 11 {
		t.Fatal(err)
	}

	v, err = tree.Find(1)

	if err != nil || v != 1 {
		t.Fatal(err)
	}

	v, err = tree.Find(99)
	if err == nil || v == 99 {
		t.Fatal(err)
	}
}
