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

func TestPriorityQueue_Dequeue(t *testing.T) {
	t.Run("test dequeue element from the priority queue", func(t *testing.T) {
		pq := New()
		pq.Enqueue(1)

		err := pq.Dequeue()
		if err != nil {
			t.Errorf("Dequeue() gotError = %v", err)
		}

		err, _ = pq.list.Get(0)
		if err == nil {
			t.Errorf("Dequeue() not working, can not dequeue element from the PriorityQueue")
		}
	})
}

func TestPriorityQueue_Size(t *testing.T) {
	t.Run("test get the size of the priority queue", func(t *testing.T) {
		pq := New()
		pq.Enqueue(0)

		size := pq.Size()
		if size != 1 {
			t.Errorf("Size() not working, got = %v, want = %v", size, 1)
		}
	})
}

func TestPriorityQueue_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "check if empty priority queue is empty",
			want: true,
		},
		{
			name: "check if priority queue is empty",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := New()

			if !tt.want {
				pq.Enqueue(0)
			}

			if got := pq.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
