package main

import (
	"fmt"

	"gods/linkedlist"
)

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
	l.Display()
	l.Delete(56)	
	Footer()	
}
