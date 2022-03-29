package priorityqueue

import (
	"errors"
	"github.com/KyriakosMilad/datastructures/dynamicarray"
)

type PriorityQueue struct {
	list *dyanmicarray.DynamicArray
}

func New() *PriorityQueue {
	return &PriorityQueue{
		list: &dyanmicarray.DynamicArray{},
	}
}

func (pq *PriorityQueue) Enqueue(val int) {
	traverser := val
	for i, v := range pq.list.Elements() {
		v, ok := v.(int)
		if !ok {
			panic("Unexpected error: v is not an integer")
		}
		if traverser < v {
			pq.list.Set(i, traverser)
			traverser = v
		}
	}
	pq.list.Append(traverser)
}

func (pq *PriorityQueue) Dequeue() error {
	return pq.list.RemoveLast()
}

func (pq *PriorityQueue) Size() int {
	return pq.list.Size()
}

func (pq *PriorityQueue) IsEmpty() bool {
	return pq.list.IsEmpty()
}

func (pq *PriorityQueue) Last() (error, interface{}) {
	if pq.IsEmpty() {
		return errors.New("can't get last element from empty priority queue"), 0
	}
	return nil, pq.list.Elements()[pq.Size()-1]
}
