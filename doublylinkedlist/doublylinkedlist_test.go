package doublylinkedlist

import (
	"reflect"
	"testing"
)

func TestDoublyLinkedList_Size(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		size int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "test get size from empty DoublyLinkedList",
			fields: fields{
				head: nil,
				tail: nil,
				size: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dll := &DoublyLinkedList{}
			if got := dll.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoublyLinkedList_IsEmpty(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		size int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "test check if empty DoublyLinkedList is empty",
			fields: fields{
				head: nil,
				tail: nil,
				size: 0,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dll := &DoublyLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				size: tt.fields.size,
			}
			if got := dll.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoublyLinkedList_AddFirst(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test add node at the beginning of the DoublyLinkedList",
			args: args{val: 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dll := &DoublyLinkedList{}
			dll.AddFirst(tt.args.val)
			if dll.head.value != tt.args.val {
				t.Errorf("AddFirst() = %v, want %v", dll.head.value, tt.args.val)
			}
		})
	}
}

func TestDoublyLinkedList_AddLast(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test add node at the end of the DoublyLinkedList",
			args: args{val: 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dll := &DoublyLinkedList{}
			dll.AddFirst(tt.args.val)
			if dll.tail.value != tt.args.val {
				t.Errorf("AddLast() = %v, want %v", dll.tail.value, tt.args.val)
			}
		})
	}
}

func TestDoublyLinkedList_Head(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		size int
	}
	tests := []struct {
		name          string
		fields        fields
		errorExpected bool
		want          interface{}
	}{
		{
			name: "test get the head from DoublyLinkedList",
			fields: fields{
				head: &Node{value: 50},
				size: 1,
			},
			errorExpected: false,
			want:          50,
		},
		{
			name: "test get the head from empty DoublyLinkedList",
			fields: fields{
				head: &Node{value: 50},
			},
			errorExpected: true,
			want:          nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dll := &DoublyLinkedList{
				head: tt.fields.head,
				size: tt.fields.size,
			}
			err, got := dll.Head()
			if (err != nil) != tt.errorExpected {
				t.Errorf("Head() error = %v, want %v", err, tt.errorExpected)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Head() got = %v, want %v", got, tt.want)
			}
		})
	}
}
