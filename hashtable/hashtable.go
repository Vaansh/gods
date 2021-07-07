package hashtable

import "fmt"

type Node struct {
	Next  *Node
	Value int
}

type HashTable struct {
	Table map[int]*Node
	Size  int
}

func (h *HashTable) Insert(i int) int {
	temp := h.HashFunction(i)
	node := Node{Next: h.Table[temp], Value: i}
	h.Table[temp] = &node
	return temp
}

func (h *HashTable) HashFunction(i int) int {
	return i % h.Size
}

func (h *HashTable) Display() {
	for i := range h.Table {
		if h.Table[i] != nil {
			temp := h.Table[i]
			for temp != nil {
				fmt.Printf("%v ->", temp.Value)
				temp = temp.Next
			}
			fmt.Println()
		}
	}
}
