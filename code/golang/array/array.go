package array

import "sort"

func Rotate(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[m-1-i][n-j-1]
			matrix[m-1-i][n-j-1] = matrix[j][m-1-i]
			matrix[j][m-1-i] = tmp
		}
	}
}

func Rotate2(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	for i := 0; i < m/2; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matrix[m-i-1][j] = matrix[m-i-1][j], matrix[i][j]
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func findKthLargest(nums []int, k int) int {
	return quic(nums, 0, len(nums)-1, k)
}

func quic(nums []int, l, r, k int) int {
	idx := quickMove(nums, l, r)
	if idx+1 == k {
		return nums[idx]
	} else if idx+1 > k {
		return quic(nums, l, idx-1, k)
	} else {
		return quic(nums, idx+1, r, k)
	}
}

func quickMove(nums []int, l, r int) int {
	flag := nums[l]

	for l < r {
		for ; l < r; r-- {
			if nums[r] > flag {
				nums[l] = nums[r]
				break
			}
		}

		for ; l < r; l++ {
			if nums[l] < flag {
				nums[r] = nums[l]
				break
			}
		}
	}

	nums[l] = flag
	return l
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ret [][]int
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		target := 0 - nums[i]
		for l < r {
			sum := nums[l] + nums[r]
			if sum == target {
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				// 调过后面连续的重复值
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if sum < target {
				l++
			} else {
				r--
			}
		}
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func rainArea(arr []int) int {
	stack := make([]int, 0)
	var area int
	for i, num := range arr {
		n := len(stack)
		if n == 0 {
			stack = append(stack, i)
		} else {
			top := stack[n-1]
			topVal := arr[top]

			for num > topVal && len(stack) > 0 {
				stack = stack[:n-1]
				n--

				if len(stack) == 0 {
					break
				}

				top := stack[n-1]
				h := min(num, arr[top]) - topVal
				w := i - top - 1
				area += h * w

				topVal = arr[top]
			}

			stack = append(stack, i)
		}
	}

	return area
}

func isValid(s string) bool {
	stack := make([]byte, 0)
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, byte(c))
		} else {
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			if c == ')' && top == '(' {
				stack = stack[:len(stack)-1]
			} else if c == ']' && top == '[' {
				stack = stack[:len(stack)-1]
			} else if c == '}' && top == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			if nums[mid] > nums[l] {
				if target >= nums[l] {
					r = mid - 1
				} else {
					l = mid + 1
				}
			} else {
				r = mid - 1
			}
		} else {
			if nums[mid] < nums[r] {
				if target <= nums[r] {
					l = mid + 1
				} else {
					r = mid - 1
				}
			} else {
				l = mid + 1
			}
		}
	}

	return -1
}
