package queue

import (
	"reflect"
	"testing"
)

func TestQueue_New(t *testing.T) {
	t.Run("test create new queue", func(t *testing.T) {
		if got := New(1); reflect.TypeOf(got) != reflect.TypeOf(&Queue{}) {
			t.Errorf("New() = %v, want %v", got, "*Queue")
		}
	})
}

func TestQueue_Enqueue(t *testing.T) {
	t.Run("test enqueue element to the queue", func(t *testing.T) {
		q := New(1)

		val := 2
		q.Enqueue(val)
		err, got := q.list.Tail()

		if err != nil {
			t.Errorf("Enqueue() gotError = %v", err)
		}
		if got != val {
			t.Errorf("Enqueue() got = %v, want %v", got, val)
		}
	})
}
