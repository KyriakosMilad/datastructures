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

func TestBinarySearchTree_Insert(t *testing.T) {
	// tree example to use with test
	//               25
	//      20               36
	//   10    22         30    40
	//  5  12           28    38  48
	bstInsertExample := &BinarySearchTree{
		root: &node{
			value: 25,
			right: &node{
				value: 36,
				left: &node{
					value: 30,
					left:  &node{value: 28},
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

	if !reflect.DeepEqual(bst, bstInsertExample) {
		t.Errorf("Insert() = %v, want %v", bst, bstInsertExample)
	}
}

func Test_node_delete(t *testing.T) {
	type fields struct {
		value int
		left  *node
		right *node
		count int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "test delete node",
			fields: fields{
				value: 5,
				left:  &node{},
				right: &node{},
				count: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				value: tt.fields.value,
				left:  tt.fields.left,
				right: tt.fields.right,
				count: tt.fields.count,
			}
			n.delete()
			if !reflect.DeepEqual(n, &node{}) {
				t.Errorf("node delete() = %v, want %v", n, &node{})
			}
		})
	}
}

func TestBinarySearchTree_Remove(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name    string
		bst     *BinarySearchTree
		args    args
		wantErr bool
		want    *BinarySearchTree
	}{
		{
			//               25
			//      20               36
			//   10    22         30       40
			//  5  12           28      38     48
			//        14                     46
			//             will be:
			//               25
			//      20                  36
			//   10                30        40
			//  5  12           28       38      48
			//        14                       46
			name: "remove node with no child (on the right to the parent)",
			bst: &BinarySearchTree{
				root: &node{
					value: 25,
					right: &node{
						value: 36,
						left: &node{
							value: 30,
							left:  &node{value: 28},
						},
						right: &node{
							value: 40,
							left:  &node{value: 38},
							right: &node{
								value: 48,
								left:  &node{value: 46},
							},
						},
					},
					left: &node{
						value: 20,
						left: &node{
							value: 10,
							left:  &node{value: 5},
							right: &node{
								value: 12,
								right: &node{value: 14},
							},
						},
						right: &node{value: 22},
					},
				},
			},
			args:    args{val: 22},
			wantErr: false,
			want: &BinarySearchTree{
				root: &node{
					value: 25,
					right: &node{
						value: 36,
						left: &node{
							value: 30,
							left:  &node{value: 28},
						},
						right: &node{
							value: 40,
							left:  &node{value: 38},
							right: &node{
								value: 48,
								left:  &node{value: 46},
							},
						},
					},
					left: &node{
						value: 20,
						left: &node{
							value: 10,
							left:  &node{value: 5},
							right: &node{
								value: 12,
								right: &node{value: 14},
							},
						},
					},
				},
			},
		},
		{
			//               25
			//      20               36
			//   10    22         30       40
			//  5  12           28      38     48
			//        14                     46
			//             will be:
			//               25
			//      20               36
			//   10    22         30       40
			//  5  12                   38     48
			//        14                     46
			name: "remove node with no child (on the left to the parent)",
			bst: &BinarySearchTree{
				root: &node{
					value: 25,
					right: &node{
						value: 36,
						left: &node{
							value: 30,
							left:  &node{value: 28},
						},
						right: &node{
							value: 40,
							left:  &node{value: 38},
							right: &node{
								value: 48,
								left:  &node{value: 46},
							},
						},
					},
					left: &node{
						value: 20,
						left: &node{
							value: 10,
							left:  &node{value: 5},
							right: &node{
								value: 12,
								right: &node{value: 14},
							},
						},
						right: &node{value: 22},
					},
				},
			},
			args:    args{val: 28},
			wantErr: false,
			want: &BinarySearchTree{
				root: &node{
					value: 25,
					right: &node{
						value: 36,
						left: &node{
							value: 30,
						},
						right: &node{
							value: 40,
							left:  &node{value: 38},
							right: &node{
								value: 48,
								left:  &node{value: 46},
							},
						},
					},
					left: &node{
						value: 20,
						left: &node{
							value: 10,
							left:  &node{value: 5},
							right: &node{
								value: 12,
								right: &node{value: 14},
							},
						},
						right: &node{value: 22},
					},
				},
			},
		},
		{
			//               25
			//      20                36
			//   10    22         30       40
			//  5  12           28      38     48
			//        14                     46
			//             will be:
			//               25
			//      20                 38
			//   10    22         30        40
			//  5  12           28             48
			//        14                     46
			name: "remove node with two children",
			bst: &BinarySearchTree{
				root: &node{
					value: 25,
					right: &node{
						value: 36,
						left: &node{
							value: 30,
							left:  &node{value: 28},
						},
						right: &node{
							value: 40,
							left:  &node{value: 38},
							right: &node{
								value: 48,
								left:  &node{value: 46},
							},
						},
					},
					left: &node{
						value: 20,
						left: &node{
							value: 10,
							left:  &node{value: 5},
							right: &node{
								value: 12,
								right: &node{value: 14},
							},
						},
						right: &node{value: 22},
					},
				},
			},
			args:    args{val: 36},
			wantErr: false,
			want: &BinarySearchTree{
				root: &node{
					value: 25,
					right: &node{
						value: 38,
						left: &node{
							value: 30,
							left:  &node{value: 28},
						},
						right: &node{
							value: 40,
							right: &node{
								value: 48,
								left:  &node{value: 46},
							},
						},
					},
					left: &node{
						value: 20,
						left: &node{
							value: 10,
							left:  &node{value: 5},
							right: &node{
								value: 12,
								right: &node{value: 14},
							},
						},
						right: &node{value: 22},
					},
				},
			},
		},
		{
			//               25
			//      20                36
			//   10    22         30       40
			//  5  12           28      38     48
			//        14                     46
			//             will be:
			//               25
			//      20                 36
			//   10    22         30         40
			//  5  12           28        38    46
			//        14
			name: "remove node with one child on the right and on the left to the parent",
			bst: &BinarySearchTree{
				root: &node{
					value: 25,
					right: &node{
						value: 36,
						left: &node{
							value: 30,
							left:  &node{value: 28},
						},
						right: &node{
							value: 40,
							left:  &node{value: 38},
							right: &node{
								value: 48,
								left:  &node{value: 46},
							},
						},
					},
					left: &node{
						value: 20,
						left: &node{
							value: 10,
							left:  &node{value: 5},
							right: &node{
								value: 12,
								right: &node{value: 14},
							},
						},
						right: &node{value: 22},
					},
				},
			},
			args:    args{val: 48},
			wantErr: false,
			want: &BinarySearchTree{
				root: &node{
					value: 25,
					right: &node{
						value: 36,
						left: &node{
							value: 30,
							left:  &node{value: 28},
						},
						right: &node{
							value: 40,
							left:  &node{value: 38},
							right: &node{
								value: 46,
							},
						},
					},
					left: &node{
						value: 20,
						left: &node{
							value: 10,
							left:  &node{value: 5},
							right: &node{
								value: 12,
								right: &node{value: 14},
							},
						},
						right: &node{value: 22},
					},
				},
			},
		},
		{
			//               25
			//      20                36
			//   10    22         30       40
			//  5  12           28             48
			//        14
			//             will be:
			//               25
			//      20                36
			//   10    22         30       48
			//  5  12           28
			//        14
			name: "remove node with one child on the right and on the right to the parent",
			bst: &BinarySearchTree{
				root: &node{
					value: 25,
					right: &node{
						value: 36,
						left: &node{
							value: 30,
							left:  &node{value: 28},
						},
						right: &node{
							value: 40,
							right: &node{
								value: 48,
							},
						},
					},
					left: &node{
						value: 20,
						left: &node{
							value: 10,
							left:  &node{value: 5},
							right: &node{
								value: 12,
								right: &node{value: 14},
							},
						},
						right: &node{value: 22},
					},
				},
			},
			args:    args{val: 40},
			wantErr: false,
			want: &BinarySearchTree{
				root: &node{
					value: 25,
					right: &node{
						value: 36,
						left: &node{
							value: 30,
							left:  &node{value: 28},
						},
						right: &node{value: 48},
					},
					left: &node{
						value: 20,
						left: &node{
							value: 10,
							left:  &node{value: 5},
							right: &node{
								value: 12,
								right: &node{value: 14},
							},
						},
						right: &node{value: 22},
					},
				},
			},
		},
		{
			//               25
			//      20                36
			//   10    22         30       40
			//  5  12           28      38     48
			//        14                     46
			//             will be:
			//               25
			//      20                36
			//   10    22         30       40
			//  5  14           28      38     48
			//                               46
			name: "remove node with one child on the left and on the right to the parent",
			bst: &BinarySearchTree{
				root: &node{
					value: 25,
					right: &node{
						value: 36,
						left: &node{
							value: 30,
							left:  &node{value: 28},
						},
						right: &node{
							value: 40,
							left:  &node{value: 38},
							right: &node{
								value: 48,
								left:  &node{value: 46},
							},
						},
					},
					left: &node{
						value: 20,
						left: &node{
							value: 10,
							left:  &node{value: 5},
							right: &node{
								value: 12,
								right: &node{value: 14},
							},
						},
						right: &node{value: 22},
					},
				},
			},
			args:    args{val: 12},
			wantErr: false,
			want: &BinarySearchTree{
				root: &node{
					value: 25,
					right: &node{
						value: 36,
						left: &node{
							value: 30,
							left:  &node{value: 28},
						},
						right: &node{
							value: 40,
							left:  &node{value: 38},
							right: &node{
								value: 48,
								left:  &node{value: 46},
							},
						},
					},
					left: &node{
						value: 20,
						left: &node{
							value: 10,
							left:  &node{value: 5},
							right: &node{value: 14},
						},
						right: &node{value: 22},
					},
				},
			},
		},
		{
			//               25
			//      20                36
			//   10    22         30       40
			//  5               28      38     48
			//                               46
			//             will be:
			//               25
			//      20                36
			//   5     22         30       40
			//                  28      38     48
			//                               46
			name: "remove node with one child on the left and on the left to the parent",
			bst: &BinarySearchTree{
				root: &node{
					value: 25,
					right: &node{
						value: 36,
						left: &node{
							value: 30,
							left:  &node{value: 28},
						},
						right: &node{
							value: 40,
							left:  &node{value: 38},
							right: &node{
								value: 48,
								left:  &node{value: 46},
							},
						},
					},
					left: &node{
						value: 20,
						left: &node{
							value: 10,
							left:  &node{value: 5},
						},
						right: &node{value: 22},
					},
				},
			},
			args:    args{val: 10},
			wantErr: false,
			want: &BinarySearchTree{
				root: &node{
					value: 25,
					right: &node{
						value: 36,
						left: &node{
							value: 30,
							left:  &node{value: 28},
						},
						right: &node{
							value: 40,
							left:  &node{value: 38},
							right: &node{
								value: 48,
								left:  &node{value: 46},
							},
						},
					},
					left: &node{
						value: 20,
						left:  &node{value: 5},
						right: &node{value: 22},
					},
				},
			},
		},
		{
			//               25
			//      20                36
			//             will be:
			//               36
			//      20
			name: "remove root has two children",
			bst: &BinarySearchTree{
				root: &node{
					value: 25,
					right: &node{value: 36},
					left:  &node{value: 20},
				},
			},
			args:    args{val: 25},
			wantErr: false,
			want: &BinarySearchTree{
				root: &node{
					value: 36,
					left:  &node{value: 20},
				},
			},
		},
		{
			//               25
			//                        36
			//             will be:
			//               36
			name: "remove root has one child on the right",
			bst: &BinarySearchTree{
				root: &node{
					value: 25,
					right: &node{value: 36},
				},
			},
			args:    args{val: 25},
			wantErr: false,
			want: &BinarySearchTree{
				root: &node{
					value: 36,
				},
			},
		},
		{
			//               25
			//      20
			//             will be:
			//               20
			name: "remove root has one child on the left",
			bst: &BinarySearchTree{
				root: &node{
					value: 25,
					left:  &node{value: 20},
				},
			},
			args:    args{val: 25},
			wantErr: false,
			want: &BinarySearchTree{
				root: &node{
					value: 20,
				},
			},
		},
		{
			//               25
			//             will be:
			//               nil
			name: "remove root has no children",
			bst: &BinarySearchTree{
				root: &node{
					value: 25,
				},
			},
			args:    args{val: 25},
			wantErr: false,
			want: &BinarySearchTree{
				root: nil,
			},
		},
		{
			name: "remove value doesn't exist",
			bst: &BinarySearchTree{
				root: &node{
					value: 25,
					right: &node{
						value: 36,
						left: &node{
							value: 30,
							left:  &node{value: 28},
						},
						right: &node{
							value: 40,
							left:  &node{value: 38},
							right: &node{
								value: 48,
								left:  &node{value: 46},
							},
						},
					},
					left: &node{
						value: 20,
						left:  &node{value: 5},
						right: &node{value: 22},
					},
				},
			},
			args:    args{val: 50},
			wantErr: true,
			want: &BinarySearchTree{
				root: &node{
					value: 25,
					right: &node{
						value: 36,
						left: &node{
							value: 30,
							left:  &node{value: 28},
						},
						right: &node{
							value: 40,
							left:  &node{value: 38},
							right: &node{
								value: 48,
								left:  &node{value: 46},
							},
						},
					},
					left: &node{
						value: 20,
						left:  &node{value: 5},
						right: &node{value: 22},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.bst.Remove(tt.args.val); ((err != nil) != tt.wantErr) || !reflect.DeepEqual(tt.want, tt.bst) {
				t.Errorf("Remove() = %#+v, want %#+v, error = %v, wantErr %v", tt.bst, tt.want, err, tt.wantErr)
			}
		})
	}
}

func TestBinarySearchTree_Contains(t *testing.T) {
	bst := &BinarySearchTree{
		root: &node{
			value: 25,
			right: &node{
				value: 36,
				left: &node{
					value: 30,
					left:  &node{value: 28},
				},
				right: &node{
					value: 40,
					left:  &node{value: 38},
					right: &node{
						value: 48,
						left:  &node{value: 46},
					},
				},
			},
			left: &node{
				value: 20,
				left:  &node{value: 5},
				right: &node{value: 22},
			},
		},
	}

	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "check if value exists (root)",
			args: args{val: 25},
			want: true,
		},
		{
			name: "check if value exists (on the left)",
			args: args{val: 5},
			want: true,
		},
		{
			name: "check if value exists (on the right)",
			args: args{val: 38},
			want: true,
		},
		{
			name: "check if non-exist value returns false",
			args: args{val: 100},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bst.Contains(tt.args.val); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_BreadthFirstSearch(t *testing.T) {
	type fields struct {
		root *node
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			//               25
			//      20               36
			//   10    22         30       40
			//  5  12           28      38     48
			//        14                     46
			name: "test BreadthFirstSearch",
			fields: fields{
				root: &node{
					value: 25,
					right: &node{
						value: 36,
						left: &node{
							value: 30,
							left:  &node{value: 28},
						},
						right: &node{
							value: 40,
							left:  &node{value: 38},
							right: &node{
								value: 48,
								left:  &node{value: 46},
							},
						},
					},
					left: &node{
						value: 20,
						left: &node{
							value: 10,
							left:  &node{value: 5},
							right: &node{
								value: 12,
								right: &node{value: 14},
							},
						},
						right: &node{value: 22},
					},
				},
			},
			want: []int{25, 36, 20, 40, 30, 22, 10, 48, 38, 28, 12, 5, 46, 14},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bst := &BinarySearchTree{
				root: tt.fields.root,
			}
			if got := bst.BreadthFirstSearch(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BreadthFirstSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
