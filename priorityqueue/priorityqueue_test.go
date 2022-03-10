package priorityqueue

import (
	"reflect"
	"testing"
)

func TestPriorityQueue_New(t *testing.T) {
	t.Run("test create new priority queue", func(t *testing.T) {
		if got := New(); reflect.TypeOf(got) != reflect.TypeOf(&PriorityQueue{}) {
			t.Errorf("New() = %v, want %v", got, "*PriorityQueue")
		}
	})
}
