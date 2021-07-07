package main

import (
	"fmt"

	"gods/bst"
	"gods/graph"
	"gods/hashtable"
	"gods/heap"
	"gods/linkedlist"
	"gods/queue"
	"gods/stack"
	"gods/trie"
)

func TestLinkedList(l linkedlist.LinkedList) {
	l.Insert(12)
	l.Insert(42)
	l.Insert(56)
	l.Insert(21)
	l.Insert(1)
	l.Insert(29)
	l.Display()
	l.Delete(21)
	l.Display()
	l.Delete(21)
	l.Display()
	l.Delete(12)
	l.Display()
	l.Delete(29)
	l.Display()
	l.Delete(1)
	l.Display()
	l.Delete(42)
	l.Display()
	l.Delete(56)
	l.Delete(56)
	Footer()
}

func TestStack(s stack.Stack) {
	s.Push(23)
	s.Display()
	s.Push(15)
	s.Display()
	s.Push(6)
	s.Display()
	s.Pop()
	s.Display()
	s.Push(34)
	s.Display()
	s.Push(11)
	s.Display()
	s.Push(11)
	s.Pop()
	s.Display()
	s.Pop()
	s.Display()
	s.Pop()
	s.Display()
	s.Pop()
	s.Display()
	s.Pop()
	s.Display()
	Footer()
}

func TestQueue(q queue.Queue) {
	q.Enqueue(6)
	q.Display()
	q.Dequeue()
	q.Display()
	q.Enqueue(22)
	q.Display()
	q.Enqueue(48)
	q.Display()
	q.Enqueue(93)
	q.Display()
	q.Enqueue(62)
	q.Display()
	q.Dequeue()
	q.Display()
	q.Dequeue()
	q.Display()
	q.Dequeue()
	q.Display()
	q.Dequeue()
	q.Display()
	Footer()
}

func TestBST(t bst.Node) {
	t.Insert(40)
	t.Insert(60)
	t.Insert(20)
	t.Insert(30)
	t.Insert(70)
	t.Insert(90)
	t.Insert(80)
	t.Display()
	Footer()
}

func TestHeap() {
	in := []int{34, 65, 7, 12, 80, 23}
	min := heap.NewMinHeap(len(in))
	for i := 0; i < len(in); i++ {
		min.Insert(in[i])
	}
	for i := 0; i < len(in); i++ {
		fmt.Println(min.Remove())
	}
	Footer()
}

func TestHashTable() {
	table := make(map[int]*hashtable.Node, 10)
	hash := hashtable.HashTable{Table: table, Size: 10}
	fmt.Println("Number of spaces:", hash.Size)
	for i := 0; i < 60; i++ {
		hash.Insert(i)
	}
	hash.Display()
	Footer()
}

func TestTrie() {
	t := trie.NewTrie()
	in := []string{
		"aragorn",
		"aragon",
		"argon",
		"eragon",
		"oregon",
		"oregano",
		"oreo",
	}
	for _, i := range in {
		t.Insert(i)
		fmt.Println("Entered", i)
	}
	fmt.Printf("Searching for 'argon': %t \n", t.Search("argon"))
	fmt.Printf("Searching for 'wizard': %t \n", t.Search("wizard"))
	Footer()
}

func TestGraph() {
	g := graph.Graph{}
	for i := 0; i < 5; i++ {
		g.AddVertex(i)
	}
	g.AddEdge(1, 2)
	g.AddEdge(1, 2)
	g.AddEdge(6, 2)
	g.AddEdge(3, 2)
	g.Display()
}

func Begin() {
	fmt.Println("==============================")
}

func Header(n string) {
	fmt.Printf("%s Output", n)
	fmt.Println("\n==============================")
}

func Footer() {
	fmt.Println("==============================")
}

func main() {
	// Begin
	Begin()

	// Test LinkedList
	Header("LinkedList")
	l := linkedlist.LinkedList{}
	TestLinkedList(l)

	// Test Stack
	Header("Stack")
	s := stack.Stack{}
	TestStack(s)

	// Test Queue
	Header("Queue")
	q := queue.Queue{}
	TestQueue(q)

	// Test BST
	Header("Binary Search Tree")
	t := bst.Node{Key: 50}
	TestBST(t)

	// Test Heap
	Header("MinHeap")
	TestHeap()

	// Test HashTable
	Header("HashTable")
	TestHashTable()

	// Test Trie
	Header("Trie")
	TestTrie()

	// Test Graph
	Header("Graph")
	TestGraph()
}
