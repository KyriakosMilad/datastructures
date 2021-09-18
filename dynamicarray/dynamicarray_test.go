package dyanmicarray

import (
	"reflect"
	"testing"
)

func TestDynamicArray_New(t *testing.T) {
	type args struct {
		capacity int
		elements []interface{}
	}
	tests := []struct {
		name             string
		args             args
		errorExpected    error
		wantDynamicArray bool
	}{
		{
			name: "test create new dynamicArray",
			args: args{
				capacity: 10,
				elements: []interface{}{'a', 'b', 'c'},
			},
			errorExpected:    nil,
			wantDynamicArray: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, da := New(tt.args.capacity, tt.args.elements...)
			if !reflect.DeepEqual(err, tt.errorExpected) {
				t.Errorf("New() error = %v, want %v", err, tt.errorExpected)
			}
			if (da != nil) != tt.wantDynamicArray {
				t.Errorf("New() dynamicArrayReturned = %v, want %v", da != nil, tt.wantDynamicArray)
			}
		})
	}
}
