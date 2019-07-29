package priority_queue

import (
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	arr := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	queue := NewPriorityQueue(arr)
	queue.Add(10);
	t.Logf("取出最大值%+v", queue.Take())
}
