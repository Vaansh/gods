package bst

import "fmt"

type Tree struct {
	root *Node
}

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(d int) {
	if n.Key < d {
		if n.Right == nil {
			n.Right = &Node{Key: d}
		} else {
			n.Right.Insert(d)
		}
	} else {
		if n.Left == nil {
			n.Left = &Node{Key: d}
		} else {
			n.Left.Insert(d)
		}
	}
}

func (n *Node) Search(d int) bool {
	if n == nil {
		return false
	}
	if n.Key < d {
		return n.Right.Search(d)
	} else if n.Key > d {
		return n.Left.Search(d)
	}
	return true
}

func (n *Node) Display() {
	if n.Left != nil {
		n.Left.Display()
	}
	fmt.Println(n.Key)
	if n.Right != nil {
		n.Right.Display()
	}
}
