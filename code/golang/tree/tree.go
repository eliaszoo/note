package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func build(a []int, index int) *TreeNode {
	if index >= len(a) {
		return nil
	}
	if a[index] == -1 {
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

func traverG2(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	var path []int
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			path = append(path, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		ret = append(ret, path)
		path = nil
	}
	return ret
}

func swap(root *TreeNode) {
	if root == nil {
		return
	}

	root.Left, root.Right = root.Right, root.Left
	swap(root.Left)
	swap(root.Right)
}

func symmetry(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if (left != nil && right == nil) || (left == nil && right != nil) {
		return false
	} else {
		if left.Val != right.Val {
			return false
		}
	}

	outside := symmetry(left.Left, right.Right)
	inside := symmetry(left.Right, right.Left)
	return outside && inside
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := height(root.Left) + 1
	r := height(root.Right) + 1
	return int(math.Max(float64(l), float64(r)))
}

func MinHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return minHeight(root)
}

func minHeight(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 1
	}

	l := height(root.Left) + 1
	r := height(root.Right) + 1
	return int(math.Min(float64(l), float64(r)))
}

func hasPathSum2(root *TreeNode, targetSum, cur int) bool {
	if root.Left == nil && root.Right == nil {
		return cur+root.Val == targetSum
	}

	cur += root.Val
	var l, r bool
	if root.Left != nil {
		l = hasPathSum2(root.Left, targetSum, cur)
	}
	if root.Right != nil {
		r = hasPathSum2(root.Right, targetSum, cur)
	}
	return l || r
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return hasPathSum2(root, targetSum, 0)
}

/*func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	tmpR := root.Right
	root.Right = root.Left
	flatten(root.Left)

}*/

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: preorder[0],
	}

	index := 0
	for ; index < len(inorder); index++ {
		if inorder[index] == root.Val {
			break
		}
	}

	root.Left = buildTree(preorder[1:index+1], inorder[:index])
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])
	return root
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var ret int

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := depth(root.Left)
	r := depth(root.Right)
	if l+r > ret {
		ret = l + r
	}

	return max(l, r) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	depth(root)
	return ret
}

func isValidBST(root *TreeNode) bool {
	pre := math.MinInt
	return isValidBSTFunc(root, &pre)
}

func isValidBSTFunc(root *TreeNode, pre *int) bool {
	if root == nil {
		return true
	}

	if !isValidBSTFunc(root.Left, pre) {
		return false
	}

	if *pre >= root.Val {
		return false
	}
	*pre = root.Val
	return isValidBSTFunc(root.Right, pre)
}

func flatten(root *TreeNode) {
	p := root
	for p != nil {
		l := p.Left
		r := p.Right
		if l != nil {
			q := l
			for ; q.Right != nil; q = q.Right {
			}

			q.Right = r
			p.Right = l
			p.Left = nil
		}

		p = p.Right
	}
}

func PreOrderIter(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		ret = append(ret, cur.Val)
	}

	return ret
}

func reverseArray(arr []int) {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

func PostOrderIter(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		ret = append(ret, cur.Val)
	}

	reverseArray(ret)
	return ret
}

func InorderIter(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var ret []int
	var stack []*TreeNode
	cur := root
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret = append(ret, top.Val)
			cur = top.Right
		}
	}

	return ret
}

func GetLeftNode(root *TreeNode, height int, maxHeight *int, ret **TreeNode) {
	if root.Left == nil && root.Right == nil {
		if height > *maxHeight {
			*maxHeight = height
			*ret = root
		}
		return
	}
	if root.Left != nil {
		GetLeftNode(root.Left, height+1, maxHeight, ret)
	}
	if root.Right != nil {
		GetLeftNode(root.Right, height+1, maxHeight, ret)
	}
}

func pathSum(root *TreeNode, targetSum int) int {
	var count int
	pathSum1(root, targetSum, &count)
	return count
}

func pathSum1(root *TreeNode, targetSum int, count *int) {
	if root == nil {
		return
	}
	pathSumFunc2(root, targetSum, count)
	pathSum1(root.Left, targetSum, count)
	pathSum1(root.Right, targetSum, count)
	return
}

func pathSumFunc2(root *TreeNode, targetSum int, count *int) {
	if root == nil {
		return
	}

	if targetSum == root.Val {
		*count++
	}

	pathSumFunc2(root.Left, targetSum-root.Val, count)
	pathSumFunc2(root.Right, targetSum-root.Val, count)
}

// 最大路径和
func MaxPathSum(root *TreeNode) int {
	var ret int
	maxPathSumFunc(root, &ret)
	return ret
}

func maxPathSumFunc(root *TreeNode, ret *int) int {
	if root == nil {
		return 0
	}

	lg := max(maxPathSumFunc(root.Left, ret), 0)
	rg := max(maxPathSumFunc(root.Right, ret), 0)

	val := lg + rg + root.Val
	if val > *ret {
		*ret = val
	}

	return root.Val + max(lg, rg)
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

	swap(root)
	fmt.Println(traverG(root))
	fmt.Println(traverG2(root))

	root2 := build([]int{1, 2, 2, 3, 4, 4, 3, 5}, 0)
	fmt.Println(traverG2(root2))
	fmt.Println(symmetry(root2.Left, root2.Right))

	fmt.Println(height(root))
	fmt.Println(height(root2))

	root3 := build([]int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, 1}, 0)
	fmt.Println(traverG2(root3))
	fmt.Println(hasPathSum(root3, 22))

	root4 := buildTree([]int{1, 2, 4, 5, 3}, []int{4, 2, 5, 1, 3})
	fmt.Println(traverG2(root4))
}
