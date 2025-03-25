package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{
		Val: vals[0],
	}

	p := head
	for i := 1; i < len(vals); i++ {
		node := &ListNode{Val: vals[i]}
		p.Next = node
		p = node
	}
	return head
}

func PrintList(head *ListNode) {
	for p := head; p != nil; p = p.Next {
		fmt.Print(fmt.Sprintf("%d->", p.Val))
	}
}

// 链表逆序
func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 递归到倒数第二个节点
	node := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}

// 查找倒数第k个节点
func findLastKth(head *ListNode, k int) *ListNode {
	fast, slow := head, head
	for i := 0; i < k; i++ {
		if fast == nil {
			return nil
		}
		fast = fast.Next
	}

	for ; fast != nil; fast = fast.Next {
		slow = slow.Next
	}
	return slow
}

func getNode(head *ListNode, n int) *ListNode {
	fast := head
	for i := 0; i < n-1; i++ {
		fast = fast.Next
	}
	return fast
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	tmp := &ListNode{Val: -1}
	tmp.Next = head

	left += 1
	right += 1

	lPrev := getNode(tmp, left-1)
	l := lPrev.Next
	r := getNode(tmp, right)
	rNext := r.Next
	r.Next = nil
	nw := reverse2(lPrev.Next)
	lPrev.Next = nw
	l.Next = rNext
	return tmp.Next
}

func reverse2(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	node := reverse2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}

func main() {
	list := NewList([]int{1, 2})
	PrintList(list)

	/*fmt.Println("")
	rvList := reverse(list)
	PrintList(rvList)

	fmt.Println("")

	node := findLastKth(rvList, 8)
	if node == nil {
		fmt.Println("not found")
	} else {
		fmt.Println(node.Val)
	}*/

	fmt.Println(reverseBetween(list, 1, 2))
}
