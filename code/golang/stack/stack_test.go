package stack

import "testing"

func TestDecodeString(t *testing.T) {
	s := "3[a2[c]]"
	t.Log(decodeString(s))
}

func TestEvalRPN(t *testing.T) {
	s := []string{"4", "13", "5", "/", "+"}
	t.Log(evalRPN(s))
}
