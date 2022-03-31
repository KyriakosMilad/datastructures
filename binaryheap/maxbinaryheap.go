package binaryheap

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
