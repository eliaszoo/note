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

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tmp := &ListNode{
		Val:  0,
		Next: head,
	}
	t := tmp
	p := head
	for p != nil && p.Next != nil {
		old := p.Next
		q := p.Next.Next
		p.Next.Next = p
		p.Next = q

		t.Next = old
		t = p
		p = q
	}
	return tmp.Next
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	pre := dummy
	next := head
	h, t := head, head
	for {
		i := 0
		for ; i < k-1 && t != nil; i++ {
			t = t.Next
		}
		if t == nil || i < k-1 {
			break
		}
		next = t.Next

		t.Next = nil
		pre.Next = reverse(h)
		h.Next = next

		pre = h
		t = next
		h = next
	}
	return dummy.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := getMiddle(head)
	rh := mid.Next
	mid.Next = nil
	l1 := sortList(head)
	l2 := sortList(rh)
	l := merge(l1, l2)
	return l
}

func getMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head.Next.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	p, q := l1, l2
	c := dummy
	for p != nil && q != nil {
		if p.Val < q.Val {
			c.Next = p
			p = p.Next
		} else {
			c.Next = q
			q = q.Next
		}
		c = c.Next
	}
	for ; p != nil; p = p.Next {
		c.Next = p
		c = c.Next
	}
	for ; q != nil; q = q.Next {
		c.Next = q
		c = c.Next
	}
	return dummy.Next
}

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	p, q := headA, headB
	for p != q {
		if p == nil {
			p = headB
		} else {
			p = p.Next
		}

		if q == nil {
			q = headA
		} else {
			q = q.Next
		}
	}
	return p
}

func isPalindrome(head *ListNode) bool {
	fast, slow := head, head
	cnt := 0
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		cnt++
	}

	p := reverse(slow)
	q := head
	for i := 0; i < cnt; i++ {
		if p.Val != q.Val {
			return false
		}
		p = p.Next
		q = q.Next
	}
	return true
}

func main() {
	list := NewList([]int{1, 2, 3, 4, 5})
	PrintList(list)
	//fmt.Println("-----\n")

	//PrintList(reverseKGroup(list, 2))

	PrintList(sortList(NewList([]int{4, 2, 1, 3})))

	/*fmt.Println("")
	rvList := reverse(list)
	PrintList(rvList)*/

	/*fmt.Println("")

	node := findLastKth(rvList, 8)
	if node == nil {
		fmt.Println("not found")
	} else {
		fmt.Println(node.Val)
	}*/

	//fmt.Println(reverseBetween(list, 1, 2))
	//PrintList(swapPairs(list))
}
