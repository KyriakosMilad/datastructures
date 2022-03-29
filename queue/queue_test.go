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

func TestQueue_Last(t *testing.T) {
	t.Run("test get last element from the queue", func(t *testing.T) {
		val := 4
		q := New(val)

		err, lastElement := q.Last()
		if err != nil {
			t.Errorf("Last() gotError = %v", err)
		}
		if lastElement != val {
			t.Errorf("Last() not working, got = %v, want = %v", lastElement, val)
		}
	})
}

func TestQueue_Size(t *testing.T) {
	t.Run("test get the size of the queue", func(t *testing.T) {
		q := New(0)

		size := q.Size()
		if size != 1 {
			t.Errorf("Size() not working, got = %v, want = %v", size, 1)
		}
	})
}

func TestQueue_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "check if empty queue is empty",
			want: true,
		},
		{
			name: "check if queue is empty",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := New(0)

			if tt.want {
				err := q.list.RemoveFirst()
				if err != nil {
					t.Errorf("IsEmpty() error: %v", err)
				}
			}

			if got := q.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
