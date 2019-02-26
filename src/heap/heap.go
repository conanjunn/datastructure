package heap

import ()

func NewHeap(arr []int64) *Heap {
	heap := &Heap{
		arr: arr,
	}
	heap.heapify()
	return heap
}

type Heap struct {
	arr []int64
}

func (this *Heap) swap(index1, index2 int64) {
	tmp := this.arr[index1]
	this.arr[index1] = this.arr[index2]
	this.arr[index2] = tmp
}

func (this *Heap) parentIndex(index int64) int64 {
	return (index - 1) / 2
}

func (this *Heap) leftIndex(index int64) int64 {
	return (index * 2) + 1
}

func (this *Heap) rightIndex(index int64) int64 {
	return (index * 2) + 2
}

func (this *Heap) heapify() {
	length := int64(len(this.arr))
	if length == 0 {
		return
	}
	// 先找到最后一个元素的父元素，即：最后一个非叶子节点
	target := this.parentIndex(length - 1)
	for i := target; i >= 0; i-- {
		// 从这个节点开始，不断的做下沉操作
		this.goDown(i)
	}
}

// 添加元素后做上浮操作
func (this *Heap) goUp(index int64) {
	pIndex := this.parentIndex(index)
	if this.arr[index] > this.arr[pIndex] {
		this.swap(index, pIndex)
		this.goUp(pIndex)
	}
}

// 取出堆顶元素后，将最后一个元素添加到堆顶，做下沉操作
func (this *Heap) goDown(index int64) {
	left := this.leftIndex(index)
	right := this.rightIndex(index)
	if left >= int64(len(this.arr)) {
		return
	}
	// 选出左孩子和右孩子中较大的一个
	baseIndex := left
	if right < int64(len(this.arr)) && this.arr[left] < this.arr[right] {
		baseIndex = right
	}
	if this.arr[index] < this.arr[baseIndex] {
		this.swap(index, baseIndex)
		this.goDown(baseIndex)
	}
}

// 增加元素到堆
func (this *Heap) Add(val int64) {
	this.arr = append(this.arr, val)
	this.goUp(int64(len(this.arr)) - 1)
}

// 取出最大值(堆顶元素)
func (this *Heap) Take() int64 {
	if len(this.arr) == 0 {
		panic("error")
	}
	ret := this.arr[0]
	// 将最后一个元素和堆顶元素交换位置
	this.swap(0, int64(len(this.arr))-1)
	// 删除数组最后一项
	this.arr = this.arr[0 : len(this.arr)-1]
	this.goDown(0)
	return ret
}

// 替换堆顶元素
func (this *Heap) replace(val int64) {
	if len(this.arr) == 0 {
		panic("error")
	}
	this.arr[0] = val
	this.goDown(0)
}
