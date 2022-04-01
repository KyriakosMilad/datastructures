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

func TestMinBinaryHeap_bubbleDown(t *testing.T) {
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
			name: "test bubbledown",
			fields: fields{
				list: []int{9, 2, 1, 4, 5, 6, 3},
			},
			args: args{
				idx: 0,
			},
			want: []int{1, 2, 3, 4, 5, 6, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mbh := &MinBinaryHeap{
				list: tt.fields.list,
			}
			mbh.bubbleDown(tt.args.idx)
			if !reflect.DeepEqual(mbh.list, tt.want) {
				t.Errorf("bubbleDown() = %v, want %v", mbh.list, tt.want)
			}
		})
	}
}

func TestMinBinaryHeap_ExtractMin(t *testing.T) {
	type fields struct {
		list []int
	}
	tests := []struct {
		name       string
		fields     fields
		want       int
		wantedList []int
		wantErr    bool
	}{
		{
			name: "test ExtractMin",
			fields: fields{
				list: []int{1, 2, 3, 4, 5, 6, 9},
			},
			want:       1,
			wantedList: []int{2, 4, 3, 9, 5, 6},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mbh := &MinBinaryHeap{
				list: tt.fields.list,
			}
			min, err := mbh.ExtractMin()
			if err != nil != tt.wantErr {
				t.Errorf("ExtractMax() err = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(min, tt.want) {
				t.Errorf("ExtractMin() = %v, want %v", min, tt.want)
			}
			if !reflect.DeepEqual(mbh.list, tt.wantedList) {
				t.Errorf("ExtractMin() list = %v, want %v", mbh.list, tt.wantedList)
			}
		})
	}
}

func TestMinBinaryHeap_Size(t *testing.T) {
	type fields struct {
		list []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "test get min binary heap size",
			fields: fields{
				list: []int{2, 4, 3, 9, 5, 6},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mbh := &MinBinaryHeap{
				list: tt.fields.list,
			}
			if got := mbh.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
