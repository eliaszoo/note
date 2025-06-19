package str

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
			k = next[k-1]
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
