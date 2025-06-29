package str

import (
	"fmt"
)

// https://leetcode.cn/problems/minimum-window-substring/submissions/636639427/?envType=study-plan-v2&envId=top-100-liked
// 最小覆盖子串
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	cur := make(map[byte]int)

	for _, c := range t {
		ch := byte(c)
		if _, ok := cur[ch]; ok {
			cur[ch]++
		} else {
			cur[ch] = 1
		}
	}

	ret := ""
	min := len(s) + 1

	l := 0
	r := 0
	for {
		if r-l < len(t) || !isOk(cur) {
			if r >= len(s) { // 不能放在for那里，因为当达到末尾时，还需要往右滑动l
				break
			}
			r++
			if _, ok := cur[s[r-1]]; ok {
				cur[s[r-1]]--
			}
		}

		if isOk(cur) {
			if r-l < min {
				min = r - l
				ret = s[l:r]
			}
			if _, ok := cur[s[l]]; ok {
				cur[s[l]]++
			}
			l++
		}
	}
	return ret
}

func isOk(m map[byte]int) bool {
	for _, v := range m {
		if v > 0 {
			return false
		}
	}
	return true
}

func getNext(s string) []int {
	next := make([]int, len(s))
	next[0] = 0
	for i := 1; i < len(s); i++ {
		k := next[i-1]
		for k > 0 && s[i] != s[k] {
			k = next[k-1] // ???
		}

		if s[i] == s[k] {
			next[i] = k + 1
		} else {
			next[i] = 0
		}
	}

	return next
}

func kmp(s, m string) bool {
	if len(m) == 0 {
		return true
	}
	if len(s) < len(m) {
		return false
	}

	var i, j = 0, 0
	next := getNext(m)
	for i < len(s) {
		if s[i] == m[j] {
			i++
			j++

			if j == len(m) {
				return true
			}
		} else {
			if j > 0 {
				j = next[j-1]
			} else {
				i++
			}
		}
	}
	return false
}

// https://leetcode.cn/problems/zigzag-conversion/description/?envType=study-plan-v2&envId=top-interview-150
// N 字形变换
func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	ret := make([][]byte, numRows)

	zCols := numRows - 2
	for i, c := range s {
		round := i / (numRows + zCols)
		idx := i % (numRows + zCols)

		if idx == 0 {
			for i := 0; i < numRows; i++ {
				block := make([]byte, 1+zCols)
				for j := 0; j < 1+zCols; j++ {
					block[j] = byte(' ')
				}
				ret[i] = append(ret[i], block...)
			}
		}

		y := round * (1 + zCols)
		if idx >= numRows {
			y += idx - (numRows - 1)
		}
		x := idx
		if idx >= numRows {
			x = numRows - 1 - (idx - numRows + 1)
		}

		ret[x][y] = byte(c)
	}

	for i := 0; i < numRows; i++ {
		fmt.Println(string(ret[i]))
	}

	news := ""
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(ret[i]); j++ {
			if ret[i][j] != ' ' {
				news += string(ret[i][j])
			}
		}
	}
	return news
}

// https://leetcode.cn/problems/text-justification/?envType=study-plan-v2&envId=top-interview-150
// 文本左右对齐
func fullJustify(words []string, maxWidth int) []string {
	l := 0
	r := l

	width := len(words[0])
	ret := make([]string, 0)
	for i := 1; i < len(words); i++ {
		if width+len(words[i])+(i-l) > maxWidth {
			r = i - 1
			wordsNum := r - l + 1
			skipNum := maxWidth - width
			if wordsNum == 1 {
				s := words[l]
				for j := 0; j < skipNum; j++ {
					s += " "
				}
				ret = append(ret, s)
			} else {
				skipNumPer := skipNum / (wordsNum - 1)
				remain := skipNum - skipNumPer*(wordsNum-1)

				var s string
				for j := l; j <= r; j++ {
					s += words[j]
					add := skipNumPer
					if remain > 0 {
						add += 1
						remain--
					}

					if j == r {
						add = 0
					}
					for n := 0; n < add; n++ {
						s += " "
					}
				}
				ret = append(ret, s)
			}

			l = i
			width = len(words[i])
		} else {
			width += len(words[i])
		}
	}

	r = len(words) - 1
	var s string
	for j := l; j <= r; j++ {
		s += words[j]
		if j != r {
			s += " "
		}
	}
	n := maxWidth - len(s)
	for j := 0; j < n; j++ {
		s += " "
	}
	ret = append(ret, s)
	return ret
}
