package binaryheap

import (
	"reflect"
	"testing"
)

func TestMaxBinaryHeap_bubbleUp(t *testing.T) {
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
			name: "test bubbleUp",
			fields: fields{
				list: []int{41, 39, 33, 18, 27, 12, 55},
			},
			args: args{
				idx: 6,
			},
			want: []int{55, 39, 41, 18, 27, 12, 33},
		},
		{
			name: "test bubbleUp with no empty child slot",
			fields: fields{
				list: []int{41, 39, 33, 18, 27, 12, 20, 55},
			},
			args: args{
				idx: 7,
			},
			want: []int{55, 41, 33, 39, 27, 12, 20, 18},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mbh := &MaxBinaryHeap{
				list: tt.fields.list,
			}
			mbh.bubbleUp(tt.args.idx)
			if !reflect.DeepEqual(mbh.list, tt.want) {
				t.Errorf("bubbleUp() = %v, want %v", mbh.list, tt.want)
			}
		})
	}
}

func TestMaxBinaryHeap_Insert(t *testing.T) {
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
			name:   "test insert",
			fields: fields{list: []int{41, 39, 33, 18, 27, 12, 20}},
			args: args{
				val: 55,
			},
			want: []int{55, 41, 33, 39, 27, 12, 20, 18},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mbh := &MaxBinaryHeap{
				list: tt.fields.list,
			}
			mbh.Insert(tt.args.val)
			if !reflect.DeepEqual(mbh.list, tt.want) {
				t.Errorf("Insert() = %v, want %v", mbh.list, tt.want)
			}
		})
	}
}
