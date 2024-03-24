// Package binarysearchtree contains tools for binary search tree
package binarysearchtree

// BinarySearchTree represents a binary search tree
type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

// NewBst creates and returns a new BinarySearchTree.
func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{data: i}
}

// Insert inserts  an int into the BinarySearchTree. insertion happen based
// on the rules of a binary search tree
func (bst *BinarySearchTree) Insert(i int) {
	if i <= bst.data {
		if bst.left == nil {
			bst.left = NewBst(i)
		} else {
			bst.left.Insert(i)
		}
	}
	if i > bst.data {
		if bst.right == nil {
			bst.right = NewBst(i)
		} else {
			bst.right.Insert(i)
		}
	}
}

// SortedData returns the ordered contents of BinarySearchTree as an []int.
// The values are in increasing order starting with the lowest int value.
func (bst *BinarySearchTree) SortedData() []int {
	// we use in order traversal here
	output := make([]int, 0)
	if bst.left != nil {
		output = append(output, bst.left.SortedData()...)
	}
	output = append(output, bst.data)
	if bst.right != nil {
		output = append(output, bst.right.SortedData()...)
	}

	return output
}
