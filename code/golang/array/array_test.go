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
