package main

import (
	"fmt"
	"math"
)

func min(arr []int) int {
	min := 0x7fffffff
	for _, a := range arr {
		if a < min {
			min = a
		}
	}
	return min
}

func CoinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	tp := make([]int, amount+1)
	tp[0] = 0
	for k := 1; k <= amount; k++ {
		arr := make([]int, 0, len(coins))
		for _, coin := range coins {
			if k-coin >= 0 && tp[k-coin] != -1 {
				tmp := 1 + tp[k-coin]
				arr = append(arr, tmp)
			}
		}
		if len(arr) == 0 {
			tp[k] = -1
		} else {
			tp[k] = min(arr)
		}
	}
	return tp[amount]

}

func cut2(p []int, n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		q := -1
		for j := 1; j <= i; j++ {
			q = max(q, dp[i-j]+p[j-1])
		}
		dp[i] = q
	}
	return dp[n]
}

func bridge(t []int) int {
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func backup(weight []int, score []int, cw int) int {
	tp := make([][]int, len(weight))
	for i := 0; i < len(weight); i++ {
		tp[i] = make([]int, cw+1)
	}

	for j := 1; j <= cw; j++ {
		for i := 0; i < len(weight); i++ {
			if i > 0 && tp[i-1][j] > tp[i][j] {
				tp[i][j] = tp[i-1][j]
			}

			if j-weight[i] >= 0 && tp[i][j-weight[i]]+score[i] > tp[i][j] {
				tp[i][j] = tp[i][j-weight[i]] + score[i]
			}
		}
	}

	return tp[len(weight)-1][cw]
}

func cut(values []int, n int) int {
	tp := make([]int, n+1)
	tp[0] = 0
	for k := 1; k <= n; k++ {
		max := -1
		for i, val := range values {
			l := i + 1
			if k >= l {
				if tp[k-l]+val > max {
					max = tp[k-l] + val
				}
			} else {
				break
			}
		}
		tp[k] = max
	}
	return tp[n]
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	tp := make([]int, len(nums)+1)
	tp[0] = 0
	tp[1] = 1
	for k := 2; k <= len(nums); k++ {
		for i := 0; i < k-1; i++ {
			if nums[k-1] > nums[i] {
				tp[k] = max(tp[k], tp[i+1]+1)
			}
		}
	}
	return tp[len(nums)]
}

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		var ok = false
		for j := 0; j < len(wordDict); j++ {
			if i >= len(wordDict[j]) {
				if s[i-len(wordDict[j]):i] == wordDict[j] {
					ok = dp[i-len(wordDict[j])]
					if ok {
						break
					}
				}
			}
		}
		dp[i] = ok
	}

	return dp[len(s)]
}

func min2(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minimumTotal(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(triangle[i]))

		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = math.MaxInt
		}
	}

	dp[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j < len(dp[i-1]) {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			}

			if j > 0 {
				dp[i][j] = min2(dp[i][j], dp[i-1][j-1]+triangle[i][j])
			}
		}
	}

	last := len(triangle) - 1
	m := dp[last][0]
	for j := 1; j < len(triangle[last]); j++ {
		m = min2(m, dp[last][j])
	}

	return m
}

func main() {
	/*val := cut([]int{1, 5, 8, 9, 10}, 4)
	fmt.Println(val)

	max := backup([]int{1, 2, 3, 4}, []int{1, 5, 8, 9}, 4)
	fmt.Println(max)*/

	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
}
