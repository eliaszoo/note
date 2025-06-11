package main

import (
	"fmt"
	"math"
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, w := range stones {
		sum += w
	}
	mid := sum / 2

	dp := make([]int, mid+1)
	for i := 0; i < len(stones); i++ {
		for j := mid; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}

	return sum - dp[mid] - dp[mid]
}

// https://leetcode.cn/problems/combination-sum-iv/
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}

	return dp[target]
}

// https://leetcode.cn/problems/coin-change/description/ 零钱兑换，求最少硬币
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
	}

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			if dp[i-coin] != math.MaxInt {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt {
		return -1
	}

	return dp[amount]
}

// https://leetcode.cn/problems/perfect-squares/ 完全平方数
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt
	}

	for i := 1; i*i <= n; i++ {
		val := i * i
		for j := val; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-val]+1)
		}
	}
	return dp[n]
}

// https://leetcode.cn/problems/house-robber/ 打家劫舍1
func robfunc(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}

// https://leetcode.cn/problems/house-robber-ii/ 打家劫舍2
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	l := robfunc(nums[:len(nums)-1])
	r := robfunc(nums[1:])
	return max(l, r)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func build(a []int, index int) *TreeNode {
	if index >= len(a) {
		return nil
	}
	if a[index] == -1 {
		return nil
	}

	root := &TreeNode{Val: a[index], Left: nil, Right: nil}
	root.Left = build(a, 2*index+1)
	root.Right = build(a, 2*index+2)
	return root
}

func traverG2(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	var path []int
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			path = append(path, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		ret = append(ret, path)
		path = nil
	}
	return ret
}

func robTree(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}

	ldp := robTree(root.Left)
	rdp := robTree(root.Right)

	val0 := ldp[1] + rdp[1] + root.Val                // 偷当前节点
	val1 := max(ldp[0], ldp[1]) + max(rdp[0], rdp[1]) // 不偷当前节点，左右孩子都可偷

	return []int{val0, val1}
	//return max(ldp[0]+rdp[0]+root.Val, ldp[1]+rdp[1])
}

// https://leetcode.cn/problems/house-robber-iii/submissions/619087157/ 打家劫舍3
func robByTree(root *TreeNode) int {
	dp := robTree(root)
	return max(dp[0], dp[1])
}

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/ 买卖股票最佳时机2
func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0] // 持有股票的最大价值
	dp[0][1] = 0          // 不持有股票的最大价值
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[len(prices)-1][1]
}

func maxProfit1(prices []int) int {
	dp := make([][2]int, len(prices))
	dp[0][0] = -prices[0] // 当天买入
	dp[0][1] = 0          //
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	fmt.Println(dp)
	return dp[len(prices)-1][1]
}

func wordBreak(s string, wordDict []string) bool {
	dp := make([]int, len(s)+1)
	dp[0] = 0
	for i := 1; i <= len(s); i++ {
		for _, word := range wordDict {
			if i >= len(word) && s[i-len(word):i] == word {
				dp[i] = max(dp[i], dp[i-len(word)]+len(word))
			}
		}
		fmt.Println(dp)
	}

	return dp[len(s)] == len(s)
}

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}

	n := sum / 2
	dp := make([]int, n+1)
	for _, num := range nums {
		for i := n; i >= num; i-- {
			dp[i] = max(dp[i-num]+num, dp[i])
		}
		fmt.Println(dp)
	}
	return dp[n] == n
}

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if text1[i] == text2[j] {
				if i < 1 || j < 1 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j-1] + 1
				}
			} else if i >= 1 && j >= 1 {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			} else if i >= 1 {
				dp[i][j] = dp[i-1][j]
			} else if j >= 1 {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	fmt.Println(dp)
	return dp[m-1][n-1]
}

func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}

	var max int
	dp := make([]int, len(s))
	for i, c := range s {
		if c == ')' {
			if i > 0 {
				if s[i-1] == '(' {
					if i > 1 {
						dp[i] = dp[i-2] + 2
					} else {
						dp[i] = 2
					}
				} else if dp[i-1] > 0 {
					idx := i - 1 - dp[i-1]
					if idx >= 0 && s[idx] == '(' {
						dp[i] = dp[i-1] + 2
						if i-dp[i] >= 0 {
							dp[i] += dp[i-dp[i]]
						}
					}
				}
				if dp[i] > max {
					max = dp[i]
				}
			}
		}
	}

	return max
}
