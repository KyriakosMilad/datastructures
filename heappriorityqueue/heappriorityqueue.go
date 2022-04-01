package heappriorityqueue

import (
	"github.com/KyriakosMilad/datastructures/binaryheap"
)

type HeapPriorityQueue struct {
	list *binaryheap.MaxBinaryHeap
}

func New() *HeapPriorityQueue {
	return &HeapPriorityQueue{list: &binaryheap.MaxBinaryHeap{}}
}
