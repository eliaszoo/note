package main

import "fmt"

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

func main() {
	val := cut([]int{1, 5, 8, 9, 10}, 4)
	fmt.Println(val)

	max := backup([]int{1, 2, 3, 4}, []int{1, 5, 8, 9}, 4)
	fmt.Println(max)
}
