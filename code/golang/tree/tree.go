package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func build(a []int, index int) *TreeNode {
	if index >= len(a) {
		return nil
	}

	root := &TreeNode{Val: a[index], Left: nil, Right: nil}
	root.Left = build(a, 2*index+1)
	root.Right = build(a, 2*index+2)
	return root
}

func preTraversal(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}

	*ret = append(*ret, root.Val)
	preTraversal(root.Left, ret)
	preTraversal(root.Right, ret)
}

func preTraversalIter(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}

	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		*ret = append(*ret, cur.Val)

		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}
}

func midTraversal(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}

	midTraversal(root.Left, ret)
	*ret = append(*ret, root.Val)
	midTraversal(root.Right, ret)
}

func midTraversalIter(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}

	var stack []*TreeNode
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		*ret = append(*ret, cur.Val)

		cur = cur.Right
	}
}

func postTraversal(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}

	postTraversal(root.Left, ret)
	postTraversal(root.Right, ret)
	*ret = append(*ret, root.Val)
}

func reverse(a []int) {
	i := 0
	j := len(a) - 1
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
}

func postTraversalIter(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}

	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		*ret = append(*ret, cur.Val)
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
	}

	reverse(*ret)
}

func traverG(root *TreeNode) []int {
	ret := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		ret = append(ret, cur.Val)
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
	return ret
}

func main() {
	root := build([]int{1, 2, 3, 4, 5, 6, 7}, 0)
	var tmp []int
	preTraversal(root, &tmp)
	fmt.Println(tmp)
	tmp = tmp[:0]
	preTraversalIter(root, &tmp)
	fmt.Println(tmp)

	tmp = tmp[:0]
	midTraversal(root, &tmp)
	fmt.Println(tmp)
	tmp = tmp[:0]
	midTraversalIter(root, &tmp)
	fmt.Println(tmp)

	tmp = tmp[:0]
	postTraversal(root, &tmp)
	fmt.Println(tmp)

	tmp = tmp[:0]
	postTraversalIter(root, &tmp)
	fmt.Println(tmp)

	fmt.Println(traverG(root))
}
