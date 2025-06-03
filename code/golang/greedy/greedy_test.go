package greedy

import (
	"fmt"
	"testing"
)

func TestJump(t *testing.T) {
	fmt.Println(jump([]int{1, 1, 1, 1}))
}

func TestProductExceptSelf(t *testing.T) {
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}

func TestCandy(t *testing.T) {
	fmt.Println(candy([]int{1, 3, 4, 5, 2}))
}

func TestSpiralOrder(t *testing.T) {
	fmt.Println(spiralOrder2([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}))
}
