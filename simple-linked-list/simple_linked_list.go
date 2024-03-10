// Package linkedlist contains solution for Simple Linked List exercise on Exercism.
package linkedlist

import (
	"errors"
)

var (
	errEmptyList = errors.New("list is empty")
)

// node defines a single node in the linked list.
type node struct {
	data int
	next *node
}

// List defines a singly linked list.
type List struct {
	Head   *node
	length int
}

// newNode creates a new node. its takes data for the
// node and the address to next node.
func newNode(data int, next *node) *node {
	return &node{data: data, next: next}
}

// New created a new singly linked list from the provided array.
func New(elements []int) *List {
	list := new(List)
	for _, item := range elements {
		list.Push(item)
	}
	return list
}

// Size method returns the size of the singly linked list.
func (l *List) Size() int {
	return l.length
}

// Push method pushes the provided element to the front singly linked list.
func (l *List) Push(element int) {
	l.Head = newNode(element, l.Head)
	l.length++

}

// Pop method pops the first element from the front of the singly linked list and returns it.
func (l *List) Pop() (int, error) {
	if l.length <= 0 {
		return 0, errEmptyList
	}
	toReturn := l.Head.data
	l.Head = l.Head.next
	l.length--
	return toReturn, nil
}

// Array returns the simply linked list in form of an array.
func (l *List) Array() []int {
	output := make([]int, l.length)
	temp := l.Head
	for i := l.length - 1; i >= 0; i-- {
		output[i] = temp.data
		temp = temp.next
	}
	return output
}

// Reverse method reverses the order of the simply linked list.
func (l *List) Reverse() *List {
	current := l.Head
	var prev *node
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	l.Head = prev
	return l
}
