package doublylinkedlist

import (
	"reflect"
	"testing"
)

func TestDoublyLinkedList_Size(t *testing.T) {
	node1 := &Node{
		value: 1,
	}
	node2 := &Node{
		value: 2,
	}
	node3 := &Node{
		value: 3,
	}
	node4 := &Node{
		value: 4,
	}
	node5 := &Node{
		value: 5,
	}
	node3.next = node4
	node4.prev = node3
	node4.next = node5
	node5.prev = node4

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
		{
			name: "test get size from one element DoublyLinkedList",
			fields: fields{
				head: node1,
				tail: node1,
				size: 1,
			},
			want: 1,
		},
		{
			name: "test get size from two element DoublyLinkedList",
			fields: fields{
				head: node1,
				tail: node2,
				size: 2,
			},
			want: 2,
		},
		{
			name: "test get size from +2 element DoublyLinkedList",
			fields: fields{
				head: node3,
				tail: node5,
				size: 3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dll := &DoublyLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				size: tt.fields.size,
			}
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

func TestDoublyLinkedList_Tail(t *testing.T) {
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
			name: "test get the tail from DoublyLinkedList",
			fields: fields{
				tail: &Node{value: 50},
				size: 1,
			},
			errorExpected: false,
			want:          50,
		},
		{
			name: "test get the tail from empty DoublyLinkedList",
			fields: fields{
				tail: &Node{value: 50},
			},
			errorExpected: true,
			want:          nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dll := &DoublyLinkedList{
				tail: tt.fields.tail,
				size: tt.fields.size,
			}
			err, got := dll.Tail()
			if (err != nil) != tt.errorExpected {
				t.Errorf("Tail() error = %v, want %v", err, tt.errorExpected)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoublyLinkedList_IndexOfNode(t *testing.T) {
	node1 := &Node{
		value: 1,
	}
	node2 := &Node{
		value: 2,
	}
	node3 := &Node{
		value: 3,
	}
	node1.next = node2
	node2.prev = node1
	node2.next = node3
	node3.prev = node2

	type fields struct {
		head *Node
		tail *Node
		size int
	}
	type args struct {
		node *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "test get the index of node in DoublyLinkedList",
			fields: fields{
				head: node1,
				tail: node2,
				size: 3,
			},
			args: args{node: node2},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dll := &DoublyLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				size: tt.fields.size,
			}
			if got := dll.IndexOfNode(tt.args.node); got != tt.want {
				t.Errorf("IndexOfNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoublyLinkedList_ContainsNode(t *testing.T) {
	node1 := &Node{
		value: 1,
	}
	node2 := &Node{
		value: 2,
	}
	node3 := &Node{
		value: 3,
	}
	node1.next = node2
	node2.prev = node1
	node2.next = node3
	node3.prev = node2

	type fields struct {
		head *Node
		tail *Node
		size int
	}
	type args struct {
		node *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "test check if non exists node returns false",
			fields: fields{
				head: node1,
				tail: node2,
				size: 3,
			},
			args: args{node: &Node{}},
			want: false,
		},
		{
			name: "test check if exists node returns true",
			fields: fields{
				head: node1,
				tail: node2,
				size: 3,
			},
			args: args{node: node3},
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
			if got := dll.ContainsNode(tt.args.node); got != tt.want {
				t.Errorf("ContainsNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoublyLinkedList_IndexOfValue(t *testing.T) {
	node1 := &Node{
		value: 1,
	}
	node2 := &Node{
		value: 2,
	}
	node3 := &Node{
		value: 3,
	}
	node1.next = node2
	node2.prev = node1
	node2.next = node3
	node3.prev = node2

	type fields struct {
		head *Node
		tail *Node
		size int
	}
	type args struct {
		val interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "test get the index of value in DoublyLinkedList",
			fields: fields{
				head: node1,
				tail: node2,
				size: 3,
			},
			args: args{val: 2},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dll := &DoublyLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				size: tt.fields.size,
			}
			if got := dll.IndexOfValue(tt.args.val); got != tt.want {
				t.Errorf("IndexOfValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoublyLinkedList_ContainsValue(t *testing.T) {
	node1 := &Node{
		value: 1,
	}
	node2 := &Node{
		value: 2,
	}
	node3 := &Node{
		value: 3,
	}
	node1.next = node2
	node2.prev = node1
	node2.next = node3
	node3.prev = node2

	type fields struct {
		head *Node
		tail *Node
		size int
	}
	type args struct {
		val interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "test check if non-exists value exists in the DoublyLinkedList",
			fields: fields{
				head: node1,
				tail: node2,
				size: 3,
			},
			args: args{val: -1},
			want: false,
		},
		{
			name: "test check if value exists in the DoublyLinkedList",
			fields: fields{
				head: node1,
				tail: node2,
				size: 3,
			},
			args: args{val: 3},
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
			if got := dll.ContainsValue(tt.args.val); got != tt.want {
				t.Errorf("ContainsValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
