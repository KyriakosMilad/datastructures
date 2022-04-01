package binaryheap

import "errors"

type MaxBinaryHeap struct {
	list []int
}

func (mbh *MaxBinaryHeap) bubbleUp(idx int) {
	parentIdx := (idx - 1) / 2
	if mbh.list[parentIdx] < mbh.list[idx] {
		temp := mbh.list[idx]
		mbh.list[idx] = mbh.list[parentIdx]
		mbh.list[parentIdx] = temp
		mbh.bubbleUp(parentIdx)
	}
}

func (mbh *MaxBinaryHeap) Insert(val int) {
	mbh.list = append(mbh.list, val)
	mbh.bubbleUp(len(mbh.list) - 1)
}

func (mbh *MaxBinaryHeap) bubbleDown(idx int) {
	leftChildIdx := (2 * idx) + 1
	rightChildIdx := (2 * idx) + 2
	if leftChildIdx > len(mbh.list) || rightChildIdx > (len(mbh.list)-1) {
		return
	}
	bigIdx := 0
	if mbh.list[leftChildIdx] > mbh.list[rightChildIdx] {
		bigIdx = leftChildIdx
	} else {
		bigIdx = rightChildIdx
	}
	if mbh.list[bigIdx] > mbh.list[idx] {
		temp := mbh.list[idx]
		mbh.list[idx] = mbh.list[bigIdx]
		mbh.list[bigIdx] = temp
		mbh.bubbleDown(bigIdx)
	}
}

func (mbh *MaxBinaryHeap) ExtractMax() (int, error) {
	if len(mbh.list) == 0 {
		return 0, errors.New("can't extract max from empty heap")
	}
	max := mbh.list[0]
	mbh.list[0] = mbh.list[len(mbh.list)-1]
	mbh.list = mbh.list[:len(mbh.list)-1]
	mbh.bubbleDown(0)
	return max, nil
}

func (mbh *MaxBinaryHeap) Size() int {
	return len(mbh.list)
}
