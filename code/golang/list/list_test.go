package main

import (
	"fmt"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	l1 := NewList([]int{2, 6, 4})
	l2 := NewList([]int{1, 5})
	fmt.Println(GetIntersectionNode(l1, l2))
}

func TestIsPalindrome(t *testing.T) {
	l1 := NewList([]int{1, 1, 2, 1})
	fmt.Println(isPalindrome(l1))
}

func TestReverseList(t *testing.T) {
	l1 := NewList([]int{1, 2, 3, 4, 5})
	fmt.Println(reverseList2(l1))
}
