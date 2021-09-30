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
