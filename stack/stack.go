package stack

import "fmt"

type Stack struct {
	stack []int
}

func (s *Stack) Push(d int) {
	s.stack = append(s.stack, d)
}

func (s *Stack) Pop() int {
	index := s.Size() - 1
	popped := s.stack[index]
	s.stack = s.stack[:index]
	return popped
}

func (s *Stack) Top() int {
	index := s.Size() - 1
	top := s.stack[index]
	return top
}

func (s *Stack) Size() int {
	return len(s.stack)
}

func (s *Stack) Display() {
	if s.Size() == 0 {
		fmt.Println("Stack is empty")
		return
	}
	for _, e := range s.stack {
		fmt.Printf("%+v ~>", e)
	}
	fmt.Println()
}
