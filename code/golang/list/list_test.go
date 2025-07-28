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

func TestPairSum(t *testing.T) {
	l1 := NewList([]int{47, 22, 81, 46, 94, 95, 90, 22, 55, 91, 6, 83, 49, 65, 10, 32, 41, 26, 83, 99, 14, 85, 42, 99, 89, 69, 30, 92, 32, 74, 9, 81, 5, 9})
	fmt.Println(pairSum(l1))
}
