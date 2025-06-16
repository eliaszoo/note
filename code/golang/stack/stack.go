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
