// Package linkedlist contains implementation of doubly linked list.
package linkedlist

import "errors"

// Node type defines a node in a doubly linked list.
type Node struct {
	Value    any
	PrevNode *Node
	NextNode *Node
}

// List type defines a doubly linked list.
type List struct {
	Length int
	Head   *Node
	Tail   *Node
}

// errListEmpty is the error for an empty list.
var errListEmpty = errors.New("unable to remove item, list is empty")

// newNode creates a new node with passed value.
func newNode(v any) *Node {
	return &Node{Value: v}
}

// takes any number of elements and returns them in form of a list.
func NewList(elements ...any) *List {
	newList := new(List)
	for _, item := range elements {
		newList.Push(item)
	}
	return newList
}

// Next method returns the address of the next node.
func (n *Node) Next() *Node {
	return n.NextNode
}

// Prev method returns the address of the previous node.
func (n *Node) Prev() *Node {
	return n.PrevNode
}

// Unshift method inserts a value at the front of the list.
func (l *List) Unshift(v any) {
	newNode := newNode(v)
	l.Length++
	if l.Head == nil { // if the list is empty then new node will be both head and tail.
		l.Tail, l.Head = newNode, newNode
		return
	}
	// link the new node with the current head of the list.
	l.Head.PrevNode, newNode.NextNode = newNode, l.Head
	l.Head = newNode // assign new node as the head.
}

// Push method inserts a value at the back of the list
func (l *List) Push(v any) {
	newNode := newNode(v)
	l.Length++
	if l.Tail == nil {
		l.Tail, l.Head = newNode, newNode
		return
	}
	l.Tail.NextNode, newNode.PrevNode = newNode, l.Tail
	l.Tail = newNode
}

// Shift method removes a value from the front of the list.
func (l *List) Shift() (any, error) {
	if l.Length == 0 {
		return nil, errListEmpty // can't remove a value from an empty list.
	}
	l.Length--
	toReturn := l.Head.Value
	l.Head = l.Head.NextNode // assign the node after the current head as head.
	if l.Head == nil {       // if the new head is nil then list is now empty.
		l.Tail = nil // set tail as nil as well and return.
		return toReturn, nil
	}
	l.Head.PrevNode = nil // unlinked previous head from the list.
	return toReturn, nil
}

// Pop method removes a value from the back of the list.
func (l *List) Pop() (any, error) {
	if l.Length == 0 {
		return nil, errListEmpty
	}
	l.Length--
	toReturn := l.Tail.Value
	l.Tail = l.Tail.PrevNode
	if l.Tail == nil {
		l.Head = nil
		return toReturn, nil
	}
	l.Tail.NextNode = nil
	return toReturn, nil
}

// Reverse method reverses the linked list.
func (l *List) Reverse() {
	next := l.Head // start from head.
	for next != nil {
		// exchange the addresses of prev and next
		next.NextNode, next.PrevNode = next.PrevNode, next.NextNode
		next = next.PrevNode // move on to next node.
	}
	l.Head, l.Tail = l.Tail, l.Head // exchange the addresses of head and tail
}

// First method returns the head node of the list.
func (l *List) First() *Node {
	return l.Head
}

// Last method returns the tail node of the list.
func (l *List) Last() *Node {
	return l.Tail
}
