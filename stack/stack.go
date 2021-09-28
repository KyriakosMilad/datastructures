package stack

import "github.com/KyriakosMilad/datastructures/doublylinkedlist"

type Stack struct {
	list doublylinkedlist.DoublyLinkedList
}

func (s *Stack) New() *Stack {
	return &Stack{doublylinkedlist.DoublyLinkedList{}}
}
