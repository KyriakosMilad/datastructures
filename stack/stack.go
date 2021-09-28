package stack

import "github.com/KyriakosMilad/datastructures/doublylinkedlist"

type Stack struct {
	list *doublylinkedlist.DoublyLinkedList
}

func New() *Stack {
	return &Stack{&doublylinkedlist.DoublyLinkedList{}}
}

func (s *Stack) Push(value interface{}) {
	s.list.AddLast(value)
}

func (s *Stack) Pop() error {
	return s.list.RemoveLast()
}

func (s *Stack) Peek() (error, interface{}) {
	return s.list.Tail()
}

func (s *Stack) Size() int {
	return s.list.Size()
}

func (s *Stack) IsEmpty() bool {
	return s.list.IsEmpty()
}
