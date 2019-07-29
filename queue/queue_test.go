package queue

import (
	"testing"
)

func TestEnqueue(t *testing.T) {
	q := &Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Enqueue(6)
	q.Enqueue(7)
	q.Enqueue(8)
	q.Enqueue(9)
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Enqueue(10)
	q.Enqueue(11)
	q.Enqueue(12)
	q.Enqueue(13)
	q.Enqueue(14)
	q.ToString()

	for i := 0; i < 10; i++ {
		t.Logf("%d ", q.arr[i])
	}
	t.Logf("head: %d", q.head)
	t.Logf("tail: %d", q.tail)
}
