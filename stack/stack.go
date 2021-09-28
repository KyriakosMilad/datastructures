package stack

import "github.com/KyriakosMilad/datastructures/doublylinkedlist"

type Stack struct {
	list *doublylinkedlist.DoublyLinkedList
}

func New() *Stack {
	return &Stack{&doublylinkedlist.DoublyLinkedList{}}
}
