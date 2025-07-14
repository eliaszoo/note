package stack

import (
	"fmt"
	"strconv"
)

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

func evalRPN(tokens []string) int {
	stack := make([]string, 0, 4)
	popNumber := func() int {
		no := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		num, _ := strconv.Atoi(no)
		fmt.Println(no)
		return num
	}

	for _, s := range tokens {
		switch s {
		case "+":
			val := popNumber() + popNumber()
			stack = append(stack, strconv.Itoa(val))
		case "-":
			a := popNumber()
			b := popNumber()
			val := b - a
			stack = append(stack, strconv.Itoa(val))
		case "*":
			val := popNumber() * popNumber()
			stack = append(stack, strconv.Itoa(val))
		case "/":
			a := popNumber()
			b := popNumber()
			stack = append(stack, strconv.Itoa(b/a))
		default:
			stack = append(stack, s)
		}
	}

	return popNumber()
}

func calculate(s string) int {
	stack := make([]string, 0, len(s))
	i := 0
	var num string
	var op string
	for ; i < len(s); i++ {
		if s[i] == '(' {
			l := 1
			for j := i + 1; j < len(s); j++ {
				if s[j] == '(' {
					l++
				} else if s[j] == ')' {
					l--
					if l == 0 {
						val := calculate(s[i+1 : j])
						stack = append(stack, strconv.Itoa(val))
						if len(stack) == 3 {
							val := calc(stack)
							stack = stack[:0]
							stack = append(stack, strconv.Itoa(val))
						}

						i = j
						break
					}
				}
			}
		} else if s[i] == ' ' {
			continue
		} else if s[i] == '+' || s[i] == '-' {
			if num != "" {
				stack = append(stack, num)
				num = ""
			}
			if len(stack) == 3 || (len(stack) == 2 && op == "-") {
				val := calc(stack)
				stack = stack[:0]
				stack = append(stack, strconv.Itoa(val))
			}

			op = string(s[i])
			stack = append(stack, string(s[i]))
		} else {
			num += string(s[i])
		}
	}
	if num != "" {
		stack = append(stack, num)
	}

	return calc(stack)
}

func calc(stack []string) int {
	if len(stack) == 1 {
		v1, _ := strconv.Atoi(stack[0])
		return v1
	} else if len(stack) == 3 {
		v1, _ := strconv.Atoi(stack[0])
		v2, _ := strconv.Atoi(stack[2])
		op := stack[1]
		stack = stack[:0]
		if op == "+" {
			return v1 + v2
		} else {
			return v1 - v2
		}

	} else {
		op := stack[0]
		v1, _ := strconv.Atoi(stack[1])
		if op == "-" {
			return -v1
		} else {
			return v1
		}
	}

	return -1
}

func decodeString2(s string) string {
	var stack []string
	var curNum, curStr string
	//var ret string
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			curNum += string(s[i])
		} else if s[i] >= 'a' && s[i] <= 'z' {
			curStr += string(s[i])
		} else {
			if s[i] == '[' {
				stack = append(stack, curNum)
				stack = append(stack, curStr)
				curNum = ""
				curStr = ""
			} else {
				n := len(stack)
				lastStr := stack[n-1]
				numStr := stack[n-2]
				stack = stack[:n-2]

				tmp := curStr
				num, _ := strconv.Atoi(numStr)
				for j := 0; j < num-1; j++ {
					curStr += tmp
				}
				curStr = lastStr + curStr
			}
		}
	}
	return curStr
}
