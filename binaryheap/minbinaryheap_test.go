package binaryheap

import (
	"reflect"
	"testing"
)

func TestMinBinaryHeap_bubbleUp(t *testing.T) {
	type fields struct {
		list []int
	}
	type args struct {
		idx int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "test bubbleup",
			fields: fields{
				list: []int{1, 2, 3, 4, 5, 6, 0},
			},
			args: args{
				idx: 6,
			},
			want: []int{0, 2, 1, 4, 5, 6, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mbh := &MinBinaryHeap{
				list: tt.fields.list,
			}
			mbh.bubbleUp(tt.args.idx)
			if !reflect.DeepEqual(mbh.list, tt.want) {
				t.Errorf("bubbleUp() = %v, want %v", mbh.list, tt.want)
			}
		})
	}
}

func TestMinBinaryHeap_Insert(t *testing.T) {
	type fields struct {
		list []int
	}
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name: "test insert",
			fields: fields{
				list: []int{1, 2, 3, 4, 5, 6},
			},
			args: args{
				val: 0,
			},
			want: []int{0, 2, 1, 4, 5, 6, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mbh := &MinBinaryHeap{
				list: tt.fields.list,
			}
			mbh.Insert(tt.args.val)
			if !reflect.DeepEqual(mbh.list, tt.want) {
				t.Errorf("Insert() = %v, want %v", mbh.list, tt.want)
			}
		})
	}
}
