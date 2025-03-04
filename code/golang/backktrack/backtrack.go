package main

import (
	"fmt"
	"sort"
)

var hash = map[string]string{
	"1": "",
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func digits(s string) []string {
	ret := make([]string, 0)
	var path string
	digitsBacktracking(s, len(s), 0, path, &ret)
	return ret
}

func digitsBacktracking(s string, k int, start int, path string, ret *[]string) {
	if len(path) == k {
		*ret = append(*ret, path)
		return
	}

	for i := start; i < len(s); i++ {
		chars, ok := hash[string(s[i])]
		if !ok {
			continue
		}

		for _, char := range chars {
			path += string(char)
			digitsBacktracking(s, k, i+1, path, ret)
			path = path[:len(path)-1]
		}
	}
}

func collect(digits, s string, index int, ret *[]string) {
	if len(digits) == index {
		*ret = append(*ret, s)
		return
	}
	for _, c := range hash[string(digits[index])] {
		collect(digits, s+string(c), index+1, ret)
	}
}

func letterCombinations(digits string) []string {
	var s string
	ret := make([]string, 0)
	collect(digits, s, 0, &ret)
	return ret
}

func premuteFunc(arr []int, path []int, used []bool, ret *[][]int) {
	if len(path) == len(arr) {
		tmp := make([]int, len(arr))
		copy(tmp, path)
		*ret = append(*ret, tmp)
		return
	}

	for i := 0; i < len(arr); i++ {
		if (used)[i] {
			continue
		}

		path = append(path, arr[i])
		(used)[i] = true
		premuteFunc(arr, path, used, ret)
		path = (path)[:len(path)-1]
		(used)[i] = false
	}
}

func permute(arr []int) [][]int {
	ret := make([][]int, 0)
	used := make([]bool, len(arr))
	path := make([]int, 0)
	premuteFunc(arr, path, used, &ret)
	return ret
}

func backtracking2(n, k, start int, numbers *[]int, ret *[][]int) {
	if len(*numbers) == k {
		tmp := make([]int, len(*numbers))
		copy(tmp, *numbers)
		*ret = append(*ret, tmp)
		return
	}

	for i := start; i <= n; i++ {
		*numbers = append(*numbers, i)
		backtracking2(n, k, i+1, numbers, ret)
		*numbers = (*numbers)[:len(*numbers)-1]
	}
}

func backtracking(n, k int) [][]int {
	ret := make([][]int, 0)
	var nums []int
	backtracking2(n, k, 1, &nums, &ret)
	return ret
}

func candidates(candidates []int, target int) [][]int {
	var ret [][]int
	var path []int
	candidatesBacktracking(candidates, target, 0, 0, &path, &ret)
	return ret
}

func candidatesBacktracking(candidates []int, target int, sum int, start int, path *[]int, ret *[][]int) {
	if sum > target {
		return
	} else if sum == target {
		tmp := make([]int, len(*path))
		copy(tmp, *path)
		*ret = append(*ret, tmp)
		return
	}

	for i := start; i < len(candidates); i++ {
		sum += candidates[i]
		*path = append(*path, candidates[i])
		candidatesBacktracking(candidates, target, sum, i, path, ret)
		*path = (*path)[:len(*path)-1]
		sum -= candidates[i]
	}
}

func candidates2(candidates []int, target int) [][]int {
	var ret [][]int
	var path []int
	sort.Ints(candidates)
	candidatesBacktracking2(candidates, target, 0, 0, &path, &ret)
	return ret
}

func candidatesBacktracking2(candidates []int, target int, sum int, start int, path *[]int, ret *[][]int) {
	if sum > target {
		return
	} else if sum == target {
		tmp := make([]int, len(*path))
		copy(tmp, *path)
		*ret = append(*ret, tmp)
		return
	}

	for i := start; i < len(candidates); i++ {
		if sum+candidates[i] > target {
			break
		}
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		sum += candidates[i]
		*path = append(*path, candidates[i])
		candidatesBacktracking2(candidates, target, sum, i+1, path, ret)
		*path = (*path)[:len(*path)-1]
		sum -= candidates[i]
	}
}

func isHuiWen(s string) bool {
	start := 0
	end := len(s) - 1

	for start < end {
		if s[start] != s[end] {
			return false
		}

		start++
		end--
	}

	return true
}

func subHuiWen(s string) []string {
	ret := make([]string, 0)
	var path string
	subHuiWenBacktracking(s, path, 0, &ret)
	return ret
}

func subHuiWenBacktracking(s string, path string, start int, ret *[]string) {
	if isHuiWen(path) {
		*ret = append(*ret, path)
	} else {
		return
	}

	for i := start; i < len(s); i++ {
		path += string(s[i])
		subHuiWenBacktracking(s, path, i+1, ret)
		path = path[:len(path)-1]
	}
}

func subHuiWen2(s string) [][]string {
	ret := make([][]string, 0)
	/*var path []string
	subHuiWenBacktracking2(s, &path, 0, 0, &ret)*/
	return ret
}

func subHuiWenBacktracking2(s string, path *[]string, tmp string, start, l int, ret *[][]string) {
	/*if l == len(s) {
		tmp := make([]string, len(*path))
		copy(tmp, *path)
		*ret = append(*ret, tmp)
		return
	}

	for i := start; i < len(s); i++ {
		tmp += string(s[i])
		subHuiWenBacktracking(s, path, tmp, i+1, , ret)
		tmp = tmp[:len(tmp)-1]
	}*/
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ret := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		target := 0 - nums[i]
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				arr := []int{nums[i], nums[left], nums[right]}
				ret = append(ret, arr)
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return ret
}

func main() {
	/*strs := letterCombinations("23")
	fmt.Println(strs)

	fmt.Println(permute([]int{1, 2, 3}))*/

	//fmt.Println(backtracking(4, 2))

	//fmt.Println(candidates([]int{2, 3, 5}, 8))
	//fmt.Println(candidates2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	//fmt.Println(digits("234"))
	//fmt.Println(subHuiWen("aab"))
	fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))
}
