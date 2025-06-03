package main

import (
	"fmt"
	"testing"
)

func TestLastStoneWeightII(t *testing.T) {
	fmt.Println(lastStoneWeightII([]int{31, 26, 33, 21, 40}))
}

func TestCoinChange(t *testing.T) {
	fmt.Println(coinChange([]int{2}, 3))
}

func TestRobTree(t *testing.T) {
	tree := build([]int{3, 2, 3, -1, 3, -1, 1}, 0)
	fmt.Println(traverG2(tree))
	fmt.Println(robByTree(tree))
}

func TestMaxProfit(t *testing.T) {
	fmt.Println(maxProfit1([]int{7, 1, 5, 3, 6, 4}))
}
