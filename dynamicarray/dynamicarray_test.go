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

func TestDynamicArray_InsertAt(t *testing.T) {
	type fields struct {
		capacity int
		elements []interface{}
	}
	type args struct {
		index   int
		element interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test insert element with capacity room available",
			fields: fields{
				capacity: 3,
				elements: []interface{}{1, 3},
			},
			args:    args{index: 1, element: 2},
			wantErr: false,
		},
		{
			name: "test insert element with no capacity room available",
			fields: fields{
				capacity: 3,
				elements: []interface{}{1, 2, 4},
			},
			args:    args{index: 2, element: 3},
			wantErr: false,
		},
		{
			name: "test insert element at the end",
			fields: fields{
				capacity: 2,
				elements: []interface{}{1},
			},
			args:    args{index: 1, element: 2},
			wantErr: false,
		},
		{
			name: "test insert element at the beginning",
			fields: fields{
				capacity: 2,
				elements: []interface{}{2},
			},
			args:    args{index: 0, element: 1},
			wantErr: false,
		},
		{
			name: "test insert element with index lower than 0",
			fields: fields{
				capacity: 2,
				elements: []interface{}{2},
			},
			args:    args{index: -1, element: 1},
			wantErr: true,
		},
		{
			name: "test insert element with index bigger than the DynamicArray size",
			fields: fields{
				capacity: 2,
				elements: []interface{}{2},
			},
			args:    args{index: 2, element: 1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, dynamicArray := New(tt.fields.capacity, tt.fields.elements...)
			if err != nil {
				t.Errorf("can't make new dynamic array, error: %v", err)
			}
			if err := dynamicArray.InsertAt(tt.args.index, tt.args.element); (err != nil) != tt.wantErr {
				t.Errorf("InsertAt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
