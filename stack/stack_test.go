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

func TestStack_Pop(t *testing.T) {
	t.Run("test removing element from the stack", func(t *testing.T) {
		s := New()

		val := 2
		s.Push(val)

		err := s.Pop()
		if err != nil {
			t.Errorf("Pop() gotError = %v", err)
		}

		err, _ = s.list.Tail()
		if err == nil {
			t.Errorf("Pop() not working, can not remove element from the Stack")
		}
	})
}

func TestStack_Peek(t *testing.T) {
	t.Run("test get last element from the stack", func(t *testing.T) {
		s := New()

		val := 2
		s.Push(val)

		err, lastElement := s.Peek()
		if err != nil {
			t.Errorf("Peek() gotError = %v", err)
		}
		if lastElement != val {
			t.Errorf("Peek() not working, got = %v, want = %v", lastElement, val)
		}
	})
}

func TestStack_Size(t *testing.T) {
	t.Run("test get the size of the stack", func(t *testing.T) {
		s := New()

		val := 2
		s.Push(val)

		size := s.Size()
		if size != 1 {
			t.Errorf("Size() not working, got = %v, want = %v", size, 1)
		}
	})
}
