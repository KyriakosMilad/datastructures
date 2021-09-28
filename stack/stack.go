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
