package main

func collect(digits, s string, index int, ret []string) {
	if len(digits) == index {
		ret = append(ret, s)
		return
	}
	for _, c := range digits {
		collect(digits, s+string(c), index+1, ret)
	}
}

func letterCombinations(digits string) []string {
	var s string
	ret := make([]string, 0)
	collect(digits, s, 0, ret)
	return ret
}

func main() {

}
