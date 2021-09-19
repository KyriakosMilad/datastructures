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

func TestDynamicArray_RemoveAt(t *testing.T) {
	type fields struct {
		capacity int
		elements []interface{}
	}
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test remove element at specific index",
			fields: fields{
				capacity: 2,
				elements: []interface{}{1, 2},
			},
			args:    args{index: 0},
			wantErr: false,
		},
		{
			name: "test remove element at index lower than 0",
			fields: fields{
				capacity: 2,
				elements: []interface{}{1, 2},
			},
			args:    args{index: -1},
			wantErr: true,
		},
		{
			name: "test remove element at index bigger than the DynamicArray size",
			fields: fields{
				capacity: 2,
				elements: []interface{}{1, 2},
			},
			args:    args{index: 3},
			wantErr: true,
		},
		{
			name: "test remove element from empty array",
			fields: fields{
				capacity: 2,
				elements: []interface{}{},
			},
			args:    args{index: 3},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, dynamicArray := New(tt.fields.capacity, tt.fields.elements...)
			if err != nil {
				t.Errorf("can't make new dynamic array, error: %v", err)
			}

			if err := dynamicArray.RemoveAt(tt.args.index); (err != nil) != tt.wantErr {
				t.Errorf("RemoveAt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDynamicArray_RemoveAllWhere(t *testing.T) {
	type fields struct {
		capacity int
		elements []interface{}
	}
	type args struct {
		element interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test remove multiple elements from DynamicArray",
			fields: fields{
				capacity: 3,
				elements: []interface{}{1, 2, 2},
			},
			args:    args{element: 2},
			wantErr: false,
		},
		{
			name: "test remove one element from DynamicArray using RemoveAllWhere() method",
			fields: fields{
				capacity: 2,
				elements: []interface{}{1, 2},
			},
			args:    args{element: 2},
			wantErr: false,
		},
		{
			name: "test remove element that does not exist using RemoveAllWhere() method",
			fields: fields{
				capacity: 2,
				elements: []interface{}{1, 2},
			},
			args:    args{element: 3},
			wantErr: false,
		},
		{
			name: "test remove element from empty array",
			fields: fields{
				capacity: 2,
				elements: []interface{}{},
			},
			args:    args{element: 3},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, dynamicArray := New(tt.fields.capacity, tt.fields.elements...)
			if err != nil {
				t.Errorf("can't make new dynamic array, error: %v", err)
			}

			if err := dynamicArray.RemoveAllWhere(tt.args.element); (err != nil) != tt.wantErr {
				t.Errorf("RemoveAllWhere() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDynamicArray_RemoveFirstWhere(t *testing.T) {
	type fields struct {
		capacity int
		elements []interface{}
	}
	type args struct {
		element interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test remove one element from DynamicArray",
			fields: fields{
				capacity: 3,
				elements: []interface{}{1, 2, 2},
			},
			args:    args{element: 2},
			wantErr: false,
		},
		{
			name: "test remove element that does not exist using RemoveFirstWhere() method",
			fields: fields{
				capacity: 2,
				elements: []interface{}{1, 2},
			},
			args:    args{element: 3},
			wantErr: false,
		},
		{
			name: "test remove element from empty array",
			fields: fields{
				capacity: 2,
				elements: []interface{}{},
			},
			args:    args{element: 3},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, dynamicArray := New(tt.fields.capacity, tt.fields.elements...)
			if err != nil {
				t.Errorf("can't make new dynamic array, error: %v", err)
			}
			if err := dynamicArray.RemoveFirstWhere(tt.args.element); (err != nil) != tt.wantErr {
				t.Errorf("RemoveFirstWhere() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDynamicArray_RemoveLast(t *testing.T) {
	type fields struct {
		capacity int
		elements []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "test remove last element from DynamicArray",
			fields: fields{
				capacity: 3,
				elements: []interface{}{1, 2, 2},
			},
			wantErr: false,
		},
		{
			name: "test remove element from empty array",
			fields: fields{
				capacity: 2,
				elements: []interface{}{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, dynamicArray := New(tt.fields.capacity, tt.fields.elements...)
			if err != nil {
				t.Errorf("can't make new dynamic array, error: %v", err)
			}
			if err := dynamicArray.RemoveLast(); (err != nil) != tt.wantErr {
				t.Errorf("RemoveLast() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDynamicArray_Elements(t *testing.T) {
	type fields struct {
		capacity int
		size     int
		elements []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   []interface{}
	}{
		{
			name: "test get DynamicArray elements",
			fields: fields{
				capacity: 2,
				elements: []interface{}{1, 2},
			},
			want: []interface{}{1, 2},
		},
		{
			name: "test get DynamicArray elements",
			fields: fields{
				capacity: 0,
				elements: []interface{}{},
			},
			want: []interface{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, dynamicArray := New(tt.fields.capacity, tt.fields.elements...)
			if err != nil {
				t.Errorf("can't make new dynamic array, error: %v", err)
			}
			if got := dynamicArray.Elements(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Elements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDynamicArray_Size(t *testing.T) {
	type fields struct {
		capacity int
		elements []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "test get DynamicArray size",
			fields: fields{
				capacity: 2,
				elements: []interface{}{1, 2},
			},
			want: 2,
		},
		{
			name: "test get DynamicArray size",
			fields: fields{
				capacity: 0,
				elements: []interface{}{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, dynamicArray := New(tt.fields.capacity, tt.fields.elements...)
			if err != nil {
				t.Errorf("can't make new dynamic array, error: %v", err)
			}
			if got := dynamicArray.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
