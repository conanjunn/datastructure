package binary_tree

import (
	"fmt"
	"queue"
)

type node struct {
	leftChild  *node
	rightChild *node
	value      int64
}

type BinaryTree struct {
	root *node
	size int64
}

func (this *BinaryTree) Size() int64 {
	return this.size
}

func (this *BinaryTree) Add(val int64) {
	this.root = this.doAdd(val, this.root)
}

func (this *BinaryTree) doAdd(val int64, ele *node) *node {
	if ele == nil {
		ele = &node{
			value: val,
		}
		this.size++
	} else if ele.value > val {
		ele.leftChild = this.doAdd(val, ele.leftChild)
	} else if ele.value < val {
		ele.rightChild = this.doAdd(val, ele.rightChild)
	}
	return ele
}

// 前序遍历 先处理父节点再处理子节点
func (this *BinaryTree) PreOrder() {
	showArr, depth := this.doPreOrder(this.root, 0, make([][]interface{}, 0))
	fmt.Printf("树深度为%d\n", depth)
	for _, v := range showArr {
		for _, v2 := range v {
			fmt.Printf("%v ", v2)
		}
		fmt.Printf("\n")
	}
}
func (this *BinaryTree) doPreOrder(ele *node, depth int64, showArr [][]interface{}) ([][]interface{}, int64) {
	if int64(len(showArr)) <= depth {
		showArr = append(showArr, make([]interface{}, 0))
	}
	var resDepth int64
	if ele == nil {
		showArr[depth] = append(showArr[depth], "*")
		resDepth = depth
	} else {
		showArr[depth] = append(showArr[depth], ele.value)
		showArr, resDepth = this.doPreOrder(ele.leftChild, depth+1, showArr)
		showArr, resDepth = this.doPreOrder(ele.rightChild, depth+1, showArr)
	}
	return showArr, resDepth
}

// 中序遍历 数据将按顺序返回
func (this *BinaryTree) InOrder() {
	this.doInOrder(this.root, 0)
}
func (this *BinaryTree) doInOrder(ele *node, depth int64) {
	if ele == nil {
		return
	}
	this.doInOrder(ele.leftChild, depth+1)
	var i int64
	for i = 0; i < depth; i++ {
		fmt.Printf("*")
	}
	fmt.Printf("%d\n", ele.value)
	this.doInOrder(ele.rightChild, depth+1)
}

// 后序遍历 先遍历子节点在遍历父节点
func (this *BinaryTree) PostOrder() {
	this.doPostOrder(this.root, 0)
}
func (this *BinaryTree) doPostOrder(ele *node, depth int64) {
	if ele == nil {
		return
	}
	this.doPostOrder(ele.leftChild, depth+1)
	this.doPostOrder(ele.rightChild, depth+1)
	var i int64
	for i = 0; i < depth; i++ {
		fmt.Printf("*")
	}
	fmt.Printf("%d\n", ele.value)
}

// 层序遍历
func (this *BinaryTree) LevelOrder() {
	q := &queue.Queue{}
	q.Enqueue(this.root)
	size := q.Size()
	for size != 0 {
		iEle := q.Dequeue()
		ele := iEle.(*node)
		if ele != nil {
			fmt.Printf("%d ", ele.value)
			q.Enqueue(ele.leftChild)
			q.Enqueue(ele.rightChild)
		}
		size = q.Size()
	}
}
