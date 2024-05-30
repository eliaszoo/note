package main

import "fmt"

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

func main() {
	strs := letterCombinations("23")
	fmt.Println(strs)

	fmt.Println(permute([]int{1, 2, 3}))
}
