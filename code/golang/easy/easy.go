package main

import "fmt"

func removeElement(nums []int, val int) int {
	count := 0
	j := len(nums) - 1
	for i := 0; i < len(nums)-count; i++ {
		if nums[i] == val {
			count++
			for nums[j] == val {
				j--
			}
			if j > i {
				nums[i] = nums[j]
			}
		}
	}
	return count
}

func isMatch(a, b byte) bool {
	if a == '(' && b == ')' {
		return true
	} else if a == '{' && b == '}' {
		return true
	} else if a == '[' && b == ']' {
		return true
	}
	return false
}

func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	for _, c := range s {
		l := len(stack)
		if l > 0 && isMatch(stack[l-1], byte(c)) {
			stack = stack[:l-1]
		} else {
			stack = append(stack, byte(c))
		}
	}

	return len(stack) == 0
}

// 滑动窗口中的最大值
func maxSlidingWindow(nums []int, k int) []int {
	ret := make([]int, 0, len(nums))
	queue := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if len(queue) > 0 && (nums[i] > queue[0] || len(queue) >= k) {
			queue = queue[:0]
			queue = append(queue, nums[i])
		} else {
			queue = append(queue, nums[i])
		}

		if i >= k-1 {
			ret = append(ret, queue[0])
		}
	}
	return ret
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	idx := m + n - 1
	for ; i >= 0 || j >= 0; idx-- {
		if i < 0 {
			nums1[idx] = nums2[j]
			j--
		} else if j < 0 {
			nums1[idx] = nums1[i]
			i--
		} else {
			if nums1[i] > nums2[j] {
				nums1[idx] = nums1[i]
				i--
			} else {
				nums1[idx] = nums2[j]
				j--
			}
		}
	}
}

func main() {
	a := []int{3, 2, 2, 3}
	fmt.Println(removeElement(a, 3))
	fmt.Println(a)

	fmt.Println(isValid("()"))

	//fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3))

	b := []int{1, 2, 3, 0, 0, 0}
	merge(b, 3, []int{2, 5, 6}, 3)
	fmt.Println(b)
}
