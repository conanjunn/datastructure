package queue

import (
	"fmt"
)

type Queue struct {
	head int64
	tail int64
	size int64
	arr  [10]interface{} // 暂时先写死10
}

func (this *Queue) Size() int64 {
	return this.size
}

func (this *Queue) nextIndex(index int64) int64 {
	if index == 9 {
		return 0
	}
	return index + 1
}

func (this *Queue) Enqueue(val interface{}) {
	if this.size+1 == 11 {
		fmt.Printf("%s\n", "队列已满")
		return
	}
	if this.size == 0 {
		this.arr[this.tail] = val
	} else {
		this.tail = this.nextIndex(this.tail)
		this.arr[this.tail] = val
	}
	this.size++
}

func (this *Queue) Dequeue() interface{} {
	if this.size == 0 {
		fmt.Printf("%s\n", "队列为空")
		return 0
	}
	tmp := this.arr[this.head]
	this.arr[this.head] = nil
	if this.size != 1 {
		this.head = this.nextIndex(this.head)
	}
	this.size--
	return tmp
}

func (this *Queue) ToString() {
	cursor := this.head
	var i int64
	for i = 0; i < this.size; i++ {
		fmt.Printf("%v ", this.arr[cursor])
		cursor = this.nextIndex(cursor)
	}
	fmt.Printf("\n")
}
