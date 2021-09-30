package queue

import "github.com/KyriakosMilad/datastructures/doublylinkedlist"

type Queue struct {
	list *doublylinkedlist.DoublyLinkedList
}

func New(val interface{}) *Queue {
	list := doublylinkedlist.DoublyLinkedList{}
	list.AddLast(val)
	return &Queue{list: &list}
}

func (q *Queue) Enqueue(val interface{}) {
	q.list.AddLast(val)
}

func (q *Queue) Dequeue() error {
	return q.list.RemoveFirst()
}

func (q *Queue) Peek() (error, interface{}) {
	return q.list.Tail()
}
