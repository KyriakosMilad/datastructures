package stack

import (
	"reflect"
	"testing"
)

func TestStack_New(t *testing.T) {
	t.Run("test create new stack", func(t *testing.T) {
		if got := New(); reflect.TypeOf(got) != reflect.TypeOf(&Stack{}) {
			t.Errorf("New() = %v, want %v", got, "*Stack")
		}
	})
}

func TestStack_Push(t *testing.T) {
	t.Run("test adding element to the stack", func(t *testing.T) {
		s := New()

		val := 2
		s.Push(val)
		err, got := s.list.Tail()

		if err != nil {
			t.Errorf("Push() gotError = %v", err)
		}
		if got != val {
			t.Errorf("Push() got = %v, want %v", got, val)
		}
	})
}
