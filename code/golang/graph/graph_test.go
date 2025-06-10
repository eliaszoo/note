package graph

import (
	"fmt"
	"testing"
)

func TestNumIslands(t *testing.T) {
	num := numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	})
	fmt.Println(num)
}

func TestOrangesRotting(t *testing.T) {
	fmt.Println(orangesRotting([][]int{
		{2, 2},
		{1, 1},
		{0, 0},
		{2, 0},
	}))
}
