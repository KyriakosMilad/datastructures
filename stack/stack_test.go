package stack

import (
	"reflect"
	"testing"
)

func TestStack_New(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test create new stack",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{}
			if got := s.New(); reflect.TypeOf(got) != reflect.TypeOf(&Stack{}) {
				t.Errorf("New() = %v, want %v", got, "*Stack")
			}
		})
	}
}
