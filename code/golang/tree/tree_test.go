package main

import (
	"fmt"
	"testing"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	root := build([]int{1, 2}, 0)
	ans := diameterOfBinaryTree(root)
	fmt.Println(ans)
}

func TestIsBSTTree(t *testing.T) {
	root := build([]int{5, 4, 6, -1, -1, 3, 7}, 0)
	fmt.Println(isValidBST(root))
}

func TestFlatten(t *testing.T) {
	root := build([]int{1, 2, 5, 3, 4, -1, 6}, 0)
	flatten(root)
	fmt.Println(traverG2(root))
}

func TestPreorderIter(t *testing.T) {
	root := build([]int{1, 2, 3, 4, 5}, 0)
	fmt.Println(PreOrderIter(root))
}

func TestPostOrderIter(t *testing.T) {
	root := build([]int{1, 2, 3, 4, 5}, 0)
	fmt.Println(PostOrderIter(root))
}

func TestInorderIter(t *testing.T) {
	root := build([]int{1, 2, 3, 4, 5}, 0)
	fmt.Println(InorderIter(root))
}

func TestGetLeftNode(t *testing.T) {
	root := build([]int{1, 2, 3, 4, 5}, 0)
	var (
		maxHeight int
		ret       *TreeNode
	)
	GetLeftNode(root, 1, &maxHeight, &ret)
	fmt.Println(ret)
	fmt.Println(maxHeight)
}

func TestPathSum(t *testing.T) {
	root := build([]int{10, 5, -3, 3, 2, -1, 11, 3, -2, -1, 1}, 0)
	fmt.Println(traverG2(root))
	fmt.Println(pathSum(root, 8))
}

func TestMaxPathSum(t *testing.T) {
	root := build([]int{-10, 10, 9, -1, -1, 10, 5}, 0)
	fmt.Println(MaxPathSum(root))
}
