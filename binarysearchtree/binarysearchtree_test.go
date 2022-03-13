package binarysearchtree

import (
	"reflect"
	"testing"
)

// tree example to use with all tests
//               25
//      20               36
//   10    22         30    40
//  5  12           28    38  48
var binarysearchtree = BinarySearchTree{
	root: &node{
		value: 25,
		right: &node{
			value: 36,
			left: &node{
				value: 30,
				left: &node{
					value: 28,
				},
			},
			right: &node{
				value: 40,
				left:  &node{value: 38},
				right: &node{value: 48},
			},
		},
		left: &node{
			value: 20,
			left: &node{
				value: 10,
				left:  &node{value: 5},
				right: &node{value: 12},
				count: 1,
			},
			right: &node{value: 22},
		},
	},
}

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

func TestBinarySearchTree_Insert(t *testing.T) {
	bstexample := &binarysearchtree
	bst := &BinarySearchTree{}
	bst.Insert(25)
	bst.Insert(20)
	bst.Insert(10)
	bst.Insert(22)
	bst.Insert(5)
	bst.Insert(12)
	bst.Insert(36)
	bst.Insert(30)
	bst.Insert(28)
	bst.Insert(40)
	bst.Insert(38)
	bst.Insert(48)
	bst.Insert(10)

	if !reflect.DeepEqual(bst, bstexample) {
		t.Errorf("New() = %v, want %v", bst, bstexample)
	}
}
