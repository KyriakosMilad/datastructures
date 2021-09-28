package doublylinkedlist

import "errors"

type node struct {
	next  *node
	prev  *node
	value interface{}
}

type DoublyLinkedList struct {
	head *node
	tail *node
	size int
}

func (dll *DoublyLinkedList) Size() int {
	return dll.size
}

func (dll *DoublyLinkedList) IsEmpty() bool {
	return dll.size == 0
}

func (dll *DoublyLinkedList) AddFirst(val interface{}) {
	newNode := &node{value: val}

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
	newNode := &node{value: val}

	switch dll.size {
	case 0:
		dll.head = newNode
		dll.tail = newNode
	case 1:
		dll.tail = newNode
		dll.tail.prev = dll.head
		dll.head.next = dll.tail
	default:
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

func (dll *DoublyLinkedList) IndexOfNode(node *node) int {
	counter := 0
	traverser := dll.head
	for traverser != nil {
		if traverser == node {
			return counter
		}
		counter += 1
		traverser = traverser.next
	}

	return -1
}

func (dll *DoublyLinkedList) ContainsNode(node *node) bool {
	return dll.IndexOfNode(node) != -1
}

func (dll *DoublyLinkedList) IndexOfValue(val interface{}) int {
	counter := 0
	traverser := dll.head
	for traverser != nil {
		if traverser.value == val {
			return counter
		}
		counter += 1
		traverser = traverser.next
	}

	return -1
}

func (dll *DoublyLinkedList) ContainsValue(val interface{}) bool {
	return dll.IndexOfValue(val) != -1
}

func (dll *DoublyLinkedList) RemoveFirst() error {
	switch dll.size {
	case 0:
		return errors.New("can't remove from empty doubly linked list")
	case 1:
		dll.head.value = nil
		dll.tail = nil
		dll.head = nil
	case 2:
		dll.head.value = nil
		dll.head = dll.tail
		dll.head.prev = nil
		dll.head.next = nil
	default:
		dll.head.value = nil
		dll.head = dll.head.next
		dll.head.prev = nil
	}

	dll.size -= 1

	return nil
}

func (dll *DoublyLinkedList) RemoveLast() error {
	switch dll.size {
	case 0:
		return errors.New("can't remove from empty doubly linked list")
	case 1:
		dll.tail.value = nil
		dll.tail = nil
		dll.head = nil
	case 2:
		dll.tail.value = nil
		dll.tail = dll.head
		dll.tail.prev = nil
		dll.tail.next = nil
	default:
		dll.tail.value = nil
		dll.tail = dll.tail.prev
		dll.tail.next = nil
	}

	dll.size -= 1

	return nil
}

func (dll *DoublyLinkedList) Remove(node *node) error {
	if !dll.ContainsNode(node) {
		return errors.New("this node does not belong to this DoublyLinkedList")
	}

	if node.prev == nil {
		return dll.RemoveFirst()
	}
	if node.next == nil {
		return dll.RemoveLast()
	}

	node.prev.next = node.next
	node.next.prev = node.prev

	node.next = nil
	node.prev = nil
	node.value = nil

	dll.size -= 1

	return nil
}

func (dll *DoublyLinkedList) RemoveAt(index int) error {
	if dll.size == 0 {
		return errors.New("can't remove from empty doubly linked list")
	}
	if index < 0 || index > dll.size {
		return errors.New("index out of range")
	}

	counter := 0
	traverser := dll.head
	for traverser != nil {
		if counter == index {
			return dll.Remove(traverser)
		}
		counter += 1
		traverser = traverser.next
	}

	return nil
}

func (dll *DoublyLinkedList) Clear() {
	traverser := dll.head
	for traverser != nil {
		next := traverser.next
		traverser.value = nil
		traverser.prev = nil
		traverser.next = nil
		dll.size -= 1
		traverser = next
	}
	dll.head = nil
	dll.tail = nil
}
