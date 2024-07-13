// https://neerc.ifmo.ru/wiki/index.php?title=2-3_%D0%B4%D0%B5%D1%80%D0%B5%D0%B2%D0%BE
package main

type Node struct {
	keys   []int
	sons   []*Node
	isLeaf bool
}

func search(root *Node, x int) int {
	t := root
	for !t.isLeaf {
		if len(t.keys) == 2 {
			if t.keys[0] < x {
				t = t.sons[1]
			} else {
				t = t.sons[0]
			}
		} else {
			if t.keys[1] < x {
				t = t.sons[2]
			} else if t.keys[0] < x {
				t = t.sons[1]
			} else {
				t = t.sons[0]
			}
		}
	}
	return t.keys[0]
}
