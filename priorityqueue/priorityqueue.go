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
