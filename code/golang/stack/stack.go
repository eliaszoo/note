package stack

import "strconv"

func decodeString(s string) string {
	numStack := []int{}
	tmpNumber := ""
	str2 := ""
	var ret string
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			tmpNumber += string(s[i])
		} else if s[i] == '[' {
			n, _ := strconv.Atoi(tmpNumber)
			numStack = append(numStack, n)
			tmpNumber = ""
			str2 = ""
		} else if s[i] == ']' {
			num := numStack[len(numStack)-1]
			for i := 0; i < num; i++ {
				ret += str2
			}
			numStack = numStack[:len(numStack)-1]
		} else {
			str2 += string(s[i])
		}

	}
	return ret
}
