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
			name: "test get size from one node DoublyLinkedList",
			fields: fields{
				head: node1,
				tail: node1,
				size: 1,
			},
			want: 1,
		},
		{
			name: "test get size from two nodes DoublyLinkedList",
			fields: fields{
				head: node1,
				tail: node2,
				size: 2,
			},
			want: 2,
		},
		{
			name: "test get size from +2 nodes DoublyLinkedList",
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
		want   bool
	}{
		{
			name: "test get size from empty DoublyLinkedList",
			fields: fields{
				head: nil,
				tail: nil,
				size: 0,
			},
			want: true,
		},
		{
			name: "test get size from one node DoublyLinkedList",
			fields: fields{
				head: node1,
				tail: node1,
				size: 1,
			},
			want: false,
		},
		{
			name: "test get size from two nodes DoublyLinkedList",
			fields: fields{
				head: node1,
				tail: node2,
				size: 2,
			},
			want: false,
		},
		{
			name: "test get size from +2 nodes DoublyLinkedList",
			fields: fields{
				head: node3,
				tail: node5,
				size: 3,
			},
			want: false,
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
	node1 := &Node{
		value: 1,
	}

	type args struct {
		val interface{}
	}
	tests := []struct {
		name        string
		args        args
		dll         DoublyLinkedList
		expectedDll DoublyLinkedList
	}{
		{
			name: "test add node at the beginning of empty DoublyLinkedList",
			args: args{val: 0},
			dll:  DoublyLinkedList{},
			expectedDll: DoublyLinkedList{
				head: &Node{value: 0},
				tail: &Node{value: 0},
				size: 1,
			},
		},
		{
			name: "test add node at the beginning of one node DoublyLinkedList",
			args: args{val: 0},
			dll: DoublyLinkedList{
				head: node1,
				tail: node1,
				size: 1,
			},
			expectedDll: DoublyLinkedList{
				head: &Node{value: 0},
				tail: &Node{value: 1},
				size: 2,
			},
		},
		{
			name: "test add node at the beginning of two nodes DoublyLinkedList",
			args: args{val: 2},
			dll: DoublyLinkedList{
				head: &Node{value: 3},
				tail: &Node{value: 6},
				size: 2,
			},
			expectedDll: DoublyLinkedList{
				head: &Node{value: 2},
				tail: &Node{value: 6},
				size: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.dll.AddFirst(tt.args.val)
			if tt.dll.head.value != tt.expectedDll.head.value {
				t.Errorf("AddFirst(): head = %v, want %v", tt.dll.head.value, tt.expectedDll.head.value)
			}
			if tt.dll.tail.value != tt.expectedDll.tail.value {
				t.Errorf("AddFirst(): tail = %v, want %v", tt.dll.tail.value, tt.expectedDll.tail.value)
			}
			if tt.dll.size != tt.expectedDll.size {
				t.Errorf("AddFirst(): size = %v, want %v", tt.dll.size, tt.expectedDll.size)
			}
		})
	}
}

func TestDoublyLinkedList_AddLast(t *testing.T) {
	node1 := &Node{
		value: 1,
	}

	type args struct {
		val interface{}
	}
	tests := []struct {
		name        string
		args        args
		dll         DoublyLinkedList
		expectedDll DoublyLinkedList
	}{
		{
			name: "test add node at the end of empty DoublyLinkedList",
			args: args{val: 0},
			dll:  DoublyLinkedList{},
			expectedDll: DoublyLinkedList{
				head: &Node{value: 0},
				tail: &Node{value: 0},
				size: 1,
			},
		},
		{
			name: "test add node at the end of one node DoublyLinkedList",
			args: args{val: 2},
			dll: DoublyLinkedList{
				head: node1,
				tail: node1,
				size: 1,
			},
			expectedDll: DoublyLinkedList{
				head: &Node{value: 1},
				tail: &Node{value: 2},
				size: 2,
			},
		},
		{
			name: "test add node at the end of two nodes DoublyLinkedList",
			args: args{val: 7},
			dll: DoublyLinkedList{
				head: &Node{value: 3},
				tail: &Node{value: 6},
				size: 2,
			},
			expectedDll: DoublyLinkedList{
				head: &Node{value: 3},
				tail: &Node{value: 7},
				size: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.dll.AddLast(tt.args.val)
			if tt.dll.head.value != tt.expectedDll.head.value {
				t.Errorf("AddFirst(): head = %v, want %v", tt.dll.head.value, tt.expectedDll.head.value)
			}
			if tt.dll.tail.value != tt.expectedDll.tail.value {
				t.Errorf("AddFirst(): tail = %v, want %v", tt.dll.tail.value, tt.expectedDll.tail.value)
			}
			if tt.dll.size != tt.expectedDll.size {
				t.Errorf("AddFirst(): size = %v, want %v", tt.dll.size, tt.expectedDll.size)
			}
		})
	}
}

func TestDoublyLinkedList_Head(t *testing.T) {
	node1 := &Node{
		value: 1,
	}

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
			name:          "test get the head from empty DoublyLinkedList",
			fields:        fields{},
			errorExpected: true,
			want:          nil,
		},
		{
			name: "test get the head from one node DoublyLinkedList",
			fields: fields{
				head: node1,
				tail: node1,
				size: 1,
			},
			want:          1,
			errorExpected: false,
		},
		{
			name: "test get the head from two nodes DoublyLinkedList",
			fields: fields{
				head: &Node{value: 3},
				tail: &Node{value: 6},
				size: 2,
			},
			want:          3,
			errorExpected: false,
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
	node1 := &Node{
		value: 1,
	}

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
			name:          "test get the tail from empty DoublyLinkedList",
			fields:        fields{},
			errorExpected: true,
			want:          nil,
		},
		{
			name: "test get the tail from one node DoublyLinkedList",
			fields: fields{
				head: node1,
				tail: node1,
				size: 1,
			},
			want:          1,
			errorExpected: false,
		},
		{
			name: "test get the tail from two nodes DoublyLinkedList",
			fields: fields{
				head: &Node{value: 3},
				tail: &Node{value: 6},
				size: 2,
			},
			want:          6,
			errorExpected: false,
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
