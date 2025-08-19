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

func exist(board [][]byte, word string) bool {
	used := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		used[i] = make([]bool, len(board[0]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				used[i][j] = true
				if dfs(board, used, i, j, 0, word) {
					return true
				}
				used[i][j] = false
			}
		}
	}
	return false
}

func dfs(board [][]byte, used [][]bool, i, j, target int, word string) bool {
	if board[i][j] != word[target] {
		return false
	}

	if target == len(word)-1 {
		return true
	}

	ret := false
	if j > 0 && !used[i][j-1] {
		used[i][j-1] = true
		ret = dfs(board, used, i, j-1, target+1, word) || ret
		used[i][j-1] = false
	}

	if j < len(board[0])-1 && !used[i][j+1] {
		used[i][j+1] = true
		ret = dfs(board, used, i, j+1, target+1, word) || ret
		used[i][j+1] = false
	}

	if i > 0 && !used[i-1][j] {
		used[i-1][j] = true
		ret = dfs(board, used, i-1, j, target+1, word) || ret
		used[i-1][j] = false
	}

	if i < len(board)-1 && !used[i+1][j] {
		used[i+1][j] = true
		ret = dfs(board, used, i+1, j, target+1, word) || ret
		used[i+1][j] = false
	}

	return ret
}

func combinationSum(candidates []int, target int) [][]int {
	var (
		ret  [][]int
		path []int
	)
	backtrack(candidates, target, 0, &ret, &path, 0)
	return ret
}

func backtrack(candidates []int, target int, start int, ret *[][]int, path *[]int, sum int) {
	if sum == target {
		tmp := make([]int, len(*path))
		copy(tmp, *path)
		*ret = append(*ret, tmp)
		return
	} else if sum > target {
		return
	}

	for i := start; i < len(candidates); i++ {
		if sum+candidates[i] <= target {
			*path = append(*path, candidates[i])
			backtrack(candidates, target, i, ret, path, sum+candidates[i])
			*path = (*path)[:len(*path)-1]
		}
	}
}

func partition(s string) [][]string {
	var (
		ret  [][]string
		path []string
	)
	backtrack2(s, 0, &ret, &path)
	return ret
}

func isHuiwen(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func backtrack2(s string, begin int, ret *[][]string, path *[]string) {
	if begin >= len(s) {
		tmp := make([]string, len(*path))
		copy(tmp, *path)
		*ret = append(*ret, tmp)
		return
	}

	for i := begin + 1; i <= len(s); i++ {
		if isHuiwen(s[begin:i]) {
			*path = append(*path, s[begin:i])
			backtrack2(s, i, ret, path)
			*path = (*path)[:len(*path)-1]
		}
	}
}

func solveNQueens(n int) [][]string {
	var (
		ret [][]string
	)
	path := make([][]byte, n)
	for i := 0; i < n; i++ {
		path[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			path[i][j] = '.'
		}
	}
	backtrackNQueues(n, 0, &ret, path)
	return ret
}

func backtrackNQueues(n, index int, ret *[][]string, path [][]byte) {
	if index >= n {
		tmp := make([]string, len(path))
		for i := 0; i < len(path); i++ {
			tmp[i] = string(path[i])
		}
		*ret = append(*ret, tmp)
		return
	}

	for i := 0; i < n; i++ {
		path[index][i] = 'Q'
		if isOk(index, i, path) {
			backtrackNQueues(n, index+1, ret, path)
		}
		path[index][i] = '.'
	}
}

func isOk(x, y int, a [][]byte) bool {
	m := len(a)
	n := m
	for j := 0; j < n; j++ {
		if j != y && a[x][j] == 'Q' {
			return false
		}
	}

	for i := 0; i < m; i++ {
		if i != x && a[i][y] == 'Q' {
			return false
		}
	}

	i, j := x, y
	for {
		i++
		j++
		if i >= m || j >= n {
			break
		}
		if a[i][j] == 'Q' {
			return false
		}
	}

	i, j = x, y
	for {
		i--
		j--
		if i < 0 || j < 0 {
			break
		}
		if a[i][j] == 'Q' {
			return false
		}
	}

	i, j = x, y
	for {
		i++
		j--
		if i >= m || j < 0 {
			break
		}
		if a[i][j] == 'Q' {
			return false
		}
	}

	i, j = x, y
	for {
		i--
		j++
		if i < 0 || j >= n {
			break
		}
		if a[i][j] == 'Q' {
			return false
		}
	}

	return true
}

func isValidSudoku(board [][]byte) bool {
	return backtrackSudo(board, 0)
}

func isValid(board [][]byte, x, y int) bool {
	for j := 0; j < len(board[x]); j++ {
		if j != y && board[x][j] == board[x][y] {
			return false
		}
	}
	for i := 0; i < len(board); i++ {
		if i != x && board[i][y] == board[x][y] {
			return false
		}
	}

	startI := (x / 3) * 3
	startJ := (y / 3) * 3
	for i := startI; i < startI+3; i++ {
		for j := startJ; j < startJ+3; j++ {
			if i != x && j != y {
				if board[i][j] == board[x][y] {
					return false
				}
			}
		}
	}

	return true
}

func backtrackSudo(board [][]byte, i int) bool {
	if i >= len(board) {
		return true
	}

	for j := 0; j < len(board[i]); j++ {
		if board[i][j] == '.' {
			continue
		}

		if !isValid(board, i, j) {
			return false
		}
	}

	return backtrackSudo(board, i+1)
}

func exist2(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				if backtracking3(board, word, i, j, 1) {
					return true
				}
			}
		}
	}
	return false
}

func backtracking3(board [][]byte, word string, i, j, idx int) bool {
	if idx >= len(word) {
		return true
	}

	direct := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for k := 0; k < len(direct); k++ {
		x := i + direct[k][0]
		y := j + direct[k][1]
		if x < 0 || y < 0 || x >= len(board) || y >= len(board[0]) {
			continue
		}

		if board[x][y] == word[idx] {
			find := backtracking3(board, word, x, y, idx+1)
			if find {
				return true
			}
		}
	}

	return false
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
	//fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))
	//fmt.Println(exist([][]byte{{'a', 'a'}}, "aaa"))
	//fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	//fmt.Println(partition("aab"))
	fmt.Println(isValidSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}))
}
