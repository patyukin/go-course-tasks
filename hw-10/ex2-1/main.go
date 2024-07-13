// https://codereview.stackexchange.com/questions/188305/2-3-tree-in-python
package main

import (
	"fmt"
)

type Node struct {
	nodeType       int
	d1, d2, d3     *int
	c1, c2, c3, c4 *Node
	parent         *Node
}

func NewNode(data int, parent *Node) *Node {
	d1 := data
	return &Node{
		nodeType: 2,
		d1:       &d1,
		parent:   parent,
	}
}

func (n *Node) Push(data int) {
	if n.nodeType == 2 {
		n.nodeType = 3
		if data < *n.d1 {
			n.d2 = n.d1
			n.d1 = &data
		} else {
			n.d2 = &data
		}
	} else if n.nodeType == 3 {
		n.nodeType = 4
		if data < *n.d1 {
			n.d3 = n.d2
			n.d2 = n.d1
			n.d1 = &data
		} else if data < *n.d2 {
			n.d3 = n.d2
			n.d2 = &data
		} else {
			n.d3 = &data
		}
	}
}

func (n *Node) Split() {
	if n.nodeType < 4 {
		return
	}

	if n.parent == nil {
		leftChild := NewNode(*n.d1, n)
		rightChild := NewNode(*n.d3, n)
		leftChild.c1, leftChild.c2 = n.c1, n.c2
		rightChild.c1, rightChild.c2 = n.c3, n.c4
		n.nodeType = 2
		n.d1, n.d2, n.d3 = n.d2, nil, nil
		n.c1, n.c2, n.c3, n.c4 = leftChild, rightChild, nil, nil
	} else if n.parent.nodeType == 2 {
		if n == n.parent.c1 {
			midChild := NewNode(*n.d3, n.parent)
			midChild.c1, midChild.c2 = n.c3, n.c4
			n.parent.Push(*n.d2)
			n.parent.c1, n.parent.c2, n.parent.c3 = n.parent.c1, midChild, n.parent.c2
			n.nodeType = 2
			n.c1, n.c2, n.c3, n.c4 = n.c1, n.c2, nil, nil
			n.d1, n.d2, n.d3 = n.d1, nil, nil
		} else if n == n.parent.c2 {
			midChild := NewNode(*n.d1, n.parent)
			midChild.c1, midChild.c2 = n.c1, n.c2
			n.parent.Push(*n.d2)
			n.parent.c1, n.parent.c2, n.parent.c3 = n.parent.c1, midChild, n.parent.c2
			n.nodeType = 2
			n.c1, n.c2, n.c3, n.c4 = n.c3, n.c4, nil, nil
			n.d1, n.d2, n.d3 = n.d3, nil, nil
		}
	} else if n.parent.nodeType == 3 {
		if n == n.parent.c1 {
			newNode := NewNode(*n.d3, n.parent)
			newNode.c1, newNode.c2 = n.c3, n.c4
			n.parent.Push(*n.d2)
			n.parent.c1, n.parent.c2, n.parent.c3, n.parent.c4 = n.parent.c1, newNode, n.parent.c2, n.parent.c3
			n.nodeType = 2
			n.c1, n.c2, n.c3, n.c4 = n.c1, n.c2, nil, nil
			n.d1, n.d2, n.d3 = n.d1, nil, nil
		} else if n == n.parent.c2 {
			newNode := NewNode(*n.d3, n.parent)
			newNode.c1, newNode.c2 = n.c3, n.c4
			n.parent.Push(*n.d2)
			n.parent.c1, n.parent.c2, n.parent.c3, n.parent.c4 = n.parent.c1, n.parent.c2, newNode, n.parent.c3
			n.nodeType = 2
			n.c1, n.c2, n.c3, n.c4 = n.c1, n.c2, nil, nil
			n.d1, n.d2, n.d3 = n.d1, nil, nil
		} else if n == n.parent.c3 {
			newNode := NewNode(*n.d1, n.parent)
			newNode.c1, newNode.c2 = n.c1, n.c2
			n.parent.Push(*n.d2)
			n.parent.c1, n.parent.c2, n.parent.c3, n.parent.c4 = n.parent.c1, n.parent.c2, newNode, n.parent.c3
			n.nodeType = 2
			n.c1, n.c2, n.c3, n.c4 = n.c3, n.c4, nil, nil
			n.d1, n.d2, n.d3 = n.d3, nil, nil
		}
		n.parent.Split()
	}
}

func (n *Node) Insert(data int) {
	if n.c1 == nil {
		n.Push(data)
		n.Split()
	} else if n.nodeType == 2 {
		if data < *n.d1 {
			n.c1.Insert(data)
		} else {
			n.c2.Insert(data)
		}
	} else if n.nodeType == 3 {
		if data < *n.d1 {
			n.c1.Insert(data)
		} else if data > *n.d3 {
			n.c3.Insert(data)
		} else {
			n.c2.Insert(data)
		}
	}
}

func (n *Node) Find(data int) bool {
	if n.c1 == nil {
		return (n.d1 != nil && data == *n.d1) || (n.d2 != nil && data == *n.d2) || (n.d3 != nil && data == *n.d3)
	} else if n.nodeType == 2 {
		if data < *n.d1 {
			return n.c1.Find(data)
		} else {
			return n.c2.Find(data)
		}
	} else if n.nodeType == 3 {
		if data < *n.d1 {
			return n.c1.Find(data)
		} else if data > *n.d3 {
			return n.c3.Find(data)
		} else {
			return n.c2.Find(data)
		}
	}
	return false
}

type TwoThreeTree struct {
	isEmpty bool
	root    *Node
}

func NewTwoThreeTree() *TwoThreeTree {
	return &TwoThreeTree{isEmpty: true}
}

func (t *TwoThreeTree) Insert(data int) {
	if t.isEmpty {
		t.isEmpty = false
		t.root = NewNode(data, nil)
	} else {
		t.root.Insert(data)
	}
}

func (t *TwoThreeTree) Find(data int) bool {
	if t.isEmpty {
		return false
	} else {
		return t.root.Find(data)
	}
}

func main() {
	tree := NewTwoThreeTree()
	tree.Insert(5)
	tree.Insert(10)
	tree.Insert(15)
	tree.Insert(20)
	tree.Insert(25)

	fmt.Println(tree.Find(10)) // true
	fmt.Println(tree.Find(30)) // false
}
