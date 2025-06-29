package str

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMinWindow(t *testing.T) {
	fmt.Println(minWindow("a", "a"))
}

func TestKmp(t *testing.T) {
	fmt.Println(kmp("abababca", "ababa"))
}

func TestGetNext(t *testing.T) {
	tests := []struct {
		s    string
		want []int
	}{
		{"abcabc", []int{0, 0, 0, 1, 2, 3}},
		{"aabaaac", []int{0, 1, 0, 1, 2, 2, 0}},
		{"abcdabca", []int{0, 0, 0, 0, 1, 2, 3, 1}},
	}
	for _, tt := range tests {
		got := getNext(tt.s)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("getNext(%q) = %v, want %v", tt.s, got, tt.want)
		}
	}
}

func TestKMP(t *testing.T) {
	tests := []struct {
		s, m string
		want bool
	}{
		{"hello", "ll", true},
		{"hello", "world", false},
		{"ababababc", "ababc", true},
		{"a", "a", true},
		{"a", "b", false},
		{"", "", true},
	}
	for _, tt := range tests {
		got := kmp(tt.s, tt.m)
		if got != tt.want {
			t.Errorf("kmp(%q, %q) = %v, want %v", tt.s, tt.m, got, tt.want)
		}
	}
}

func TestConvert(t *testing.T) {
	fmt.Println(convert("PAYPALISHIRING", 3))
}
