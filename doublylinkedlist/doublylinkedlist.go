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
