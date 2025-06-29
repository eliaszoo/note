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

func TestWordBreak(t *testing.T) {
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
}

func TestCanPartition(t *testing.T) {
	fmt.Println(canPartition([]int{3, 3, 3, 4, 5}))
}

func TestGetlongestCommonSubsequencee(t *testing.T) {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
}

func TestLongestValidParentheses(t *testing.T) {
	fmt.Println(longestValidParentheses(")()()(())("))

	fmt.Println(longestValidParentheses("(()"))    // 2
	fmt.Println(longestValidParentheses(")()())")) // 4
	fmt.Println(longestValidParentheses(""))       // 0
	fmt.Println(longestValidParentheses("()(()"))  // 2
	fmt.Println(longestValidParentheses("()(())")) // 6

}

func TestIsInterleave(t *testing.T) {
	fmt.Println(isInterleave("ab", "bc", "babc"))
}

func TestMaxProfit3(t *testing.T) {
	fmt.Println(maxProfit3([]int{2, 1, 4, 5, 2, 9, 7}))
}

func TestMaximalSquare(t *testing.T) {
	fmt.Println(maximalSquare([][]byte{{'0', '0', '0', '1'}, {'1', '1', '0', '1'}, {'1', '1', '1', '1'}, {'0', '1', '1', '1'}, {'0', '1', '1', '1'}}))
}

func TestIsSubsequenc(t *testing.T) {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
}
