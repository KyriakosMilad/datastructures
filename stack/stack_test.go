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
