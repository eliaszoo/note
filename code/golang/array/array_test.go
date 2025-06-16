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

func TestMedian(t *testing.T) {
	m := Constructor()
	m.AddNum(1)
	m.AddNum(2)
	fmt.Println(m.FindMedian())
	m.AddNum(3)
	fmt.Println(m.FindMedian())
}

func TestFindDuplicate(t *testing.T) {
	fmt.Println(findDuplicate2([]int{3, 1, 3, 4, 2}))
}

func TestCanCompleteCircuit(t *testing.T) {
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 3, 2, 4, 1, 5, 3, 2, 4}, []int{1, 1, 1, 3, 2, 4, 3, 6, 7, 4, 3, 1}))
}

func TestCandy(t *testing.T) {
	fmt.Println(candy([]int{1, 2, 87, 87, 87, 2, 1}))
}
