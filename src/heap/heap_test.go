package heap

import (
	"testing"
)

func TestHeapify(t *testing.T) {
	arr := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	heap := NewHeap(arr)
	t.Logf("%+v", heap.arr)
}

func TestAdd(t *testing.T) {
	arr := make([]int64, 0)
	heap := NewHeap(arr)
	heap.Add(1)
	heap.Add(2)
	heap.Add(3)
	heap.Add(4)
	heap.Add(5)
	heap.Take()
	t.Logf("%+v ", heap.arr)
}
