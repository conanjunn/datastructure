package binary_tree

import (
	"fmt"
	"testing"
)

func TestPreOrder(t *testing.T) {
	tree := &BinaryTree{}
	tree.Add(3)
	tree.Add(2)
	tree.Add(4)
	tree.Add(5)
	fmt.Printf("pre\n")
	tree.PreOrder()
	fmt.Printf("in\n")
	tree.InOrder()
	fmt.Printf("post\n")
	tree.PostOrder()
	fmt.Printf("level\n")
	tree.LevelOrder()
}
