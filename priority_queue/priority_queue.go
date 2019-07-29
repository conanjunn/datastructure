package priority_queue

import (
	"datastructure/heap"
)

func NewPriorityQueue(arr []int64) *PriorityQueue {
	queue := &PriorityQueue{
		Heap: heap.NewHeap(arr),
	}
	return queue
}

type PriorityQueue struct {
	*heap.Heap
}
