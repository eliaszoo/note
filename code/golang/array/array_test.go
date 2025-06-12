package array

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	m := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	Rotate(m)
	fmt.Println(m)

	m = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	Rotate(m)
	fmt.Println(m)
}

func TestRotate2(t *testing.T) {
	m := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	Rotate2(m)
	fmt.Println(m)

	m = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	Rotate2(m)
	fmt.Println(m)
}

func TestFindKthLargest(t *testing.T) {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}

func TestThreeSum(t *testing.T) {
	fmt.Println(threeSum([]int{2, -3, 0, -2, -5, -5, -4, 1, 2, -2, 2, 0, 2, -4, 5, 5, -10}))
}

func TestRainArea(t *testing.T) {
	fmt.Println(rainArea([]int{4, 2, 2, 1, 3}))
}

func TestIsValid(t *testing.T) {
	fmt.Println(isValid("()"))
}

func TestSearch(t *testing.T) {
	fmt.Println(search([]int{3, 1}, 1))
}
