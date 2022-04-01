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

func TestMaxBinaryHeap_bubbleDown(t *testing.T) {
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
			name:   "test bubbleDown",
			fields: fields{list: []int{33, 39, 41, 18, 27, 12}},
			args: args{
				idx: 0,
			},
			want: []int{41, 39, 33, 18, 27, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mbh := &MaxBinaryHeap{
				list: tt.fields.list,
			}
			mbh.bubbleDown(tt.args.idx)
			if !reflect.DeepEqual(mbh.list, tt.want) {
				t.Errorf("bubbleDown() = %v, want %v", mbh.list, tt.want)
			}
		})
	}
}

func TestMaxBinaryHeap_ExtractMax(t *testing.T) {
	type fields struct {
		list []int
	}
	tests := []struct {
		name       string
		fields     fields
		wantedList []int
		want       int
		wantErr    bool
	}{
		{
			name:       "test ExtractMax",
			fields:     fields{list: []int{55, 39, 41, 18, 27, 12, 33}},
			wantedList: []int{41, 39, 33, 18, 27, 12},
			want:       55,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mbh := &MaxBinaryHeap{
				list: tt.fields.list,
			}
			max, err := mbh.ExtractMax()
			if err != nil != tt.wantErr {
				t.Errorf("ExtractMax() err = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(max, tt.want) {
				t.Errorf("ExtractMax() = %v, want %v", max, tt.want)
			}
			if !reflect.DeepEqual(mbh.list, tt.wantedList) {
				t.Errorf("ExtractMax() list = %v, want %v", mbh.list, tt.wantedList)
			}
		})
	}
}

func TestMaxBinaryHeap_Size(t *testing.T) {
	type fields struct {
		list []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "test get max binary heap size",
			fields: fields{
				list: []int{41, 39, 33, 18, 27, 12},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mbh := &MaxBinaryHeap{
				list: tt.fields.list,
			}
			if got := mbh.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
