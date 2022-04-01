package heappriorityqueue

import (
	"github.com/KyriakosMilad/datastructures/binaryheap"
	"reflect"
	"testing"
)

func TestHeapPriorityQueue_New(t *testing.T) {
	t.Run("test create new heap priority queue", func(t *testing.T) {
		if got := New(); !reflect.DeepEqual(got, &HeapPriorityQueue{list: &binaryheap.MaxBinaryHeap{}}) {
			t.Errorf("New() = %v, want %v", got, "*HeapPriorityQueue")
		}
	})
}
