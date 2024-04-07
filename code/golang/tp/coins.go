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

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
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
	//val := CoinChange([]int{1, 2, 5}, 11)
	//fmt.Println(val)

	fmt.Println(lengthOfLIS([]int{4, 10, 4, 3, 8, 9}))
}
