package stack

import "testing"

func TestDecodeString(t *testing.T) {
	s := "3[a2[c]]"
	t.Log(decodeString(s))
}
