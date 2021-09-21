package doublylinkedlist

import "errors"

type Node struct {
	next  *Node
	prev  *Node
	value interface{}
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func (dll *DoublyLinkedList) Size() int {
	return dll.size
}

func (dll *DoublyLinkedList) IsEmpty() bool {
	return dll.size == 0
}

func (dll *DoublyLinkedList) AddFirst(val interface{}) {
	newNode := &Node{value: val}

	switch dll.size {
	case 0:
		dll.head = newNode
		dll.tail = newNode
	case 1:
		dll.head = newNode
		dll.head.next = dll.tail
		dll.tail.prev = dll.head
	default:
		newNode.next = dll.head
		dll.head.prev = newNode
		dll.head = newNode
	}

	dll.size += 1
}

func (dll *DoublyLinkedList) AddLast(val interface{}) {
	newNode := &Node{value: val}
	if dll.IsEmpty() {
		dll.tail = newNode
		dll.head = newNode
	} else {
		newNode.prev = dll.tail
		dll.tail.next = newNode
		dll.tail = newNode
	}

	dll.size += 1
}

func (dll *DoublyLinkedList) Head() (error, interface{}) {
	if dll.IsEmpty() {
		return errors.New("can't get the head of empty doubly linked list"), nil
	}

	return nil, dll.head.value
}

func (dll *DoublyLinkedList) Tail() (error, interface{}) {
	if dll.IsEmpty() {
		return errors.New("can't get the head of empty doubly linked list"), nil
	}

	return nil, dll.tail.value
}
