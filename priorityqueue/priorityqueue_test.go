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

func TestPriorityQueue_Enqueue(t *testing.T) {
	t.Run("test enqueue element to the priority queue", func(t *testing.T) {
		pq := New()

		val := 2
		pq.Enqueue(val)
		err, got := pq.list.Get(0)

		if err != nil {
			t.Errorf("Enqueue() gotError = %v", err)
		}
		if got != val {
			t.Errorf("Enqueue() got = %v, want %v", got, val)
		}
	})
}
