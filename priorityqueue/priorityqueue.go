package priorityqueue

import (
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
