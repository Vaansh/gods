package main

import (
	"fmt"

	"gods/bst"
	"gods/heap"
	"gods/linkedlist"
	"gods/queue"
	"gods/stack"
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

func TestHeap(h heap.MinHeap, in []int) {
	for i := 0; i < len(in); i++ {
		h.Insert(in[i])
	}
	for i := 0; i < len(in); i++ {
		fmt.Println(h.Remove())
	}
	Footer()
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
	in := []int{34, 65, 7, 12, 80, 23}
	min := heap.NewMinHeap(len(in))
	TestHeap(*min, in)
}
