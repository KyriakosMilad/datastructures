package binarysearchtree

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	n := &node{value: 3}
	type args struct {
		n *node
	}
	tests := []struct {
		name string
		args args
		want *BinarySearchTree
	}{
		{
			name: "test create new binarysearchtree",
			args: args{
				n: n,
			},
			want: &BinarySearchTree{root: n},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
