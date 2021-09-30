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

func TestQueue_Dequeue(t *testing.T) {
	t.Run("test dequeue element from the queue", func(t *testing.T) {
		q := New(1)

		err := q.Dequeue()
		if err != nil {
			t.Errorf("Dequeue() gotError = %v", err)
		}

		err, _ = q.list.Tail()
		if err == nil {
			t.Errorf("Dequeue() not working, can not dequeue element from the Queue")
		}
	})
}

func TestQueue_Peek(t *testing.T) {
	t.Run("test get last element from the queue", func(t *testing.T) {
		val := 4
		q := New(val)

		err, lastElement := q.Peek()
		if err != nil {
			t.Errorf("Peek() gotError = %v", err)
		}
		if lastElement != val {
			t.Errorf("Peek() not working, got = %v, want = %v", lastElement, val)
		}
	})
}
