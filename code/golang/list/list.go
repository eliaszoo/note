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

func main() {
	list := NewList([]int{1, 2, 3, 4, 5, 6, 7, 8})
	PrintList(list)

	fmt.Println("")
	rvList := reverse(list)
	PrintList(rvList)

	fmt.Println("")

	node := findLastKth(rvList, 8)
	if node == nil {
		fmt.Println("not found")
	} else {
		fmt.Println(node.Val)
	}

}
