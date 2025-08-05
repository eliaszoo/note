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

func TestCanFinish(t *testing.T) {
	fmt.Println(canFinish(5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}))
}

func TestTrie(t *testing.T) {
}

func TestMinOrder(t *testing.T) {
	fmt.Println(minReorder(6, [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}))
}

func TestNearestExit(t *testing.T) {
	fmt.Println(nearestExit([][]byte{
		{'+', '.', '+', '+', '+', '+', '+'},
		{'+', '.', '+', '.', '.', '.', '+'},
		{'+', '.', '+', '.', '+', '.', '+'},
		{'+', '.', '.', '.', '+', '.', '+'},
		{'+', '+', '+', '+', '+', '.', '+'},
	}, []int{0, 1}))
}
