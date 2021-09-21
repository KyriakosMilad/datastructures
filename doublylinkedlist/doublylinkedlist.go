package doublylinkedlist

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
	if dll.IsEmpty() {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.next = dll.head
		dll.head.prev = newNode
		dll.head = newNode
	}

	dll.size += 1
}
