package main

import "sort"

// Node - узел 2-3 дерева
type Node struct {
	keys     []int
	children []*Node
	parent   *Node
}

// Tree - дерево
type Tree struct {
	root *Node
}

// newNode создает новый узел
func newNode(parent *Node) *Node {
	return &Node{
		keys:     make([]int, 0, 3),
		children: make([]*Node, 0, 4),
		parent:   parent,
	}
}

// newTree создает новое дерево
func newTree() *Tree {
	return &Tree{
		root: newNode(nil),
	}
}

// insert вставляет ключ в дерево
func (tree *Tree) insert(key int) {
	tree.root = insert(tree.root, key)
	if len(tree.root.keys) == 3 {
		tree.root = splitRoot(tree.root)
	}
}

func insert(node *Node, key int) *Node {
	if len(node.children) == 0 {
		node.keys = insertIntoLeaf(node.keys, key)
		return node
	}

	childIndex := findChildIndex(node.keys, key)
	child := insert(node.children[childIndex], key)
	if len(child.keys) == 3 {
		node = splitChild(node, childIndex)
	}

	return node
}

func insertIntoLeaf(keys []int, key int) []int {
	keys = append(keys, key)
	sort.Ints(keys)
	return keys
}

func splitRoot(node *Node) *Node {
	midKey := node.keys[1]
	left := &Node{keys: []int{node.keys[0]}, parent: nil}
	right := &Node{keys: []int{node.keys[2]}, parent: nil}

	if len(node.children) > 0 {
		left.children = append(left.children, node.children[0], node.children[1])
		right.children = append(right.children, node.children[2], node.children[3])
		for _, child := range left.children {
			child.parent = left
		}
		for _, child := range right.children {
			child.parent = right
		}
	}

	newRoot := &Node{keys: []int{midKey}, children: []*Node{left, right}}
	left.parent = newRoot
	right.parent = newRoot
	return newRoot
}

func splitChild(parent *Node, childIndex int) *Node {
	child := parent.children[childIndex]
	midKey := child.keys[1]
	left := &Node{keys: []int{child.keys[0]}, parent: parent}
	right := &Node{keys: []int{child.keys[2]}, parent: parent}

	if len(child.children) > 0 {
		left.children = append(left.children, child.children[0], child.children[1])
		right.children = append(right.children, child.children[2], child.children[3])
		for _, c := range left.children {
			c.parent = left
		}
		for _, c := range right.children {
			c.parent = right
		}
	}

	parent.keys = append(parent.keys, 0)
	copy(parent.keys[childIndex+1:], parent.keys[childIndex:])
	parent.keys[childIndex] = midKey

	parent.children = append(parent.children, nil)
	copy(parent.children[childIndex+2:], parent.children[childIndex+1:])
	parent.children[childIndex] = left
	parent.children[childIndex+1] = right

	return parent
}

func findChildIndex(keys []int, key int) int {
	for i := range keys {
		if key < keys[i] {
			return i
		}
	}

	return len(keys)
}

func main() {
	tree := newTree()
	keys := []int{10, 20, 5, 6, 15, 30, 25, 35, 4}
	for _, key := range keys {
		tree.insert(key)
	}
}
