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

func TestDynamicArray_Append(t *testing.T) {
	type fields struct {
		capacity int
		elements []interface{}
	}
	type args struct {
		element interface{}
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		indexExpected int
	}{
		{
			name: "test append element with capacity room available",
			fields: fields{
				capacity: 2,
				elements: []interface{}{1},
			},
			args:          args{element: 2},
			indexExpected: 1,
		},
		{
			name: "test append element with no capacity room available",
			fields: fields{
				capacity: 2,
				elements: []interface{}{1, 2},
			},
			args:          args{element: 3},
			indexExpected: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, dynamicArray := New(tt.fields.capacity, tt.fields.elements...)
			if err != nil {
				t.Errorf("can't make new dynamic array, error: %v", err)
			}
			if got := dynamicArray.Append(tt.args.element); got != tt.indexExpected {
				t.Errorf("Append() = %v, indexExpected %v", got, tt.indexExpected)
			}
		})
	}
}
