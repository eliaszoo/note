package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func print(a [][]int) {
	for _, row := range a {
		for _, cell := range row {
			fmt.Print(cell, " ")
		}
		fmt.Println()
	}
}

func printArr(a []int) {
	for _, row := range a {
		fmt.Print(row, " ")
	}
	fmt.Println()
}

func backup01(weight []int, score []int, cap int) int {
	dp := make([][]int, len(weight))
	for i := range dp {
		dp[i] = make([]int, cap+1)
	}

	for i := 0; i < len(weight); i++ {
		dp[i][0] = 0
	}

	for j := 1; j <= cap; j++ {
		if j >= weight[0] {
			dp[0][j] = score[0]
		}
	}

	for j := 1; j <= cap; j++ {
		for i := 1; i < len(weight); i++ {
			if j >= weight[i] {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+score[i])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	print(dp)
	return dp[len(weight)-1][cap]
}

func backup01_2(weight []int, score []int, cap int) int {
	dp := make([]int, cap+1)
	for i := 0; i < len(weight); i++ {
		for j := cap; j >= weight[i]; j-- {
			dp[j] = max(dp[j], dp[j-weight[i]]+score[i])
		}
		printArr(dp)
	}

	return dp[cap]
}

func backup01_3(weight []int, score []int, cap int) int {
	dp := make([]int, cap+1)
	for i := 0; i < len(weight); i++ {
		for j := weight[i]; j <= cap; j++ {
			dp[j] = max(dp[j], dp[j-weight[i]]+score[i])
		}
		printArr(dp)
	}

	return dp[cap]
}

func main() {
	max := backup01([]int{1, 2, 3, 4}, []int{10, 15, 20, 35}, 10)
	fmt.Println(max)

	fmt.Println("-----------")
	max = backup01_2([]int{1, 2, 3, 4}, []int{10, 15, 20, 35}, 10)
	fmt.Println(max)
}
