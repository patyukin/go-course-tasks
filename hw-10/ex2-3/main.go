// https://github.com/strictlysimpledesign/23tree/blob/master/23tree.py
package main

import (
	"fmt"
	"sort"
)

type Node struct {
	data   []int
	childs map[string]*Node
	parent *Node
}

var path []string
var root *Node

func NewNode(data int, parent *Node) *Node {
	return &Node{
		data:   []int{data},
		childs: make(map[string]*Node),
		parent: parent,
	}
}

func (n *Node) insert(value int) {
	path = []string{}
	insertNode := n.search(value)
	insertNode.add(value)
}

func (n *Node) split() {
	if n.parent == nil && len(n.childs) > 0 {
		branch := path[len(path)-1]
		path = path[:len(path)-1]
		newNodeLeft := NewNode(n.data[0], n)
		newNodeRight := NewNode(n.data[1], n)
		n.data = n.data[1:]
		switch branch {
		case "left":
			newNodeLeft.childs["left"] = n.childs["left"]
			newNodeLeft.childs["right"] = n.childs["overflow"]
			newNodeRight.childs["left"] = n.childs["mid"]
			newNodeRight.childs["right"] = n.childs["right"]
		case "mid":
			newNodeLeft.childs["left"] = n.childs["left"]
			newNodeLeft.childs["right"] = n.childs["mid"]
			newNodeRight.childs["left"] = n.childs["overflow"]
			newNodeRight.childs["right"] = n.childs["right"]
		case "right":
			newNodeLeft.childs["left"] = n.childs["left"]
			newNodeLeft.childs["right"] = n.childs["mid"]
			newNodeRight.childs["left"] = n.childs["right"]
			newNodeRight.childs["right"] = n.childs["overflow"]
		}
		newNodeLeft.childs["left"].parent = newNodeLeft
		newNodeLeft.childs["right"].parent = newNodeLeft
		newNodeRight.childs["left"].parent = newNodeRight
		newNodeRight.childs["right"].parent = newNodeRight
		n.childs["left"] = newNodeLeft
		n.childs["right"] = newNodeRight
		delete(n.childs, "mid")
	} else if n.parent != nil && len(n.childs) > 0 {
		branch := path[len(path)-1]
		path = path[:len(path)-1]
		newNode := NewNode(n.data[2], n.parent)
		n.data = n.data[:2]
		n.parent.childs["overflow"] = newNode
		switch branch {
		case "left":
			newNode.childs["left"] = n.childs["mid"]
			newNode.childs["right"] = n.childs["right"]
			n.childs["right"] = n.childs["overflow"]
		case "mid":
			newNode.childs["left"] = n.childs["overflow"]
			newNode.childs["right"] = n.childs["right"]
			n.childs["right"] = n.childs["mid"]
		case "right":
			newNode.childs["left"] = n.childs["right"]
			newNode.childs["right"] = n.childs["overflow"]
			n.childs["right"] = n.childs["mid"]
		}
		newNode.childs["left"].parent = newNode
		newNode.childs["right"].parent = newNode
		delete(n.childs, "mid")
	} else if n.parent == nil && len(n.childs) == 0 {
		n.childs["left"] = NewNode(n.data[0], n)
		n.childs["right"] = NewNode(n.data[1], n)
		n.data = n.data[1:]
	} else if n.parent != nil && len(n.childs) == 0 {
		n.parent.childs["overflow"] = NewNode(n.data[2], n.parent)
		n.data = n.data[:2]
	}
}

func (n *Node) add(value int) {
	if !contains(n.data, value) {
		n.data = append(n.data, value)
		sort.Ints(n.data)
		if len(n.data) == 3 {
			n.split()
			if n.parent != nil {
				n.parent.add(n.data[2])
				n.data = n.data[:2]
			}
		} else {
			if _, exists := n.childs["overflow"]; exists {
				branch := path[len(path)-1]
				path = path[:len(path)-1]
				switch branch {
				case "left":
					n.childs["mid"] = n.childs["overflow"]
				case "right":
					n.childs["mid"] = n.childs["right"]
					n.childs["right"] = n.childs["overflow"]
				}
				delete(n.childs, "overflow")
			}
		}
	}
}

func (n *Node) search(value int) *Node {
	if len(n.childs) > 0 {
		boundLeft := minValue(n.data)
		boundRight := maxValue(n.data)
		if value < boundLeft {
			path = append(path, "left")
			return n.childs["left"].search(value)
		} else if value > boundRight {
			path = append(path, "right")
			return n.childs["right"].search(value)
		} else {
			path = append(path, "mid")
			return n.childs["mid"].search(value)
		}
	}
	return n
}

func (n *Node) element(value int) bool {
	if contains(n.data, value) {
		return true
	} else if len(n.childs) > 0 {
		boundLeft := minValue(n.data)
		boundRight := maxValue(n.data)
		if value < boundLeft {
			return n.childs["left"].element(value)
		} else if value > boundRight {
			return n.childs["right"].element(value)
		} else {
			return n.childs["mid"].element(value)
		}
	}
	return false
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func minValue(slice []int) int {
	m := slice[0]
	for _, v := range slice {
		if v < m {
			m = v
		}
	}

	return m
}

func maxValue(slice []int) int {
	m := slice[0]
	for _, v := range slice {
		if v > m {
			m = v
		}
	}

	return m
}

func main() {
	root = NewNode(10, nil)
	root.insert(20)
	root.insert(5)
	root.insert(15)

	fmt.Println("Root data:", root.data)
	fmt.Println("Root left child data:", root.childs["left"].data)
	fmt.Println("Root right child data:", root.childs["right"].data)
}
