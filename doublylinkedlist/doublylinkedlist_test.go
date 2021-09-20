package doublylinkedlist

import "testing"

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
