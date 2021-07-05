package linkedlist

import "fmt"

type Node struct {
	next *Node
	data int
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) Insert(d int) {
	n := &Node{
		next: nil,
		data: d,
	}
	n.next = l.head
	l.head = n
	l.length++
}

func (l *LinkedList) Delete(d int) {
	if l.length == 0 {
		fmt.Println("List is empty")
		return
	}
	if l.head.data == d {
		l.head = l.head.next
		l.length--
		return
	}
	temp := l.head
	for temp.next.data != d {
		if temp.next.next == nil {
			fmt.Println("Element does not exist")
			return
		}
		temp = temp.next
	}
	temp.next = temp.next.next
	l.length--
}

func (l *LinkedList) Display() {
	temp := l.head
	for temp != nil {
		fmt.Printf("%+v ->", temp.data)
		temp = temp.next
	}
	fmt.Println()
}
