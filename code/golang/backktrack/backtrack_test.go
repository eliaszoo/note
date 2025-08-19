package main

import "testing"

func TestExist(t *testing.T) {
	t.Log(exist2([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCB"))
}
