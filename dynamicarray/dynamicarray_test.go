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
		name                 string
		args                 args
		errorExpected        bool
		wantDynamicArray     bool
		dynamicArrayExpected DynamicArray
	}{
		{
			name: "test create new dynamicArray",
			args: args{
				capacity: 5,
				elements: []interface{}{1, 2, 3},
			},
			errorExpected:    false,
			wantDynamicArray: true,
			dynamicArrayExpected: DynamicArray{
				capacity: 5,
				size:     3,
				elements: []interface{}{1, 2, 3, nil, nil},
			},
		},
		{
			name: "test create new dynamicArray with zero elements",
			args: args{
				capacity: 3,
				elements: []interface{}{},
			},
			errorExpected:    false,
			wantDynamicArray: true,
			dynamicArrayExpected: DynamicArray{
				capacity: 3,
				size:     0,
				elements: []interface{}{nil, nil, nil},
			},
		},
		{
			name: "test create new dynamicArray with capacity lower than 0",
			args: args{
				capacity: -1,
				elements: []interface{}{1, 2, 3},
			},
			errorExpected:    true,
			wantDynamicArray: false,
			dynamicArrayExpected: DynamicArray{
				capacity: 0,
				size:     0,
				elements: nil,
			},
		},
		{
			name: "test create new dynamicArray with capacity lower than number of elements",
			args: args{
				capacity: 2,
				elements: []interface{}{1, 2, 3},
			},
			errorExpected:    true,
			wantDynamicArray: false,
			dynamicArrayExpected: DynamicArray{
				capacity: 0,
				size:     0,
				elements: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, dynamicArray := New(tt.args.capacity, tt.args.elements...)
			if (err != nil) != tt.errorExpected {
				t.Errorf("New() error = %v, want %v", err, tt.errorExpected)
			}

			if (dynamicArray != nil) != tt.wantDynamicArray {
				t.Errorf("New() dynamicArrayReturned = %v, want %v", dynamicArray != nil, tt.wantDynamicArray)
			}

			if (dynamicArray != nil) && (dynamicArray.size != tt.dynamicArrayExpected.size) && tt.wantDynamicArray {
				t.Errorf("New() wrong dynamic array size expected, got = %v, expected %v", dynamicArray.size, tt.dynamicArrayExpected.size)
			}
			if (dynamicArray != nil) && (dynamicArray.capacity != tt.dynamicArrayExpected.capacity) && tt.wantDynamicArray {
				t.Errorf("New() wrong dynamic array capacity expected, got = %v, expected %v", dynamicArray.capacity, tt.dynamicArrayExpected.capacity)
			}
			if (dynamicArray != nil) && (len(dynamicArray.elements) == len(tt.dynamicArrayExpected.elements)) && (len(dynamicArray.elements) == 0) && !reflect.DeepEqual(dynamicArray.elements, tt.dynamicArrayExpected.elements) && tt.wantDynamicArray {
				t.Errorf("New() wrong dynamic array elements expected, got = %v, expected %v", dynamicArray.elements, tt.dynamicArrayExpected.elements)
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
		name                 string
		fields               fields
		args                 args
		indexExpected        int
		wantDynamicArray     bool
		dynamicArrayExpected DynamicArray
	}{
		{
			name: "test append element with capacity room available",
			fields: fields{
				capacity: 2,
				elements: []interface{}{1},
			},
			args:             args{element: 2},
			indexExpected:    1,
			wantDynamicArray: true,
			dynamicArrayExpected: DynamicArray{
				capacity: 2,
				size:     2,
				elements: []interface{}{1, 2},
			},
		},
		{
			name: "test append element with no capacity room available",
			fields: fields{
				capacity: 2,
				elements: []interface{}{1, 2},
			},
			args:             args{element: 3},
			indexExpected:    2,
			wantDynamicArray: true,
			dynamicArrayExpected: DynamicArray{
				capacity: 4,
				size:     3,
				elements: []interface{}{1, 2, 3, nil},
			},
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

			if (dynamicArray != nil) && (dynamicArray.size != tt.dynamicArrayExpected.size) && tt.wantDynamicArray {
				t.Errorf("New() wrong dynamic array size expected, got = %v, expected %v", dynamicArray.size, tt.dynamicArrayExpected.size)
			}
			if (dynamicArray != nil) && (dynamicArray.capacity != tt.dynamicArrayExpected.capacity) && tt.wantDynamicArray {
				t.Errorf("New() wrong dynamic array capacity expected, got = %v, expected %v", dynamicArray.capacity, tt.dynamicArrayExpected.capacity)
			}
			if (dynamicArray != nil) && (len(dynamicArray.elements) == len(tt.dynamicArrayExpected.elements)) && (len(dynamicArray.elements) == 0) && !reflect.DeepEqual(dynamicArray.elements, tt.dynamicArrayExpected.elements) && tt.wantDynamicArray {
				t.Errorf("New() wrong dynamic array elements expected, got = %v, expected %v", dynamicArray.elements, tt.dynamicArrayExpected.elements)
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
		name                 string
		fields               fields
		args                 args
		wantErr              bool
		wantDynamicArray     bool
		dynamicArrayExpected DynamicArray
	}{
		{
			name: "test insert element with capacity room available",
			fields: fields{
				capacity: 3,
				elements: []interface{}{1, 3},
			},
			args:             args{index: 1, element: 2},
			wantErr:          false,
			wantDynamicArray: true,
			dynamicArrayExpected: DynamicArray{
				capacity: 3,
				size:     3,
				elements: []interface{}{1, 2, 3},
			},
		},
		{
			name: "test insert element with no capacity room available",
			fields: fields{
				capacity: 3,
				elements: []interface{}{1, 2, 4},
			},
			args:             args{index: 2, element: 3},
			wantErr:          false,
			wantDynamicArray: true,
			dynamicArrayExpected: DynamicArray{
				capacity: 6,
				size:     4,
				elements: []interface{}{1, 2, 3, 4, nil, nil},
			},
		},
		{
			name: "test insert element at the end",
			fields: fields{
				capacity: 2,
				elements: []interface{}{1},
			},
			args:             args{index: 1, element: 2},
			wantErr:          false,
			wantDynamicArray: true,
			dynamicArrayExpected: DynamicArray{
				capacity: 2,
				size:     2,
				elements: []interface{}{1, 2},
			},
		},
		{
			name: "test insert element at the beginning",
			fields: fields{
				capacity: 2,
				elements: []interface{}{2},
			},
			args:             args{index: 0, element: 1},
			wantErr:          false,
			wantDynamicArray: true,
			dynamicArrayExpected: DynamicArray{
				capacity: 2,
				size:     2,
				elements: []interface{}{1, 2},
			},
		},
		{
			name: "test insert element with index lower than 0",
			fields: fields{
				capacity: 2,
				elements: []interface{}{2},
			},
			args:             args{index: -1, element: 1},
			wantErr:          true,
			wantDynamicArray: false,
			dynamicArrayExpected: DynamicArray{
				capacity: 0,
				size:     0,
				elements: nil,
			},
		},
		{
			name: "test insert element with index bigger than the DynamicArray size",
			fields: fields{
				capacity: 2,
				elements: []interface{}{2},
			},
			args:             args{index: 2, element: 1},
			wantErr:          true,
			wantDynamicArray: false,
			dynamicArrayExpected: DynamicArray{
				capacity: 0,
				size:     0,
				elements: nil,
			},
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

			if (dynamicArray != nil) && (dynamicArray.size != tt.dynamicArrayExpected.size) && tt.wantDynamicArray {
				t.Errorf("New() wrong dynamic array size expected, got = %v, expected %v", dynamicArray.size, tt.dynamicArrayExpected.size)
			}
			if (dynamicArray != nil) && (dynamicArray.capacity != tt.dynamicArrayExpected.capacity) && tt.wantDynamicArray {
				t.Errorf("New() wrong dynamic array capacity expected, got = %v, expected %v", dynamicArray.capacity, tt.dynamicArrayExpected.capacity)
			}
			if (dynamicArray != nil) && (len(dynamicArray.elements) == len(tt.dynamicArrayExpected.elements)) && (len(dynamicArray.elements) == 0) && !reflect.DeepEqual(dynamicArray.elements, tt.dynamicArrayExpected.elements) && tt.wantDynamicArray {
				t.Errorf("New() wrong dynamic array elements expected, got = %v, expected %v", dynamicArray.elements, tt.dynamicArrayExpected.elements)
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
		name                 string
		fields               fields
		args                 args
		wantErr              bool
		wantDynamicArray     bool
		dynamicArrayExpected DynamicArray
	}{
		{
			name: "test remove element at specific index",
			fields: fields{
				capacity: 2,
				elements: []interface{}{1, 2},
			},
			args:             args{index: 0},
			wantErr:          false,
			wantDynamicArray: true,
			dynamicArrayExpected: DynamicArray{
				capacity: 1,
				size:     1,
				elements: []interface{}{2},
			},
		},
		{
			name: "test remove element at index lower than 0",
			fields: fields{
				capacity: 2,
				elements: []interface{}{1, 2},
			},
			args:             args{index: -1},
			wantErr:          true,
			wantDynamicArray: false,
			dynamicArrayExpected: DynamicArray{
				capacity: 0,
				size:     0,
				elements: nil,
			},
		},
		{
			name: "test remove element at index bigger than the DynamicArray size",
			fields: fields{
				capacity: 2,
				elements: []interface{}{1, 2},
			},
			args:             args{index: 3},
			wantErr:          true,
			wantDynamicArray: false,
			dynamicArrayExpected: DynamicArray{
				capacity: 0,
				size:     0,
				elements: nil,
			},
		},
		{
			name: "test remove element from empty array",
			fields: fields{
				capacity: 2,
				elements: []interface{}{},
			},
			args:             args{index: 3},
			wantErr:          true,
			wantDynamicArray: false,
			dynamicArrayExpected: DynamicArray{
				capacity: 0,
				size:     0,
				elements: nil,
			},
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

			if (dynamicArray != nil) && (dynamicArray.size != tt.dynamicArrayExpected.size) && tt.wantDynamicArray {
				t.Errorf("New() wrong dynamic array size expected, got = %v, expected %v", dynamicArray.size, tt.dynamicArrayExpected.size)
			}
			if (dynamicArray != nil) && (dynamicArray.capacity != tt.dynamicArrayExpected.capacity) && tt.wantDynamicArray {
				t.Errorf("New() wrong dynamic array capacity expected, got = %v, expected %v", dynamicArray.capacity, tt.dynamicArrayExpected.capacity)
			}
			if (dynamicArray != nil) && (len(dynamicArray.elements) == len(tt.dynamicArrayExpected.elements)) && (len(dynamicArray.elements) == 0) && !reflect.DeepEqual(dynamicArray.elements, tt.dynamicArrayExpected.elements) && tt.wantDynamicArray {
				t.Errorf("New() wrong dynamic array elements expected, got = %v, expected %v", dynamicArray.elements, tt.dynamicArrayExpected.elements)
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
		name                 string
		fields               fields
		args                 args
		wantErr              bool
		wantDynamicArray     bool
		dynamicArrayExpected DynamicArray
	}{
		{
			name: "test remove multiple elements from DynamicArray",
			fields: fields{
				capacity: 4,
				elements: []interface{}{1, 2, 2},
			},
			args:             args{element: 2},
			wantErr:          false,
			wantDynamicArray: true,
			dynamicArrayExpected: DynamicArray{
				capacity: 1,
				size:     1,
				elements: []interface{}{1},
			},
		},
		{
			name: "test remove one element from DynamicArray using RemoveAllWhere() method",
			fields: fields{
				capacity: 2,
				elements: []interface{}{1, 2},
			},
			args:             args{element: 2},
			wantErr:          false,
			wantDynamicArray: true,
			dynamicArrayExpected: DynamicArray{
				capacity: 1,
				size:     1,
				elements: []interface{}{1},
			},
		},
		{
			name: "test remove element that does not exist using RemoveAllWhere() method",
			fields: fields{
				capacity: 2,
				elements: []interface{}{1, 2},
			},
			args:             args{element: 3},
			wantErr:          false,
			wantDynamicArray: true,
			dynamicArrayExpected: DynamicArray{
				capacity: 2,
				size:     2,
				elements: []interface{}{1, 2},
			},
		},
		{
			name: "test remove element from empty array",
			fields: fields{
				capacity: 2,
				elements: []interface{}{},
			},
			args:             args{element: 3},
			wantErr:          true,
			wantDynamicArray: false,
			dynamicArrayExpected: DynamicArray{
				capacity: 0,
				size:     0,
				elements: nil,
			},
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

			if (dynamicArray != nil) && (dynamicArray.size != tt.dynamicArrayExpected.size) && tt.wantDynamicArray {
				t.Errorf("New() wrong dynamic array size expected, got = %v, expected %v", dynamicArray.size, tt.dynamicArrayExpected.size)
			}
			if (dynamicArray != nil) && (dynamicArray.capacity != tt.dynamicArrayExpected.capacity) && tt.wantDynamicArray {
				t.Errorf("New() wrong dynamic array capacity expected, got = %v, expected %v", dynamicArray.capacity, tt.dynamicArrayExpected.capacity)
			}
			if (dynamicArray != nil) && (len(dynamicArray.elements) == len(tt.dynamicArrayExpected.elements)) && (len(dynamicArray.elements) == 0) && !reflect.DeepEqual(dynamicArray.elements, tt.dynamicArrayExpected.elements) && tt.wantDynamicArray {
				t.Errorf("New() wrong dynamic array elements expected, got = %v, expected %v", dynamicArray.elements, tt.dynamicArrayExpected.elements)
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
		name                 string
		fields               fields
		args                 args
		wantErr              bool
		wantDynamicArray     bool
		dynamicArrayExpected DynamicArray
	}{
		{
			name: "test remove one element from DynamicArray",
			fields: fields{
				capacity: 3,
				elements: []interface{}{1, 2, 2},
			},
			args:             args{element: 2},
			wantErr:          false,
			wantDynamicArray: true,
			dynamicArrayExpected: DynamicArray{
				capacity: 2,
				size:     2,
				elements: []interface{}{1, 2},
			},
		},
		{
			name: "test remove element that does not exist using RemoveFirstWhere() method",
			fields: fields{
				capacity: 2,
				elements: []interface{}{1, 2},
			},
			args:             args{element: 3},
			wantErr:          false,
			wantDynamicArray: true,
			dynamicArrayExpected: DynamicArray{
				capacity: 2,
				size:     2,
				elements: []interface{}{1, 2},
			},
		},
		{
			name: "test remove element from empty array",
			fields: fields{
				capacity: 2,
				elements: []interface{}{},
			},
			args:             args{element: 3},
			wantErr:          true,
			wantDynamicArray: false,
			dynamicArrayExpected: DynamicArray{
				capacity: 0,
				size:     0,
				elements: nil,
			},
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

			if (dynamicArray != nil) && (dynamicArray.size != tt.dynamicArrayExpected.size) && tt.wantDynamicArray {
				t.Errorf("New() wrong dynamic array size expected, got = %v, expected %v", dynamicArray.size, tt.dynamicArrayExpected.size)
			}
			if (dynamicArray != nil) && (dynamicArray.capacity != tt.dynamicArrayExpected.capacity) && tt.wantDynamicArray {
				t.Errorf("New() wrong dynamic array capacity expected, got = %v, expected %v", dynamicArray.capacity, tt.dynamicArrayExpected.capacity)
			}
			if (dynamicArray != nil) && (len(dynamicArray.elements) == len(tt.dynamicArrayExpected.elements)) && (len(dynamicArray.elements) == 0) && !reflect.DeepEqual(dynamicArray.elements, tt.dynamicArrayExpected.elements) && tt.wantDynamicArray {
				t.Errorf("New() wrong dynamic array elements expected, got = %v, expected %v", dynamicArray.elements, tt.dynamicArrayExpected.elements)
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
		name                 string
		fields               fields
		wantErr              bool
		wantDynamicArray     bool
		dynamicArrayExpected DynamicArray
	}{
		{
			name: "test remove last element from DynamicArray",
			fields: fields{
				capacity: 3,
				elements: []interface{}{1, 2, 2},
			},
			wantErr:          false,
			wantDynamicArray: true,
			dynamicArrayExpected: DynamicArray{
				capacity: 2,
				size:     2,
				elements: []interface{}{1, 2},
			},
		},
		{
			name: "test remove element from empty array",
			fields: fields{
				capacity: 2,
				elements: []interface{}{},
			},
			wantErr:          true,
			wantDynamicArray: false,
			dynamicArrayExpected: DynamicArray{
				capacity: 0,
				size:     0,
				elements: nil,
			},
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

			if (dynamicArray != nil) && (dynamicArray.size != tt.dynamicArrayExpected.size) && tt.wantDynamicArray {
				t.Errorf("New() wrong dynamic array size expected, got = %v, expected %v", dynamicArray.size, tt.dynamicArrayExpected.size)
			}
			if (dynamicArray != nil) && (dynamicArray.capacity != tt.dynamicArrayExpected.capacity) && tt.wantDynamicArray {
				t.Errorf("New() wrong dynamic array capacity expected, got = %v, expected %v", dynamicArray.capacity, tt.dynamicArrayExpected.capacity)
			}
			if (dynamicArray != nil) && (len(dynamicArray.elements) == len(tt.dynamicArrayExpected.elements)) && (len(dynamicArray.elements) == 0) && !reflect.DeepEqual(dynamicArray.elements, tt.dynamicArrayExpected.elements) && tt.wantDynamicArray {
				t.Errorf("New() wrong dynamic array elements expected, got = %v, expected %v", dynamicArray.elements, tt.dynamicArrayExpected.elements)
			}
		})
	}
}

func TestDynamicArray_Elements(t *testing.T) {
	type fields struct {
		capacity int
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

func TestDynamicArray_IsEmpty(t *testing.T) {
	type fields struct {
		capacity int
		elements []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "test check if two elements DynamicArray is empty",
			fields: fields{
				capacity: 2,
				elements: []interface{}{1, 2},
			},
			want: false,
		},
		{
			name: "test check if zero elements DynamicArray is empty",
			fields: fields{
				capacity: 0,
				elements: []interface{}{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, dynamicArray := New(tt.fields.capacity, tt.fields.elements...)
			if err != nil {
				t.Errorf("can't make new dynamic array, error: %v", err)
			}
			if got := dynamicArray.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDynamicArray_Capacity(t *testing.T) {
	type fields struct {
		capacity int
		size     int
		elements []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "test get DynamicArray capacity",
			fields: fields{
				capacity: 2,
				elements: []interface{}{1, 2},
			},
			want: 2,
		},
		{
			name: "test get DynamicArray capacity",
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
			if got := dynamicArray.Capacity(); got != tt.want {
				t.Errorf("Capacity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDynamicArray_Get(t *testing.T) {
	type fields struct {
		capacity int
		elements []interface{}
	}
	type args struct {
		index int
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		want          interface{}
		errorExpected bool
	}{
		{
			name: "test get element from DynamicArray by index",
			fields: fields{
				capacity: 3,
				elements: []interface{}{1, 2, 3},
			},
			args:          args{index: 0},
			want:          1,
			errorExpected: false,
		},
		{
			name: "test get element from DynamicArray by index lower than 0",
			fields: fields{
				capacity: 3,
				elements: []interface{}{1, 2, 3},
			},
			args:          args{index: -1},
			want:          nil,
			errorExpected: true,
		},
		{
			name: "test get element from DynamicArray by index bigger than DynamicArray size",
			fields: fields{
				capacity: 3,
				elements: []interface{}{1, 2, 3},
			},
			args:          args{index: 3},
			want:          nil,
			errorExpected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, dynamicArray := New(tt.fields.capacity, tt.fields.elements...)
			if err != nil {
				t.Errorf("can't make new dynamic array, error: %v", err)
			}
			err, value := dynamicArray.Get(tt.args.index)
			if (err != nil) != tt.errorExpected {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.errorExpected)
			}
			if !reflect.DeepEqual(value, tt.want) {
				t.Errorf("Get() got = %v, want %v", value, tt.want)
			}
		})
	}
}

func TestDynamicArray_Set(t *testing.T) {
	type fields struct {
		capacity int
		elements []interface{}
	}
	type args struct {
		index int
		value interface{}
	}
	tests := []struct {
		name                 string
		fields               fields
		args                 args
		wantErr              bool
		wantDynamicArray     bool
		dynamicArrayExpected DynamicArray
	}{
		{
			name: "test set element value in DynamicArray by index",
			fields: fields{
				capacity: 3,
				elements: []interface{}{1, 2, 3},
			},
			args:             args{index: 0, value: 0},
			wantErr:          false,
			wantDynamicArray: true,
			dynamicArrayExpected: DynamicArray{
				capacity: 3,
				size:     3,
				elements: []interface{}{0, 2, 3},
			},
		},
		{
			name: "test set element value in DynamicArray by index lower than 0",
			fields: fields{
				capacity: 3,
				elements: []interface{}{1, 2, 3},
			},
			args:             args{index: -1, value: 0},
			wantErr:          true,
			wantDynamicArray: false,
			dynamicArrayExpected: DynamicArray{
				capacity: 0,
				size:     0,
				elements: nil,
			},
		},
		{
			name: "test set element value in DynamicArray by index bigger than DynamicArray capacity",
			fields: fields{
				capacity: 3,
				elements: []interface{}{1, 2, 3},
			},
			args:             args{index: 3, value: 0},
			wantErr:          true,
			wantDynamicArray: false,
			dynamicArrayExpected: DynamicArray{
				capacity: 0,
				size:     0,
				elements: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, dynamicArray := New(tt.fields.capacity, tt.fields.elements...)
			if err != nil {
				t.Errorf("can't make new dynamic array, error: %v", err)
			}

			if err := dynamicArray.Set(tt.args.index, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}

			if (dynamicArray != nil) && (dynamicArray.size != tt.dynamicArrayExpected.size) && tt.wantDynamicArray {
				t.Errorf("New() wrong dynamic array size expected, got = %v, expected %v", dynamicArray.size, tt.dynamicArrayExpected.size)
			}
			if (dynamicArray != nil) && (dynamicArray.capacity != tt.dynamicArrayExpected.capacity) && tt.wantDynamicArray {
				t.Errorf("New() wrong dynamic array capacity expected, got = %v, expected %v", dynamicArray.capacity, tt.dynamicArrayExpected.capacity)
			}
			if (dynamicArray != nil) && (len(dynamicArray.elements) == len(tt.dynamicArrayExpected.elements)) && (len(dynamicArray.elements) == 0) && !reflect.DeepEqual(dynamicArray.elements, tt.dynamicArrayExpected.elements) && tt.wantDynamicArray {
				t.Errorf("New() wrong dynamic array elements expected, got = %v, expected %v", dynamicArray.elements, tt.dynamicArrayExpected.elements)
			}
		})
	}
}

func TestDynamicArray_IndexOf(t *testing.T) {
	type fields struct {
		capacity int
		elements []interface{}
	}
	type args struct {
		element interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "test get index of value in DynamicArray",
			fields: fields{
				capacity: 3,
				elements: []interface{}{1, 2, 3},
			},
			args: args{element: 3},
			want: 2,
		},
		{
			name: "test get index of value does not exist in DynamicArray",
			fields: fields{
				capacity: 3,
				elements: []interface{}{1, 2, 3},
			},
			args: args{element: 4},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, dynamicArray := New(tt.fields.capacity, tt.fields.elements...)
			if err != nil {
				t.Errorf("can't make new dynamic array, error: %v", err)
			}
			if got := dynamicArray.IndexOf(tt.args.element); got != tt.want {
				t.Errorf("IndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDynamicArray_Contains(t *testing.T) {
	type fields struct {
		capacity int
		elements []interface{}
	}
	type args struct {
		element interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "test check if DynamicArray contains element that exists",
			fields: fields{
				capacity: 2,
				elements: []interface{}{1, 2},
			},
			args: args{element: 2},
			want: true,
		},
		{
			name: "test check if DynamicArray contains element that does not exist",
			fields: fields{
				capacity: 2,
				elements: []interface{}{1, 2},
			},
			args: args{element: 3},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, dynamicArray := New(tt.fields.capacity, tt.fields.elements...)
			if err != nil {
				t.Errorf("can't make new dynamic array, error: %v", err)
			}
			if got := dynamicArray.Contains(tt.args.element); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDynamicArray_Clear(t *testing.T) {
	type fields struct {
		capacity int
		elements []interface{}
	}
	tests := []struct {
		name                 string
		fields               fields
		wantDynamicArray     bool
		dynamicArrayExpected DynamicArray
	}{
		{
			name: "test clear DynamicArray",
			fields: fields{
				capacity: 2,
				elements: []interface{}{1, 2},
			},
			wantDynamicArray: true,
			dynamicArrayExpected: DynamicArray{
				capacity: 2,
				size:     0,
				elements: []interface{}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err, dynamicArray := New(tt.fields.capacity, tt.fields.elements...)
			if err != nil {
				t.Errorf("can't make new dynamic array, error: %v", err)
			}

			dynamicArray.Clear()

			if (dynamicArray != nil) && (dynamicArray.size != tt.dynamicArrayExpected.size) && tt.wantDynamicArray {
				t.Errorf("New() wrong dynamic array size expected, got = %v, expected %v", dynamicArray.size, tt.dynamicArrayExpected.size)
			}
			if (dynamicArray != nil) && (dynamicArray.capacity != tt.dynamicArrayExpected.capacity) && tt.wantDynamicArray {
				t.Errorf("New() wrong dynamic array capacity expected, got = %v, expected %v", dynamicArray.capacity, tt.dynamicArrayExpected.capacity)
			}
			if (dynamicArray != nil) && (len(dynamicArray.elements) == len(tt.dynamicArrayExpected.elements)) && (len(dynamicArray.elements) == 0) && !reflect.DeepEqual(dynamicArray.elements, tt.dynamicArrayExpected.elements) && tt.wantDynamicArray {
				t.Errorf("New() wrong dynamic array elements expected, got = %v, expected %v", dynamicArray.elements, tt.dynamicArrayExpected.elements)
			}
		})
	}
}
